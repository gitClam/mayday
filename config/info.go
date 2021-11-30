// Package classification JieTong API.
//
// 除了用户登录和注册以及头像获取三个接口
// 其他的都需要用户携带TOKEN进行用户验证，否则无法访问接口
//
// TOKEN 格式 ： KEY：Authorization VALUE： "JWT " + 登录时返回的对应TOKEN   （放在请求的header中）
//
//     Schemes: http
//     Host: 47.107.108.127
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//
// swagger:meta
package classification
