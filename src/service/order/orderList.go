package order

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/timedecoder"
	"mayday/src/model/order"
	"mayday/src/model/user"
)

type WorkOrder struct {
	Classify int
	Context  iris.Context
}

type workOrderInfo struct {
	Id            int
	UserId        int
	WorkflowId    int
	CreateTime    timedecoder.LocalTime
	Title         string
	UrgeLastTime  timedecoder.LocalTime
	UrgeCount     int
	RelatedPerson json.RawMessage
	IsDenied      int
	IsEnd         int
	State         json.RawMessage
	IsDeleted     int
	Principals    string `json:"principals"`
	StateName     string `json:"state_name"`
	DataClassify  int    `json:"data_classify"`
	ProcessName   string `json:"process_name"`
}

func NewWorkOrder(classify int, c iris.Context) *WorkOrder {
	return &WorkOrder{
		Classify: classify,
		Context:  c,
	}
}

func (w *WorkOrder) PureWorkOrderList() (result interface{}, err error) {
	var (
		//workOrderInfoList []workOrderInfo
		processorInfo user.SdUser
	)

	personSelectValue := "(JSON_CONTAINS(sd_order.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(sd_order.state, JSON_OBJECT('process_method', 'person')))"
	//roleSelectValue := "(JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('process_method', 'role')))"
	//departmentSelectValue := "(JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('processor', %v)) and JSON_CONTAINS(p_work_order_info.state, JSON_OBJECT('process_method', 'department')))"

	startTime := w.Context.FormValue("startTime")
	endTime := w.Context.FormValue("endTime")
	isEnd := w.Context.FormValue("isEnd")
	processor := w.Context.FormValue("processor")
	//priority := w.context.FormValue("priority")
	creator := w.Context.FormValue("creator")
	db := global.GVA_DB.NewSession()
	db.Begin()
	if startTime != "" {
		db = db.Where("create_time >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("create_time <= ?", endTime)
	}
	if isEnd != "" {
		db = db.Where("is_end = ?", isEnd)
	}
	if creator != "" {
		db = db.Where("user_id = ?", creator)
	}
	if processor != "" && w.Classify != 1 {
		err = global.GVA_DB.Id(processor).Find(&processorInfo)
		if err != nil {
			return
		}
		//or %v or %v
		db = db.Where(fmt.Sprintf("(%v) and is_end = 0",
			fmt.Sprintf(personSelectValue, processorInfo.Id),
			//fmt.Sprintf(roleSelectValue, processorInfo.RoleId),
			//fmt.Sprintf(departmentSelectValue, processorInfo.DeptId),
		))
	}
	//if priority != "" {
	//	db = db.Where("p_work_order_info.priority = ?", priority)
	//}

	// 获取当前用户信息
	switch w.Classify {
	case 1:
		// 待办工单
		// 1. 个人
		personSelect := fmt.Sprintf(personSelectValue, w.Context.Values().Get("user").(user.SdUser).Id)

		//// 2. 角色
		//roleSelect := fmt.Sprintf(roleSelectValue, w.context.Values().Get("user").(user.SdUser).Id)
		//
		//// 3. 部门
		//var userInfo system.SysUser
		//err = orm.Eloquent.Model(&system.SysUser{}).
		//	Where("user_id = ?", tools.GetUserId(w.GinObj)).
		//	Find(&userInfo).Error
		//if err != nil {
		//	return
		//}
		//departmentSelect := fmt.Sprintf(departmentSelectValue, userInfo.DeptId)

		// 4. 变量会转成个人数据
		//db = db.Where(fmt.Sprintf("(%v or %v or %v or %v) and is_end = 0", personSelect, personGroupSelect, departmentSelect, variableSelect))
		db = db.Where(fmt.Sprintf("(%v) and is_end = 0", personSelect))
	case 2:
		// 我创建的
		db = db.Where("user_id = ?", w.Context.Values().Get("user").(user.SdUser).Id)
	case 3:
		// 我相关的
		db = db.Where(fmt.Sprintf("JSON_CONTAINS(related_person, '%v')", w.Context.Values().Get("user").(user.SdUser).Id))
	case 4:
	// 所有工单
	default:
		return nil, fmt.Errorf("请确认查询的数据类型是否正确")
	}

	//db = db.Join("left join p_process_info on p_work_order_info.process = p_process_info.id").
	//	Select("p_work_order_info.*, p_process_info.name as process_name")

	//result, err = pagination.Paging(&pagination.Param{
	//	C:  w.Context,
	//	DB: db,
	//}, &workOrderInfoList, map[string]map[string]interface{}{}, "p_process_info")
	var sdOrder []order.SdOrder
	err = db.Find(&sdOrder)
	if err != nil {
		global.GVA_LOG.Info("查询工单列表失败")
		return nil, err
	}
	result = sdOrder
	return
}

func (w *WorkOrder) WorkOrderList() (result interface{}, err error) {

	var (
		principals        string
		StateList         []map[string]interface{}
		workOrderInfoList []workOrderInfo
		minusTotal        int
	)

	result, err = w.PureWorkOrderList()
	if err != nil {
		return nil, err
	}

	for i, v := range result.([]order.SdOrder) {
		var (
			stateName    string
			structResult map[string]interface{}
			authStatus   bool
		)
		err = json.Unmarshal(v.State, &StateList)
		if err != nil {
			err = fmt.Errorf("json反序列化失败，%v", err.Error())
			return
		}
		if len(StateList) != 0 {
			// 仅待办工单需要验证
			if w.Classify == 1 {
				structResult, err = MakeProcessStructure(w.Context, v.WorkflowId, v.Id)
				if err != nil {
					return
				}

				authStatus, err = JudgeUserAuthority(w.Context, v.Id, structResult["workOrder"].(WorkOrderData).CurrentState)
				if err != nil {
					return
				}
				if !authStatus {
					minusTotal += 1
					continue
				}
			} else {
				authStatus = true
			}

			processorList := make([]int, 0)
			if len(StateList) > 1 {
				for _, s := range StateList {
					for _, p := range s["processor"].([]interface{}) {
						if int(p.(float64)) == w.Context.Values().Get("user").(user.SdUser).Id {
							processorList = append(processorList, int(p.(float64)))
						}
					}
					if len(processorList) > 0 {
						stateName = s["label"].(string)
						break
					}
				}
			}
			if len(processorList) == 0 {
				for _, v := range StateList[0]["processor"].([]interface{}) {
					processorList = append(processorList, int(v.(float64)))
				}
				stateName = StateList[0]["label"].(string)
			}
			principals, err = GetPrincipal(processorList, StateList[0]["process_method"].(string))
			if err != nil {
				err = fmt.Errorf("查询处理人名称失败，%v", err.Error())
				return
			}
		}

		workOrderDetails := result.([]order.SdOrder)
		workOrderDetail := struct {
			Id            int
			UserId        int
			WorkflowId    int
			CreateTime    timedecoder.LocalTime
			Title         string
			UrgeLastTime  timedecoder.LocalTime
			UrgeCount     int
			RelatedPerson json.RawMessage
			IsDenied      int
			IsEnd         int
			State         json.RawMessage
			IsDeleted     int
			Principals    string `json:"principals"`
			StateName     string `json:"state_name"`
			DataClassify  int    `json:"data_classify"`
			ProcessName   string `json:"process_name"`
		}{
			Id:            workOrderDetails[i].Id,
			UserId:        workOrderDetails[i].UserId,
			WorkflowId:    workOrderDetails[i].WorkflowId,
			CreateTime:    workOrderDetails[i].CreateTime,
			Title:         workOrderDetails[i].Title,
			UrgeLastTime:  workOrderDetails[i].UrgeLastTime,
			UrgeCount:     workOrderDetails[i].UrgeCount,
			RelatedPerson: workOrderDetails[i].RelatedPerson,
			IsDenied:      workOrderDetails[i].IsDenied,
			IsEnd:         workOrderDetails[i].IsEnd,
			State:         workOrderDetails[i].State,
			IsDeleted:     workOrderDetails[i].IsDeleted,
			Principals:    principals,
			StateName:     stateName,
		}

		//workOrderDetails[i].DataClassify = v.Classify
		if authStatus {
			workOrderInfoList = append(workOrderInfoList, workOrderDetail)
		}
	}

	//result.(*pagination.Paginator).Data = &workOrderInfoList
	//result.(*pagination.Paginator).TotalCount -= minusTotal

	return workOrderInfoList, nil
}

//func Paging(p *Param, result interface{}, args ...interface{}) (*Paginator, error) {
//	var (
//		param     ListRequest
//		paginator Paginator
//		count     int
//		offset    int
//		tableName string
//	)
//
//	if err := p.C.Bind(&param); err != nil {
//		logger.Errorf("参数绑定失败，错误：%v", err)
//		return nil, err
//	}
//
//	db := p.DB
//
//	if p.ShowSQL {
//		db = db.Debug()
//	}
//
//	if param.Page < 1 {
//		param.Page = 1
//	}
//
//	if param.PerPage == 0 {
//		param.PerPage = 10
//	}
//
//	if param.Sort == 0 || param.Sort == -1 {
//		db = db.Order("id desc")
//	}
//
//	if len(args) > 1 {
//		tableName = fmt.Sprintf("`%s`.", args[1].(string))
//	}
//
//	if len(args) > 0 {
//		for paramType, paramsValue := range args[0].(map[string]map[string]interface{}) {
//			if paramType == "like" {
//				for key, value := range paramsValue {
//					db = db.Where(fmt.Sprintf("%v%v like ?", tableName, key), fmt.Sprintf("%%%v%%", value))
//				}
//			} else if paramType == "equal" {
//				for key, value := range paramsValue {
//					db = db.Where(fmt.Sprintf("%v%v = ?", tableName, key), value)
//				}
//			}
//		}
//	}
//
//	done := make(chan bool, 1)
//
//	go countRecords(db, result, done, &count)
//
//	if param.Page == 1 {
//		offset = 0
//	} else {
//		offset = (param.Page - 1) * param.PerPage
//	}
//	err := db.Limit(param.PerPage).Offset(offset).Scan(result).Error
//	if err != nil {
//		logger.Errorf("数据查询失败，错误：%v", err)
//		return nil, err
//	}
//	<-done
//
//	paginator.TotalCount = count
//	paginator.Data = result
//	paginator.Page = param.Page
//	paginator.PerPage = param.PerPage
//	paginator.TotalPage = int(math.Ceil(float64(count) / float64(param.PerPage)))
//
//	return &paginator, nil
//}
