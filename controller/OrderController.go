package controller

import (
	"net/http"
	"project-pertama/model"
	"project-pertama/repository"
	"project-pertama/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepository repository.OrderRepositoryImpl
}

func NewOrderController(orderRepository repository.OrderRepositoryImpl) *OrderController {
	return &OrderController{
		orderRepository: orderRepository,
	}
}

// Create Order godoc
// @Summary Create Order
// @Schemes
// @Description Create Order
// @Tags order
// @Accept json
// @Produce json
// @Param request body model.OrderHandler.request true "query params"
// @Success 200 {object} model.OrderHandler.response
// @Router /order [post]
func (oc *OrderController) Create(ctx *gin.Context) {

	var order model.Order

	err := ctx.ShouldBindJSON(&order)

	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	// Check if ordered_at is null or empty
	if order.OrderedAt.IsZero() {
		// Handle case where ordered_at is not provided or empty
		var r model.Response = model.Response{
			Success: false,
			Error:   "ordered_at field is required",
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	// Check if customer_name is null or empty
	if order.CustomerName == "" {
		// Handle case where customer_name is not provided or empty
		var r model.Response = model.Response{
			Success: false,
			Error:   "customer_name field is required",
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	createOrder, err := oc.orderRepository.Create(order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createOrder, ""))
}

func (oc *OrderController) Get(ctx *gin.Context) {

	createOrder, err := oc.orderRepository.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createOrder, ""))
}

func (oc *OrderController) Update(ctx *gin.Context) {

	var order model.Order

	err := ctx.ShouldBindJSON(&order)

	// parse the id to uint
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	uid := uint(id)

	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	// Check if ordered_at is null or empty
	if order.OrderedAt.IsZero() {
		// Handle case where ordered_at is not provided or empty
		var r model.Response = model.Response{
			Success: false,
			Error:   "ordered_at field is required",
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	// Check if customer_name is null or empty
	if order.CustomerName == "" {
		// Handle case where customer_name is not provided or empty
		var r model.Response = model.Response{
			Success: false,
			Error:   "customer_name field is required",
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	createOrder, err := oc.orderRepository.Update(uid, order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createOrder, ""))

}

func (oc *OrderController) Delete(ctx *gin.Context) {

	// parse the id to uint
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	uid := uint(id)

	createOrder, err := oc.orderRepository.Delete(uid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createOrder, ""))

}
