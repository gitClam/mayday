package middleware

import (
	"fmt"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	"mayday/src/model/user"
	"mayday/src/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"time"

	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type (
	errorHandler func(context.Context, int, error)

	TokenExtractor func(context.Context) (string, error)

	Jwts struct {
		Config Config
	}
)

var (
	jwts *Jwts
	lock sync.Mutex
)

func Serve(ctx context.Context) bool {
	ConfigJWT()
	if !jwts.CheckJWT(ctx) {
		return false
	}
	return ParseTokenTest(ctx)
}
func ParseTokenTest(ctx context.Context) bool {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)

	id, ok1 := mapClaims["id"].(float64)
	name, ok2 := mapClaims["name"].(string)

	if ok1 && ok2 {
		ctx.Values().Set("user", user.SdUser{Id: int(id), Name: name})
		return true
	}

	return false
}

// ParseToken 解析token的信息为当前用户
func ParseToken(ctx context.Context) (*user.SdUser, bool) {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)

	id, ok1 := mapClaims["id"].(float64)
	name, ok2 := mapClaims["name"].(string)

	if !ok1 || !ok2 {
		return nil, false
	}
	return &user.SdUser{Id: int(id), Name: name}, true
}

func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "jwt" {
		return "", fmt.Errorf("Authorization header format must be JWT {token}")
	}

	return authHeaderParts[1], nil
}

func (m *Jwts) logf(format string, args ...interface{}) {
	if m.Config.Debug {
		global.GVA_LOG.Info(fmt.Sprintf(format, args...))
	}
}

func (m *Jwts) Get(ctx context.Context) *jwt.Token {
	return ctx.Values().Get(m.Config.ContextKey).(*jwt.Token)
}

func (m *Jwts) CheckJWT(ctx context.Context) bool {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return true
		}
	}

	token, err := m.Config.Extractor(ctx)
	if err != nil {
		m.Config.ErrorHandler(ctx, resultcode.TokenExactFail, err)
		return false
	}

	if token == "" {
		if m.Config.CredentialsOptional {
			m.logf(" Error: No credentials found (CredentialsOptional=true)")
			return true
		}

		m.Config.ErrorHandler(ctx, resultcode.TokenParseFailAndEmpty, fmt.Errorf(" Error: No credentials found (CredentialsOptional=false)"))
		return false
	}

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	if err != nil {
		m.Config.ErrorHandler(ctx, resultcode.TokenParseFail, err)
		return false
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		m.Config.ErrorHandler(ctx, resultcode.TokenParseFail, fmt.Errorf(message)) // 算法错误
		return false
	}

	if !parsedToken.Valid {
		m.Config.ErrorHandler(ctx, resultcode.TokenParseFailAndInvalid, fmt.Errorf("无效的TOKEN"))
		return false
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				m.Config.ErrorHandler(ctx, resultcode.TokenExpire, fmt.Errorf("TOKEN已过期"))
				return false
			}
		}
	}

	ctx.Values().Set(m.Config.ContextKey, parsedToken)
	return true
}

// ------------------------------------------------------------------------

// ConfigJWT jwt中间件配置
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if jwts != nil {
		return
	}

	c := Config{
		ContextKey: global.GVA_CONFIG.JWT.DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(global.GVA_CONFIG.JWT.Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, code int, errMsg error) {
			utils.Responser.Fail(ctx, code, errMsg)
		},
		// 指定func用于提取请求中的token
		Extractor:           FromAuthHeader,
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	jwts = &Jwts{Config: c}
}

type Claims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 在登录成功的时候生成token
func GenerateToken(user *user.SdUser) (string, error) {

	expireTime := time.Now().Add(time.Duration(global.GVA_CONFIG.JWT.JWTTimeout) * time.Second)

	claims := Claims{
		user.Id,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(global.GVA_CONFIG.JWT.Secret))
	return token, err
}

type Config struct {
	// The function that will return the Key to validate the JWT.
	// It can be either a shared secret or a public key.
	// Default value: nil
	ValidationKeyGetter jwt.Keyfunc
	// The name of the property in the request where the user (&token) information
	// from the JWT will be stored.
	// Default value: "jwts"
	ContextKey string
	// The function that will be called when there's an error validating the token
	// Default value:
	ErrorHandler errorHandler
	// A boolean indicating if the credentials are required or not
	// Default value: false
	CredentialsOptional bool
	// A function that extracts the token from the request
	// Default: FromAuthHeader (i.e., from Authorization header as bearer token)
	Extractor TokenExtractor
	// Debug flag turns on debugging output
	// Default: false
	Debug bool
	// When set, all requests with the OPTIONS method will use authentication
	// if you enable this option you should register your route with iris.Options(...) also
	// Default: false
	EnableAuthOnOptions bool
	// When set, the middelware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// Default: nil
	SigningMethod jwt.SigningMethod
	// When set, the expiration time of token will be check every time
	// if the token was expired, expiration error will be returned
	// Default: false
	Expiration bool
}
