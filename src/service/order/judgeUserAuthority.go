package order

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/order"
	"mayday/src/model/user"
	"mayday/src/model/workflow"
)

func JudgeUserAuthority(c iris.Context, workOrderId int, currentState string) (status bool, err error) {
	/*
		person 人员
		persongroup 人员组
		department 部门
		variable 变量
	*/
	var (
		//userDept          department.SdDepartment
		workOrderInfo     order.SdOrder
		userInfo          user.SdUser
		cirHistoryList    []order.SdOrderCirculationHistory
		stateValue        map[string]interface{}
		processInfo       workflow.SdWorkflow
		processState      ProcessState
		currentStateList  []map[string]interface{}
		currentStateValue map[string]interface{}
		currentUserInfo   user.SdUser
	)
	// 获取工单信息
	has, err1 := global.GVA_DB.Id(workOrderId).Get(&workOrderInfo)
	if !has || err1 != nil {
		err = err1
		return
	}

	// 获取流程信息
	has, err = global.GVA_DB.Id(workOrderInfo.WorkflowId).Get(&processInfo)
	if !has || err != nil {
		return
	}

	if processInfo.Structure != nil && len(processInfo.Structure) > 0 {
		err = json.Unmarshal(processInfo.Structure, &processState.Structure)
		if err != nil {
			return
		}
	}

	stateValue, err = processState.GetNode(currentState)
	if err != nil {
		return
	}

	err = json.Unmarshal(workOrderInfo.State, &currentStateList)
	if err != nil {
		return
	}

	for _, v := range currentStateList {
		if v["id"].(string) == currentState {
			currentStateValue = v
			break
		}
	}

	// 获取当前用户信息
	has, err = global.GVA_DB.Id(c.Values().Get("user").(user.SdUser).Id).Get(&currentUserInfo)
	if !has || err != nil {
		return
	}

	// 会签
	if currentStateValue["processor"] != nil && len(currentStateValue["processor"].([]interface{})) >= 1 {
		if isCounterSign, ok := stateValue["isCounterSign"]; ok {
			if isCounterSign.(bool) {
				err = global.GVA_DB.Where("order_id = ?", workOrderId).Desc("id").Find(&cirHistoryList)
				if err != nil {
					return
				}
				for _, cirHistoryValue := range cirHistoryList {
					if cirHistoryValue.Source != stateValue["id"] {
						break
					} else if cirHistoryValue.Source == stateValue["id"] {
						if currentStateValue["process_method"].(string) == "person" {
							// 验证个人会签
							if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
								return
							}
						} else if currentStateValue["process_method"].(string) == "role" {
							// 验证角色会签
							if stateValue["fullHandle"].(bool) {
								if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
									return
								}
							} else {
								var roleUserInfo user.SdUser
								err = global.GVA_DB.Id(cirHistoryValue.ProcessorId).Find(&roleUserInfo)
								if err != nil {
									return
								}
								//if roleUserInfo.RoleId == tools.GetRoleId(c) {
								//	return
								//}
							}
						} else if currentStateValue["process_method"].(string) == "department" {
							//部门会签
							if stateValue["fullHandle"].(bool) {
								if cirHistoryValue.ProcessorId == c.Values().Get("user").(user.SdUser).Id {
									return
								}
							} else {
								var (
									deptUserInfo user.SdUser
								)
								err = global.GVA_DB.Id(cirHistoryValue.ProcessorId).Find(&deptUserInfo)
								if err != nil {
									return
								}

								//if deptUserInfo.DeptId == currentUserInfo.DeptId {
								//	return
								//}
							}
						}
					}
				}
			}
		}
	}

	switch currentStateValue["process_method"].(string) {
	case "person":
		for _, processorValue := range currentStateValue["processor"].([]interface{}) {
			if int(processorValue.(float64)) == c.Values().Get("user").(user.SdUser).Id {
				status = true
			}
		}
	case "role":
		for _, processorValue := range currentStateValue["processor"].([]interface{}) {
			if int(processorValue.(float64)) == c.Values().Get("user").(user.SdUser).Id {
				status = true
			}
		}
	case "department":
		for _, _ = range currentStateValue["processor"].([]interface{}) {
			//if int(processorValue.(float64)) == currentUserInfo.DeptId {
			status = true
			//}
		}
	case "variable":
		for _, p := range currentStateValue["processor"].([]interface{}) {
			switch int(p.(float64)) {
			case 1:
				if workOrderInfo.UserId == c.Values().Get("user").(user.SdUser).Id {
					status = true
				}
			case 2:
				has, err = global.GVA_DB.Id(workOrderInfo.UserId).Get(&userInfo)
				if !has || err != nil {
					return
				}
				//err = orm.Eloquent.Model(&userDept).Id(userInfo.DeptId).Find(&userDept).Error
				//if err != nil {
				//	return
				//}

				//if userDept.Leader == tools.GetUserId(c) {
				status = true
				//}
			}
		}
	}
	return
}
