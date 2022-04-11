package order

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	"mayday/src/model/order"
	"mayday/src/model/process"
	"mayday/src/model/user"
	order2 "mayday/src/service/order"
	"mayday/src/utils"
	"strconv"
	"time"
)

//processId
//workOrderId
// 流程结构包括节点，流转和模版
func ProcessStructure(ctx iris.Context) {
	processId := ctx.FormValue("WorkflowId")
	if processId == "" {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}
	workOrderId := ctx.FormValue("OrderId")
	if workOrderId == "" {
		workOrderId = "0"
	}
	workOrderIdInt, _ := strconv.Atoi(workOrderId)
	processIdInt, _ := strconv.Atoi(processId)
	result, err := order2.MakeProcessStructure(ctx, processIdInt, workOrderIdInt)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.Fail, err)
		return
	}
	if workOrderIdInt != 0 {
		currentState := result["workOrder"].(order.SdOrder).CurrentState
		userAuthority, err := order2.JudgeUserAuthority(ctx, workOrderIdInt, currentState)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.PermissionsLess, err)
			return
		}
		result["userAuthority"] = userAuthority
	}
	utils.Responser.OkWithDetails(ctx, result)
}

// 工单列表
//title := w.GinObj.FormValue("title")
//startTime := w.GinObj.FormValue("startTime")
//endTime := w.GinObj.FormValue("endTime")
//isEnd := w.GinObj.FormValue("isEnd")
//processor := w.GinObj.FormValue("processor")
//priority := w.GinObj.FormValue("priority")
//creator := w.GinObj.FormValue("creator")
//classify := ctx.FormValue("classify")
func WorkOrderList(ctx iris.Context) {
	/*
		1. 待办工单
		2. 我创建的
		3. 我相关的
		4. 所有工单
	*/

	var (
		result      interface{}
		err         error
		classifyInt int
	)
	classify := ctx.FormValue("Classify")
	if classify == "" {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}

	classifyInt, _ = strconv.Atoi(classify)
	w := order2.WorkOrder{
		Classify: classifyInt,
		Context:  ctx,
	}
	result, _, err = w.WorkOrderList()
	if err != nil {
		global.GVA_LOG.Warn("查询工单数据失败，%v", zap.Error(err))
		return
	}

	utils.Responser.OkWithDetails(ctx, result)
}

func WorkOrderListLength(ctx iris.Context) {
	/*
		1. 待办工单
		2. 我创建的
		3. 我相关的
		4. 所有工单
	*/

	var (
		//result      interface{}
		err         error
		classifyInt int
	)
	classify := ctx.FormValue("Classify")
	if classify == "" {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}

	classifyInt, _ = strconv.Atoi(classify)
	w := order2.WorkOrder{
		Classify: classifyInt,
		Context:  ctx,
	}
	_, len, err := w.WorkOrderList()
	if err != nil {
		global.GVA_LOG.Warn("查询工单数据失败，%v", zap.Error(err))
		return
	}

	utils.Responser.OkWithDetails(ctx, struct {
		Length int
	}{Length: len})
}

// 处理工单
func ProcessWorkOrder(ctx iris.Context) {
	var (
		err           error
		userAuthority bool
		handle        order2.Handle
		params        struct {
			Tasks          []string
			TargetState    string                   `json:"target_state"`    // 目标状态
			SourceState    string                   `json:"source_state"`    // 源状态
			WorkOrderId    int                      `json:"work_order_id"`   // 工单ID
			Circulation    string                   `json:"circulation"`     // 流转ID
			FlowProperties int                      `json:"flow_properties"` // 流转类型 0 拒绝，1 同意，2 其他
			Remarks        string                   `json:"remarks"`         // 处理的备注信息
			Tpls           []map[string]interface{} `json:"tpls"`            // 表单数据
			IsExecTask     bool                     `json:"is_exec_task"`    // 是否执行任务
		}
	)

	err = ctx.ReadJSON(&params)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}

	// 处理工单
	userAuthority, err = order2.JudgeUserAuthority(ctx, params.WorkOrderId, params.SourceState)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.PermissionsLess, err)
		return
	}
	if !userAuthority {
		utils.Responser.Fail(ctx, resultcode.PermissionsLess, err)
		return
	}

	err = handle.HandleWorkOrder(
		ctx,
		params.WorkOrderId,    // 工单ID
		params.Tasks,          // 任务列表
		params.TargetState,    // 目标节点
		params.SourceState,    // 源节点
		params.Circulation,    // 流转标题
		params.FlowProperties, // 流转属性
		params.Remarks,        // 备注信息
		params.Tpls,           // 工单数据更新
		params.IsExecTask,     // 是否执行任务
	)
	if err != nil {
		global.GVA_LOG.Error("工单处理失败", zap.Error(err))
		utils.Responser.Fail(ctx, resultcode.Fail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// 结束工单
func UnityWorkOrder(c iris.Context) {
	var (
		err           error
		workOrderId   string
		workOrderInfo order.SdOrder
		userInfo      user.SdUser
	)

	workOrderId = c.FormValue("workOrderId")
	if workOrderId == "" {
		utils.Responser.Fail(c, resultcode.DataReceiveFail)
		return
	}
	tx := global.GVA_DB.NewSession()
	err = tx.Begin()
	if err != nil {
		utils.Responser.Fail(c, resultcode.DataReceiveFail, err)
		return
	}

	// 查询工单信息
	has, err := tx.Where("id = ?", workOrderId).Get(&workOrderInfo)
	if !has || err != nil {
		utils.Responser.Fail(c, resultcode.DataSelectFail, err)
		return
	}
	if workOrderInfo.IsEnd == 1 {
		utils.Responser.Fail(c, resultcode.OrderIsEnd, err)
		return
	}

	// 更新工单状态
	effect, err := tx.Where("id = ?", workOrderId).Update(map[string]interface{}{"is_end": 1})
	if effect <= 0 || err != nil {
		tx.Rollback()
		utils.Responser.Fail(c, resultcode.DataUpdateFail, err)
		return
	}

	// 获取当前用户信息
	err = tx.Where("id = ?", c.Values().Get("user").(user.SdUser).Id).Find(&userInfo)
	if err != nil {
		tx.Rollback()
		utils.Responser.Fail(c, resultcode.DataSelectFail, err)
		return
	}

	// 写入历史
	effect, err = tx.Insert(&process.CirculationHistory{
		Title:       workOrderInfo.Title,
		WorkOrder:   workOrderInfo.Id,
		State:       "结束工单",
		Circulation: "结束",
		Processor:   userInfo.Name,
		ProcessorId: c.Values().Get("user").(user.SdUser).Id,
		Remarks:     "手动结束工单。",
		Status:      2,
	})
	if effect <= 0 || err != nil {
		utils.Responser.Fail(c, resultcode.DataDeleteFail, err)
		return
	}

	tx.Commit()

	utils.Responser.Ok(c)
}

//form
//			WorkOrderId int    `json:"work_order_id"`
//			NodeId      string `json:"node_id"`
//			UserId      int    `json:"user_id"`
//			Remarks     string `json:"remarks"`
// 转交工单

func InversionWorkOrder(c iris.Context) {
	var (
		cirHistoryValue   []order.SdOrderCirculationHistory
		err               error
		workOrderInfo     order.SdOrder
		stateList         []map[string]interface{}
		stateValue        []byte
		currentState      map[string]interface{}
		userInfo          user.SdUser
		currentUserInfo   user.SdUser
		costDurationValue int64
		params            struct {
			WorkOrderId int    `json:"work_order_id"`
			NodeId      string `json:"node_id"`
			UserId      int    `json:"user_id"`
			Remarks     string `json:"remarks"`
		}
	)

	// 获取当前用户信息
	has, err := global.GVA_DB.Id(c.Values().Get("user").(user.SdUser).Id).Get(&currentUserInfo)
	if !has || err != nil {
		utils.Responser.Fail(c, resultcode.DataSelectFail, err)
		return
	}
	err = c.ReadForm(&params)
	if err != nil {
		utils.Responser.Fail(c, resultcode.DataReceiveFail, err)
		return
	}

	// 查询工单信息
	has, err = global.GVA_DB.Where("id = ?", params.WorkOrderId).Get(&workOrderInfo)
	if !has || err != nil {
		utils.Responser.Fail(c, resultcode.DataSelectFail, err)
		return
	}

	// 序列化节点数据
	err = json.Unmarshal(workOrderInfo.State, &stateList)
	if err != nil {
		global.GVA_LOG.Error("节点数据反序列化失败，%v", zap.Error(err))
		return
	}

	for _, s := range stateList {
		if s["id"].(string) == params.NodeId {
			s["processor"] = []interface{}{params.UserId}
			s["process_method"] = "person"
			currentState = s
			break
		}
	}

	stateValue, err = json.Marshal(stateList)
	if err != nil {
		global.GVA_LOG.Error("节点数据反序列化失败，%v", zap.Error(err))
		return
	}

	tx := global.GVA_DB.NewSession()
	tx.Begin()

	// 更新数据
	_, err = tx.Table(new(order.SdOrder)).
		Where("id = ?", params.WorkOrderId).
		Update(map[string]interface{}{"state": stateValue})
	if err != nil {
		utils.Responser.Fail(c, resultcode.DataUpdateFail, err)
		return
	}

	// 查询用户信息
	has, err = global.GVA_DB.Id(params.UserId).Get(&userInfo)
	if !has || err != nil {
		utils.Responser.Fail(c, resultcode.DataSelectFail, err)
		return
	}

	// 流转历史写入
	err = global.GVA_DB.Where("order_id = ?", params.WorkOrderId).Desc("create_time").Find(&cirHistoryValue)
	if err != nil {
		tx.Rollback()
		utils.Responser.Fail(c, resultcode.DataCreateFail, err)
		return
	}
	for _, t := range cirHistoryValue {
		if t.Source != currentState["id"].(string) {
			costDuration := time.Since(time.Time(t.CreateTime))
			costDurationValue = int64(costDuration) / 1000 / 1000 / 1000
		}
	}

	// 添加转交历史
	_, err = tx.Insert(&order.SdOrderCirculationHistory{
		Title:        workOrderInfo.Title,
		OrderId:      workOrderInfo.Id,
		State:        currentState["label"].(string),
		Circulation:  "转交",
		Processor:    currentUserInfo.Name,
		ProcessorId:  c.Values().Get("user").(user.SdUser).Id,
		Remarks:      fmt.Sprintf("此阶段负责人已转交给《%v》", userInfo.Name),
		Status:       2, // 其他
		CostDuration: costDurationValue,
	})
	if err != nil {
		utils.Responser.Fail(c, resultcode.DataCreateFail, err)
		return
	}

	tx.Commit()

	utils.Responser.Ok(c)
}

//
//// 催办工单
//func UrgeWorkOrder(c iris.Context) {
//	var (
//		workOrderInfo  process.WorkOrderInfo
//		sendToUserList []system.SysUser
//		stateList      []interface{}
//		userInfo       system.SysUser
//	)
//	workOrderId := c.DefaultQuery("workOrderId", "")
//	if workOrderId == "" {
//		app.Error(c, -1, errors.New("参数不正确，缺失workOrderId"), "")
//		return
//	}
//
//	// 查询工单数据
//	err := orm.Eloquent.Model(&process.WorkOrderInfo{}).Where("id = ?", workOrderId).Find(&workOrderInfo).Error
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("查询工单信息失败，%v", err.Error()))
//		return
//	}
//
//	// 确认是否可以催办
//	if workOrderInfo.UrgeLastTime != 0 && (int(time.Now().Unix())-workOrderInfo.UrgeLastTime) < 600 {
//		app.Error(c, -1, errors.New("十分钟内无法多次催办工单。"), "")
//		return
//	}
//
//	// 获取当前工单处理人信息
//	err = json.Unmarshal(workOrderInfo.State, &stateList)
//	if err != nil {
//		app.Error(c, -1, err, "")
//		return
//	}
//	sendToUserList, err = service.GetPrincipalUserInfo(stateList, workOrderInfo.Creator)
//
//	// 查询创建人信息
//	err = orm.Eloquent.Model(&system.SysUser{}).Where("user_id = ?", workOrderInfo.Creator).Find(&userInfo).Error
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("创建人信息查询失败，%v", err.Error()))
//		return
//	}
//
//	// 发送催办提醒
//	bodyData := notify.BodyData{
//		SendTo: map[string]interface{}{
//			"userList": sendToUserList,
//		},
//		Subject:     "您被催办工单了，请及时处理。",
//		Description: "您有一条待办工单，请及时处理，工单描述如下",
//		Classify:    []int{1}, // todo 1 表示邮箱，后续添加了其他的在重新补充
//		ProcessId:   workOrderInfo.Process,
//		Id:          workOrderInfo.Id,
//		Title:       workOrderInfo.Title,
//		Creator:     userInfo.NickName,
//		Priority:    workOrderInfo.Priority,
//		CreatedAt:   workOrderInfo.CreatedAt.Format("2006-01-02 15:04:05"),
//	}
//	err = bodyData.SendNotify()
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("催办提醒发送失败，%v", err.Error()))
//		return
//	}
//
//	// 更新数据库
//	err = orm.Eloquent.Model(&process.WorkOrderInfo{}).Where("id = ?", workOrderInfo.Id).Updates(map[string]interface{}{
//		"urge_count":     workOrderInfo.UrgeCount + 1,
//		"urge_last_time": int(time.Now().Unix()),
//	}).Error
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("更新催办信息失败，%v", err.Error()))
//		return
//	}
//
//	app.OK(c, "", "")
//}
//
//// 主动处理
//func ActiveOrder(c iris.Context) {
//	var (
//		workOrderId string
//		err         error
//		stateValue  []struct {
//			ID            string `json:"id"`
//			Label         string `json:"label"`
//			ProcessMethod string `json:"process_method"`
//			Processor     []int  `json:"processor"`
//		}
//		stateValueByte []byte
//	)
//
//	workOrderId = c.Param("id")
//
//	err = c.ShouldBind(&stateValue)
//	if err != nil {
//		app.Error(c, -1, err, "")
//		return
//	}
//
//	stateValueByte, err = json.Marshal(stateValue)
//	if err != nil {
//		app.Error(c, -1, fmt.Errorf("转byte失败，%v", err.Error()), "")
//		return
//	}
//
//	err = orm.Eloquent.Model(&process.WorkOrderInfo{}).
//		Where("id = ?", workOrderId).
//		Update("state", stateValueByte).Error
//	if err != nil {
//		app.Error(c, -1, fmt.Errorf("接单失败，%v", err.Error()), "")
//		return
//	}
//
//	app.OK(c, "", "接单成功，请及时处理")
//}
//

//
// 删除工单
func DeleteWorkOrder(c iris.Context) {

	workOrderId := c.URLParam("id")
	effect, err := global.GVA_DB.Where("id = ?", workOrderId).Delete(new(order.SdOrder))
	if effect <= 0 || err != nil {
		utils.Responser.Fail(c, resultcode.DataDeleteFail, err)
		return
	}

	utils.Responser.Ok(c)
}

//
//// 重开工单
//func ReopenWorkOrder(c *gin.Context) {
//	var (
//		err           error
//		id            string
//		workOrder     process.WorkOrderInfo
//		processInfo   process.Info
//		structure     map[string]interface{}
//		startId       string
//		label         string
//		jsonState     []byte
//		relatedPerson []byte
//		newWorkOrder  process.WorkOrderInfo
//		workOrderData []*process.TplData
//	)
//
//	id = c.Param("id")
//
//	// 查询当前ID的工单信息
//	err = orm.Eloquent.Find(&workOrder, id).Error
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("查询工单信息失败, %s", err.Error()))
//		return
//	}
//
//	// 创建新的工单
//	err = orm.Eloquent.Find(&processInfo, workOrder.Process).Error
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("查询流程信息失败, %s", err.Error()))
//		return
//	}
//	err = json.Unmarshal(processInfo.Structure, &structure)
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("Json序列化失败, %s", err.Error()))
//		return
//	}
//	for _, node := range structure["nodes"].([]interface{}) {
//		if node.(map[string]interface{})["clazz"] == "start" {
//			startId = node.(map[string]interface{})["id"].(string)
//			label = node.(map[string]interface{})["label"].(string)
//		}
//	}
//
//	state := []map[string]interface{}{
//		{"id": startId, "label": label, "processor": []int{tools.GetUserId(c)}, "process_method": "person"},
//	}
//	jsonState, err = json.Marshal(state)
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("Json序列化失败, %s", err.Error()))
//		return
//	}
//
//	relatedPerson, err = json.Marshal([]int{tools.GetUserId(c)})
//	if err != nil {
//		app.Error(c, -1, err, fmt.Sprintf("Json序列化失败, %s", err.Error()))
//		return
//	}
//
//	tx := orm.Eloquent.Begin()
//
//	newWorkOrder = process.WorkOrderInfo{
//		Title:         workOrder.Title,
//		Priority:      workOrder.Priority,
//		Process:       workOrder.Process,
//		Classify:      workOrder.Classify,
//		State:         jsonState,
//		RelatedPerson: relatedPerson,
//		Creator:       tools.GetUserId(c),
//	}
//	err = tx.Create(&newWorkOrder).Error
//	if err != nil {
//		tx.Rollback()
//		app.Error(c, -1, err, fmt.Sprintf("新建工单失败, %s", err.Error()))
//		return
//	}
//
//	// 查询工单数据
//	err = orm.Eloquent.Model(&process.TplData{}).Where("work_order = ?", id).Find(&workOrderData).Error
//	if err != nil {
//		tx.Rollback()
//		app.Error(c, -1, err, fmt.Sprintf("查询工单数据失败, %s", err.Error()))
//		return
//	}
//
//	for _, d := range workOrderData {
//		d.WorkOrder = newWorkOrder.Id
//		d.Id = 0
//		err = tx.Create(d).Error
//		if err != nil {
//			tx.Rollback()
//			app.Error(c, -1, err, fmt.Sprintf("创建工单数据失败, %s", err.Error()))
//			return
//		}
//	}
//
//	tx.Commit()
//
//	app.OK(c, nil, "")
//}
//func JudgeUserAuthority(c iris.Context, workOrderId int, currentState string) (status bool, err error) {
//	/*
//		person 人员
//		persongroup 人员组
//		department 部门
//		variable 变量
//	*/
//	var (
//		//userDept          department.SdDepartment
//		workOrderInfo     order.SdOrder
//		userInfo          user.SdUser
//		cirHistoryList    []order.SdOrderCirculationHistory
//		stateValue        map[string]interface{}
//		processInfo       workflow.SdWorkflow
//		processState      order2.ProcessState
//		currentStateList  []map[string]interface{}
//		currentStateValue map[string]interface{}
//		currentUserInfo   user.SdUser
//	)
//	// 获取工单信息
//	err = global.GVA_DB.Id(workOrderId).Find(&workOrderInfo)
//	if err != nil {
//		return
//	}
//
//	// 获取流程信息
//	err = global.GVA_DB.Id(workOrderInfo.WorkflowId).Find(&processInfo)
//	if err != nil {
//		return
//	}
//
//	if processInfo.Structure != nil && len(processInfo.Structure) > 0 {
//		err = json.Unmarshal(processInfo.Structure, &processState.Structure)
//		if err != nil {
//			return
//		}
//	}
//
//	stateValue, err = processState.GetNode(currentState)
//	if err != nil {
//		return
//	}
//
//	err = json.Unmarshal(workOrderInfo.State, &currentStateList)
//	if err != nil {
//		return
//	}
//
//	for _, v := range currentStateList {
//		if v["id"].(string) == currentState {
//			currentStateValue = v
//			break
//		}
//	}
//
//	// 获取当前用户信息
//	err = global.GVA_DB.Id(c.Values().Get("user").(user.SdUser).Id).Find(&currentUserInfo)
//	if err != nil {
//		return
//	}
//
//	// 会签
//	if currentStateValue["processor"] != nil && len(currentStateValue["processor"].([]interface{})) >= 1 {
//		if isCounterSign, ok := stateValue["isCounterSign"]; ok {
//			if isCounterSign.(bool) {
//				err = global.GVA_DB.Where("order_id = ?", workOrderId).Desc("id").Find(&cirHistoryList)
//				if err != nil {
//					return
//				}
//				for _, cirHistoryValue := range cirHistoryList {
//					if cirHistoryValue.Source != stateValue["id"] {
//						break
//					} else if cirHistoryValue.Source == stateValue["id"] {
//						if currentStateValue["process_method"].(string) == "person" {
//							// 验证个人会签
//							if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
//								return
//							}
//						} else if currentStateValue["process_method"].(string) == "role" {
//							// 验证角色会签
//							if stateValue["fullHandle"].(bool) {
//								if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
//									return
//								}
//							} else {
//								var roleUserInfo user.SdUser
//								err = global.GVA_DB.Id(cirHistoryValue.ProcessorId).Find(&roleUserInfo)
//								if err != nil {
//									return
//								}
//								//if roleUserInfo.RoleId == tools.GetRoleId(c) {
//								//	return
//								//}
//							}
//						} else if currentStateValue["process_method"].(string) == "department" {
//							//部门会签
//							if stateValue["fullHandle"].(bool) {
//								if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
//									return
//								}
//							} else {
//								var (
//									deptUserInfo user.SdUser
//								)
//								err = global.GVA_DB.Id(cirHistoryValue.ProcessorId).Find(&deptUserInfo)
//								if err != nil {
//									return
//								}
//
//								//if deptUserInfo.DeptId == currentUserInfo.DeptId {
//								//	return
//								//}
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//
//	switch currentStateValue["process_method"].(string) {
//	case "person":
//		for _, processorValue := range currentStateValue["processor"].([]interface{}) {
//			if int(processorValue.(float64)) == c.Values().Get("user").(user.SdUser).Id {
//				status = true
//			}
//		}
//	case "role":
//		for _, processorValue := range currentStateValue["processor"].([]interface{}) {
//			if int(processorValue.(float64)) == c.Values().Get("user").(user.SdUser).Id {
//				status = true
//			}
//		}
//	case "department":
//		for _, _ = range currentStateValue["processor"].([]interface{}) {
//			//if int(processorValue.(float64)) == currentUserInfo.DeptId {
//			status = true
//			//}
//		}
//	case "variable":
//		for _, p := range currentStateValue["processor"].([]interface{}) {
//			switch int(p.(float64)) {
//			case 1:
//				if workOrderInfo.UserId == c.Values().Get("user").(user.SdUser).Id {
//					status = true
//				}
//			case 2:
//				err = global.GVA_DB.Id(workOrderInfo.UserId).Find(&userInfo)
//				if err != nil {
//					return
//				}
//				//err = orm.Eloquent.Model(&userDept).Id(userInfo.DeptId).Find(&userDept).Error
//				//if err != nil {
//				//	return
//				//}
//
//				//if userDept.Leader == tools.GetUserId(c) {
//				status = true
//				//}
//			}
//		}
//	}
//	return
//}
