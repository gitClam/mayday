package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/model/order"
	UserModel "mayday/src/model/user"
	"mayday/src/model/workflow"
	"strconv"
)

type WorkOrderData struct {
	order.SdOrder
	CurrentState string `json:"current_state"`
}

func MakeProcessStructure(c iris.Context, processId int, workOrderId int) (result map[string]interface{}, err error) {
	var (
		processValue            workflow.SdWorkflow
		processStructureDetails map[string]interface{}
		processNode             []map[string]interface{}
		tplDetails              []*workflow.SdTable
		workOrderInfo           order.SdOrder
		workOrderTpls           []*order.SdOrderTable
		workOrderHistory        []*order.SdOrderCirculationHistory
		stateList               []map[string]interface{}
	)

	err = global.GVA_DB.Id(processId).Find(&processValue)

	if processValue.Structure != nil && len(processValue.Structure) > 0 {
		byteData, err1 := processValue.Structure.MarshalJSON()
		if err1 != nil {
			err = err1
			global.GVA_LOG.Error("json转byte失败，%v", zap.Error(err))
			return
		}
		err = json.Unmarshal(byteData, &processStructureDetails)
		if err != nil {
			global.GVA_LOG.Error("json转map失败，%v", zap.Error(err))
			return
		}

		// 排序，使用冒泡
		p := processStructureDetails["nodes"].([]interface{})
		if len(p) > 1 {
			for i := 0; i < len(p); i++ {
				for j := 1; j < len(p)-i; j++ {
					if p[j].(map[string]interface{})["sort"] == nil || p[j-1].(map[string]interface{})["sort"] == nil {
						return nil, errors.New("流程未定义顺序属性，请确认")
					}
					leftInt, _ := strconv.Atoi(p[j].(map[string]interface{})["sort"].(string))
					rightInt, _ := strconv.Atoi(p[j-1].(map[string]interface{})["sort"].(string))
					if leftInt < rightInt {
						p[j], p[j-1] = p[j-1], p[j]
					}
				}
			}
			for _, node := range processStructureDetails["nodes"].([]interface{}) {
				processNode = append(processNode, node.(map[string]interface{}))
			}
		} else {
			processNode = processStructureDetails["nodes"].([]map[string]interface{})
		}
	}

	processValue.Structure = nil
	result = map[string]interface{}{
		"process": processValue,
		"nodes":   processNode,
		"edges":   processStructureDetails["edges"],
	}

	// 获取历史记录
	err = global.GVA_DB.Where("order_id = ?", workOrderId).OrderBy("id desc").Find(&workOrderHistory)
	if err != nil {
		return
	}
	result["circulationHistory"] = workOrderHistory

	if workOrderId == 0 {
		// 查询流程模版
		var tplIdList []int
		byteData, err1 := processValue.Tables.MarshalJSON()
		if err1 != nil {
			err = err1
			global.GVA_LOG.Error("json转byte失败，%v", zap.Error(err))
			return
		}
		err = json.Unmarshal(byteData, &tplIdList)
		if err != nil {
			err = fmt.Errorf("json转map失败，%v", err.Error())
			return
		}
		err = global.GVA_DB.In("id", tplIdList).Find(&tplDetails)
		if err != nil {
			err = fmt.Errorf("查询模版失败，%v", err.Error())
			return
		}
		result["tpls"] = tplDetails
	} else {
		// 查询工单信息
		has, err := global.GVA_DB.Where("id = ?", workOrderId).Get(&workOrderInfo)
		if !has || err != nil {
			return nil, fmt.Errorf("查询order数据失败", err)
		}
		// 获取当前节点
		err = json.Unmarshal(workOrderInfo.State, &stateList)
		if err != nil {
			err = fmt.Errorf("序列化节点列表失败，%v", err.Error())
			return nil, err
		}
		if len(stateList) == 0 {
			err = errors.New("当前工单没有下一节点数据")
			return nil, err
		}

		// 整理需要并行处理的数据
		if len(stateList) > 1 {
		continueHistoryTag:
			for _, v := range workOrderHistory {
				status := false
				for i, s := range stateList {
					if v.Source == s["id"].(string) && v.Target != "" {
						status = true
						stateList = append(stateList[:i], stateList[i+1:]...)
						continue continueHistoryTag
					}
				}
				if !status {
					break
				}
			}
		}

		if len(stateList) > 0 {
		breakStateTag:
			for _, stateValue := range stateList {
				if processStructureDetails["nodes"] != nil {
					for _, processNodeValue := range processStructureDetails["nodes"].([]interface{}) {
						if stateValue["id"].(string) == processNodeValue.(map[string]interface{})["id"] {
							if _, ok := stateValue["processor"]; ok {
								for _, userId := range stateValue["processor"].([]interface{}) {
									if int(userId.(float64)) == c.Values().Get("user").(UserModel.SdUser).Id {
										workOrderInfo.CurrentState = stateValue["id"].(string)
										break breakStateTag
									}
								}
							} else {
								err = errors.New("未查询到对应的处理人字段，请确认。")
								return nil, err
							}
						}
					}
				}
			}

			if workOrderInfo.CurrentState == "" {
				workOrderInfo.CurrentState = stateList[0]["id"].(string)
			}
		}

		result["workOrder"] = workOrderInfo

		// 查询工单表单数据
		err = global.GVA_DB.Where("order_history_id = ?", workOrderId).Find(&workOrderTpls)
		if err != nil {
			return nil, err
		}
		result["tpls"] = workOrderTpls
	}
	return result, nil
}
