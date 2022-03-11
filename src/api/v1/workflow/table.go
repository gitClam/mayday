package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/model/common/resultcode"
	workflowModel "mayday/src/model/workflow"
	workflowSever "mayday/src/service/workflow"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Table
// @Summary 获取表单详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息 错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /table/get/table [get]
func GetTableById(ctx iris.Context) {
	var tableId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		tableId = append(tableId, num)
	}
	workflowSever.GetTableById(ctx, tableId)
}

// @Tags Table
// @Summary 获取当前用户的表单草稿列表
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息"
// @Router /table/get/table-draft/user [get]
func GetTableDraftByUser(ctx iris.Context) {
	workflowSever.GetTableDraftByUser(ctx)
}

// @Tags Table
// @Summary 获取表单草稿详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单草稿id"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回表单的详细信息"
// @Router /table/get/table-draft/id [get]
func GetTableDraftById(ctx iris.Context) {
	var ids []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		ids = append(ids, num)
	}
	workflowSever.GetTableDraftById(ctx, ids)
}

// @Tags Table
// @Summary 创建表单
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflow.CreateTableReq true "表单信息"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes}
// @Router /table/create/table [post]
func CreateTable(ctx iris.Context) {

	var tableReq workflowModel.CreateTableReq
	if err := ctx.ReadJSON(&tableReq); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	workflowSever.CreateTable(ctx, tableReq)
}

// @Tags Table
// @Summary 创建表单草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflow.CreateTableDraftReq true "表单草稿信息"
// @Success 200 {object} utils.Response
// @Router /table/create/table-draft [post]
func CreateTableDraft(ctx iris.Context) {

	var tableDraftReq workflowModel.CreateTableDraftReq
	if err := ctx.ReadJSON(&tableDraftReq); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	workflowSever.CreateTableDraft(ctx, tableDraftReq)
}

// @Tags Table
// @Summary 修改表单信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflow.UpdateTableReq true "表单信息"
// @Success 200 {object} utils.Response
// @Router /table/update/table [post]
func UpdateTable(ctx iris.Context) {

	var tableReq workflowModel.UpdateTableReq
	if err := ctx.ReadJSON(&tableReq); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	workflowSever.UpdateTable(ctx, tableReq)

}

// @Tags Table
// @Summary 修改表单草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflow.UpdateTableDraftReq true "表单草稿信息"
// @Success 200 {object} utils.Response
// @Router /table/update/table-draft [post]
func UpdateTableDraft(ctx iris.Context) {

	var tableDraftReq workflowModel.UpdateTableDraftReq
	if err := ctx.ReadJSON(&tableDraftReq); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	workflowSever.UpdateTableDraft(ctx, tableDraftReq)
}

// @Tags Table
// @Summary 删除表单
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id"
// @Success 200 {object} utils.Response
// @Router /table/delete/table [delete]
func DeleteTable(ctx iris.Context) {

	var tableId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		tableId = append(tableId, num)
	}
	workflowSever.DeleteTable(ctx, tableId)
}

// @Tags Table
// @Summary 删除表单草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "表单id"
// @Success 200 {object} utils.Response
// @Router /table/delete/table-draft [delete]
func DeleteTableDraft(ctx iris.Context) {
	var tableId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		tableId = append(tableId, num)
	}
	workflowSever.DeleteTableDraft(ctx, tableId)

}
