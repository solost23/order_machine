package create_order

import (
	"context"
	"errors"
	"fmt"
	"github.com/solost23/go_interface/gen_go/order_machine"
	"gorm.io/gorm"
	"order_machine/internal/models"
	"order_machine/internal/service/base"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *order_machine.CreateOrderRequest) (reply *order_machine.CreateOrderResponse, err error) {
	// 查看课程是否存在
	// 查看课程库存量是否满足
	db := a.GetMysqlConnect()
	sqlCourses, err := (&models.Course{}).WhereAll(db, "id IN ?", request.GetCourseIds())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	courseIdToInfoMaps := make(map[uint]*models.Course, len(sqlCourses))
	for _, sqlCourse := range sqlCourses {
		courseIdToInfoMaps[sqlCourse.ID] = sqlCourse
	}
	for index, courseId := range request.GetCourseIds() {
		courseInfo, ok := courseIdToInfoMaps[uint(courseId)]
		if !ok {
			return nil, errors.New(fmt.Sprintf("课程ID[%d]不存在，请求参数错误", courseId))
		}
		// 检查当前物品库存量
		if courseInfo.Stock < uint64(request.GetCourseNums()[index]) {
			return nil, errors.New(fmt.Sprintf("课程ID[%d]库存量不足，请求参数错误", courseId))
		}
	}
	// 数据校验都正确，生成订单
	// 减小库存
	// 生成订单信息
	// 生成关系
	begin := db.Begin()
	var totalAmount float64 = 0
	for index, courseId := range request.GetCourseIds() {
		courseInfo := courseIdToInfoMaps[uint(courseId)]
		value := map[string]interface{}{
			"stock": courseInfo.Stock - uint64(request.GetCourseNums()[index]),
		}
		err = (&models.Course{}).Update(begin, value, "id = ?", courseId)
		if err != nil {
			begin.Rollback()
			return nil, err
		}
		// 统计订单金额
		totalAmount = totalAmount + courseInfo.Amount*float64(request.GetCourseNums()[index])
	}
	order := &models.Order{
		CreatorBase: models.CreatorBase{
			CreatorId: uint(request.GetHeader().GetOperatorUid()),
		},
		// todo: 优化成从枚举值写入
		Status:          10,
		EstimatedAmount: totalAmount,
		RealAmount:      totalAmount,
		Source:          uint(request.GetOrderSource()),
	}
	err = order.Insert(begin)
	if err != nil {
		begin.Rollback()
		return nil, err
	}
	for _, courseId := range request.GetCourseIds() {
		err = (&models.OrderCourse{
			CreatorBase: models.CreatorBase{
				CreatorId: uint(request.GetHeader().GetOperatorUid()),
			},
			OrderId:  order.ID,
			CourseId: uint(courseId),
		}).Insert(begin)
		if err != nil {
			begin.Rollback()
			return nil, err
		}
	}
	begin.Commit()
	reply = &order_machine.CreateOrderResponse{
		OrderId: uint32(order.ID),
	}
	return reply, nil
}
