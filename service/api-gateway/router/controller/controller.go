package controller

import (
	"log"
	grpcclient "order-service-platform/service/api-gateway/grpc"
	"order-service-platform/service/api-gateway/model"
	"order-service-platform/service/api-gateway/router/base_controller"

	"order-service-platform/errcode"
	"order-service-platform/proto/proto/pb"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	base_controller.BaseController
	// ES *elasticsearch.Client
}

func NewController() *Controller {
	// cfg := elasticsearch.Config{
	// 	Addresses: []string{"http://elasticsearch:9200"}, // docker-compose service 名稱
	// }
	// es, err := elasticsearch.NewClient(cfg)
	// if err != nil {
	// 	panic(err)
	// }
	return &Controller{
		// ES: es,
	}
}

func (ctrl *Controller) GetOrder(c *gin.Context) {
	ctrl.JsonResponse(c, errcode.Success, nil)
}

// order
//
//	@Router			/v1/order [get]
//	@Tags			order
//	@Summary		query chess rtp model
//	@Description	query chess rtp model
//	@Security		BearerAuth
//	@Accept			application/json
//	@Produce		application/json
//	@Param			game_id		query		int32	false	"game id 遊戲ID"
//	@Param			model		query		int32	false	"model 模板"
//	@Param			page_index	query		int32	ture	"page index"
//	@Param			page_size	query		int32	ture	"page size"
//	@Success		200			{object}	errorcode.Resp
func (ctrl *Controller) Order(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.JsonResponse(c, errcode.CommonConvertError, err)
		return
	}

	_, err := grpcclient.OrderClient.Order(c, &pb.OrderRequest{
		UserId:   req.UserID,
		EventId:  req.EventID,
		Quantity: req.Quantity,
	})
	if err != nil {
		log.Printf("RPC 呼叫失敗: %v", err)
		ctrl.JsonResponse(c, errcode.CommonGRPCError, err)
	}

	ctrl.JsonResponse(c, errcode.Success, nil)
}
