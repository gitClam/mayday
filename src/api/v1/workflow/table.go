package workflow

import "github.com/kataras/iris/v12"

// @Tags Table
// @Summary 获取表单详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息"
// @Router /table/get/table/{id:int} [get]
func GetTableById(ctx iris.Context) {

}

// @Tags Table
// @Summary 获取当前用户的表单草稿列表
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息"
// @Router /table/get/table-draft [get]
func GetTableDraftByUser(ctx iris.Context) {

}

// @Tags Table
// @Summary 获取表单草稿详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单草稿id"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息"
// @Router /table/get/table-draft/{id:int} [get]
func GetTableDraftById(ctx iris.Context) {

}

// @Tags Table
// @Summary 创建表单
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "表单信息"
// @Success 200 {object} utils.Response
// @Router /table/create/table [post]
func CreateTable(ctx iris.Context) {

}

// @Tags Table
// @Summary 创建表单
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "表单信息"
// @Success 200 {object} utils.Response
// @Router /table/update/table-draft [post]
func CreateTableDraft(ctx iris.Context) {

}

// @Tags Table
// @Summary 修改表单信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "表单信息"
// @Success 200 {object} utils.Response
// @Router /table/update/table [post]
func UpdateTable(ctx iris.Context) {

}

// @Tags Table
// @Summary 修改表单草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "表单信息"
// @Success 200 {object} utils.Response
// @Router /table/update/table-draft [post]
func UpdateTableDraft(ctx iris.Context) {

}

// @Tags Table
// @Summary 删除表单
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id"
// @Success 200 {object} utils.Response
// @Router /table/delete/table [post]
func DeleteTable(ctx iris.Context) {

}

// @Tags Table
// @Summary 删除表单草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id"
// @Success 200 {object} utils.Response
// @Router /table/delete/table-draft [post]
func DeleteTableDraft(ctx iris.Context) {

}
