package order

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/model/common/timedecoder"
	"mayday/src/model/order"
	"mayday/src/model/user"
	"mayday/src/model/workflow"
	"time"

	userModel "mayday/src/model/user"
	//"mayday/src/utils/responser"
	//"mayday/src/utils/responser/vo"
	//"mayday/middleware/jwts"
)

func CreateOrderService(ctx iris.Context, user *user.SdUser) (err error) {
	var (
		//taskList       []string
		stateList     []interface{}       //节点列表
		userInfo      userModel.SdUser    //请求用户信息
		variableValue []interface{}       //节点列表
		workflowValue workflow.SdWorkflow //流程信息
		//sendToUserList []model.SdUser //要通知的人的列表
		//noticeList     []int            //流程的通知列表
		//handle       Handle
		processState ProcessState //流程结构
		//condExprStatus bool
		//tpl []byte
		//sourceEdges    []map[string]interface{}
		//targetEdges    []map[string]interface{}
		//currentNode    map[string]interface{} //流程的开始节点（不知道为什么是个列表）
		workOrderValue struct { //请求的全部数据
			order.SdOrder
			Tpls        map[string][]interface{} `json:"tpls"` //表单结构和数据
			SourceState string                   `json:"source_state"`
			Tasks       json.RawMessage          `json:"tasks"` ////////
			Source      string                   `json:"source"`
			IsExecTask  bool                     `json:"is_exec_task"`
		}
		//paramsValue struct {
		//Id       int           `json:"id"`
		//Title    string        `json:"title"`
		//Priority int           `json:"priority"`
		//FormData []interface{} `json:"form_data"`
		//}
	)
	//获取请求的全部数据
	err = ctx.ReadJSON(&workOrderValue)
	if err != nil {
		global.GVA_LOG.Warn("数据接收失败", zap.Error(err))
		return
	}
	//设置参与人
	relatedPerson, err := json.Marshal([]int{user.Id})
	if err != nil {
		global.GVA_LOG.Warn("序列化参与人信息失败", zap.Error(err))
		return
	}
	// 获取节点数据 variableValue：节点列表
	err = json.Unmarshal(workOrderValue.State, &variableValue)
	if err != nil {
		global.GVA_LOG.Warn("获取节点列表失败", zap.Error(err))
		return
	}

	//检查节点处理人的ID是否存在数据库
	err = GetVariableValue(variableValue, user.Id)
	if err != nil {
		global.GVA_LOG.Warn("获取处理人变量值失败", zap.Error(err))
		return
	}

	// 创建工单数据    tx:数据库链接对象
	tx := global.GVA_DB.NewSession()
	defer tx.Close()
	err = tx.Begin()
	if err != nil {
		tx.Rollback()
		global.GVA_LOG.Warn("数据库事务创建失败", zap.Error(err))
		return
	}

	// 从数据库查询流程信息
	has, err := tx.ID(workOrderValue.WorkflowId).Get(&workflowValue)
	if !has || err != nil {
		tx.Rollback()
		global.GVA_LOG.Warn("获取流程信息失败", zap.Error(err))
		return err
	}
	//取出流程结构
	err = json.Unmarshal(workflowValue.Structure, &processState.Structure)
	if err != nil {
		global.GVA_LOG.Warn("流程结构解析失败", zap.Error(err))
		return
	}
	//找到流程的开始节点
	//for _, node := range processState.Structure["nodes"] {
	//	if node["clazz"] == "start" {
	//		currentNode = node
	//	}
	//}
	//获取第一个节点的详细信息
	//nodeValue, err := processState.GetNode(variableValue[0].(map[string]interface{})["id"].(string))
	//if err != nil {
	//	global.GVA_LOG.Warn("获取第一个节点的详细信息失败", zap.Error(err))
	//	return
	//}

	//获取请求中的表单数据
	//for _, _ = range workOrderValue.Tpls["form_data"] {
	//	//解析每个表单数据
	//	tpl, err = json.Marshal(v)
	//	if err != nil {
	//		global.GVA_LOG.Warn("解析表单数据失败", zap.Error(err))
	//		return
	//	}
	//	//储存全部表单数据（一个个加到末尾）handle.WorkOrderData：form_data解析后
	//	handle.WorkOrderData = append(handle.WorkOrderData, tpl)
	//}
	//var nodeValue,condExprStatus,sourceEdges,targetEdges
	//-------------------------------------------------------如果当前节点是网关节点------------------------------------
	/*switch nodeValue["clazz"] {
	// 排他网关
	case "exclusiveGateway":
		var sourceEdges []map[string]interface{}
		sourceEdges, err = processState.GetEdge(nodeValue["id"].(string), "source")
		if err != nil {
			return
		}
	breakTag:
		for _, edge := range sourceEdges {
			edgeCondExpr := make([]map[string]interface{}, 0)
			err = json.Unmarshal([]byte(edge["conditionExpression"].(string)), &edgeCondExpr)
			if err != nil {
				return
			}
			for _, condExpr := range edgeCondExpr {
				// 条件判断
				condExprStatus, err = ConditionalJudgment(condExpr)
				if err != nil {
					return
				}
				if condExprStatus {
					// 进行节点跳转
					nodeValue, err = processState.GetNode(edge["target"].(string))
					if err != nil {
						return
					}

					if nodeValue["clazz"] == "userTask" || nodeValue["clazz"] == "receiveTask" {
						if nodeValue["assignValue"] == nil || nodeValue["assignType"] == "" {
							err = errors.New("处理人不能为空")
							return
						}
					}
					variableValue[0].(map[string]interface{})["id"] = nodeValue["id"].(string)
					variableValue[0].(map[string]interface{})["label"] = nodeValue["label"]
					variableValue[0].(map[string]interface{})["processor"] = nodeValue["assignValue"]
					variableValue[0].(map[string]interface{})["process_method"] = nodeValue["assignType"]
					break breakTag
				}
			}
		}
		if !condExprStatus {
			err = errors.New("所有流转均不符合条件，请确认。")
			return
		}
	case "parallelGateway":
		// 入口，判断
		sourceEdges, err = processState.GetEdge(nodeValue["id"].(string), "source")
		if err != nil {
			err = fmt.Errorf("查询流转信息失败，%v", err.Error())
			return
		}

		targetEdges, err = processState.GetEdge(nodeValue["id"].(string), "target")
		if err != nil {
			err = fmt.Errorf("查询流转信息失败，%v", err.Error())
			return
		}

		if len(sourceEdges) > 0 {
			nodeValue, err = processState.GetNode(sourceEdges[0]["target"].(string))
			if err != nil {
				return
			}
		} else {
			err = errors.New("并行网关流程不正确")
			return
		}

		if len(sourceEdges) > 1 && len(targetEdges) == 1 {
			// 入口
			variableValue = []interface{}{}
			for _, edge := range sourceEdges {
				targetStateValue, err := processState.GetNode(edge["target"].(string))
				if err != nil {
					return err
				}
				variableValue = append(variableValue, map[string]interface{}{
					"id":             edge["target"].(string),
					"label":          targetStateValue["label"],
					"processor":      targetStateValue["assignValue"],
					"process_method": targetStateValue["assignType"],
				})
			}
		} else {
			err = errors.New("并行网关流程配置不正确")
			return
		}
	}

	// 再次检查节点的处理人（如果不是网关节点应该是没用的）
	//err = GetVariableValue(variableValue, tools.GetUserId(c))
	//if err != nil {
	//	return
	//}
	//把处理过后的节点数据赋值到请求数据（修改请求数据）（如果不是网关节点应该也是没用的）
	workOrderValue.State, err = json.Marshal(variableValue)
	if err != nil {
		return
	}
	*/
	//-------------------------------如果不是网关节点应该会直接跳到这里------------------------------------
	//新建一个变量存储请求数据   对应数据库表  p_work_order_info
	var OrderInfo = order.SdOrder{
		Title:         workOrderValue.Title,
		WorkflowId:    workOrderValue.WorkflowId,
		State:         workOrderValue.State,
		RelatedPerson: relatedPerson,
		CreateTime:    timedecoder.LocalTime(time.Now()),
		UserId:        user.Id,
	}
	//数据库插入新订单的记录
	effect, err := tx.Insert(&OrderInfo)
	if effect <= 0 || err != nil {
		tx.Rollback()
		global.GVA_LOG.Warn("插入订单数据失败", zap.Error(err))
		return err
	}

	//创建工单模版关联数据
	//遍历所有表单
	for i := 0; i < len(workOrderValue.Tpls["form_structure"]); i++ {
		var (
			formDataJson      []byte //表单数据
			formStructureJson []byte //表单结构
		)
		formDataJson, err = json.Marshal(workOrderValue.Tpls["form_data"][i])
		if err != nil {
			tx.Rollback()
			global.GVA_LOG.Warn("生成json数据失败", zap.Error(err))
			return
		}
		formStructureJson, err = json.Marshal(workOrderValue.Tpls["form_structure"][i])
		if err != nil {
			tx.Rollback()
			global.GVA_LOG.Warn("生成json数据失败", zap.Error(err))
			return
		}
		//数据库存储对象 对应数据库表 p_work_order_tpl_data
		formData := order.SdOrderTable{
			OrderId:       OrderInfo.Id,
			FormStructure: formStructureJson,
			FormData:      formDataJson,
		}
		//插入
		effect, err1 := tx.Insert(formData)
		if effect <= 0 || err1 != nil {
			err = err1
			tx.Rollback()
			global.GVA_LOG.Warn("创建工单模版关联数据失败", zap.Error(err))
			return
		}
	}

	has, err1 := tx.Id(user.Id).Get(&userInfo)
	if !has || err1 != nil || userInfo.IsDeleted == 1 {
		tx.Rollback()
		err = err1
		global.GVA_LOG.Warn("数据库查询错误或用户名不存在", zap.Error(err))
		return
	}

	//当前用户昵称信息
	nameValue := userInfo.Name

	// 创建历史记录
	//获取当前节点数据并创建历史记录 对应数据库表 p_work_order_circulation_history
	err = json.Unmarshal(OrderInfo.State, &stateList)
	if err != nil {
		tx.Rollback()
		global.GVA_LOG.Warn("生成json数据失败", zap.Error(err))
		return
	}
	//插入操作
	if affect, err1 := tx.Insert(order.SdOrderCirculationHistory{
		Title:       workOrderValue.Title,
		OrderId:     OrderInfo.Id,
		State:       workOrderValue.SourceState,
		Source:      workOrderValue.Source,
		Target:      stateList[0].(map[string]interface{})["id"].(string),
		Circulation: "新建",
		Processor:   nameValue,
		CreateTime:  timedecoder.LocalTime(time.Now()),
		ProcessorId: userInfo.Id,
		Status:      2, // 其他
	}); affect <= 0 || err1 != nil {
		tx.Rollback()
		err = err1
		global.GVA_LOG.Warn("新建历史数据错误", zap.Error(err))
		return
	}

	// 更新流程提交数量统计
	if affect, err1 := tx.
		Table(new(workflow.SdWorkflow)).
		Id(workOrderValue.WorkflowId).
		Update(map[string]interface{}{"ceiling_count": workflowValue.CeilingCount + 1}); affect <= 0 || err1 != nil {
		tx.Rollback()
		err = err1
		global.GVA_LOG.Warn("更新流程统计数据失败", zap.Error(err))
		return
	}
	//数据库保存
	tx.Commit()
	return
}

func ConditionalJudgment(stateList []interface{}) (bool, error) {
	return true, nil
}

func GetVariableValue(stateList []interface{}, creator int) (err error) {
	for _, stateItem := range stateList {
		if stateItem.(map[string]interface{})["process_method"] == "variable" {
			for processorIndex, _ := range stateItem.(map[string]interface{})["processor"].([]interface{}) {
				stateItem.(map[string]interface{})["processor"].([]interface{})[processorIndex] = creator
			}
			stateItem.(map[string]interface{})["process_method"] = "person"
		}
	}
	return
}
