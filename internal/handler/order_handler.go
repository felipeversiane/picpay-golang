package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	"github.com/felipeversiane/picpay-golang.git/config/validation"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/felipeversiane/picpay-golang.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(
	orderService service.OrderService,
) OrderHandler {
	return &orderHandler{
		orderService,
	}
}

type OrderHandler interface {
	InsertOrderHandler(c *gin.Context)
	FindOrderByIDHandler(c *gin.Context)
	DeleteOrderHandler(c *gin.Context)
}

// InsertOrderHandler Creates a new order
// @Summary Insert a new order
// @Description Insert a new order with the provided order information
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderRequest body request.OrderRequest true "Order information for registration"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {object} http_error.HttpError
// @Failure 500 {object} http_error.HttpError
// @Router /order [post]
func (oh *orderHandler) InsertOrderHandler(c *gin.Context) {
	var orderRequest request.OrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		logger.Error("Error trying to validate order info", err,
			zap.String("journey", "createOrder"))
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	payee, payeeErr := uuid.Parse(orderRequest.Payee)
	if payeeErr != nil {
		logger.Error("Error trying to parse Payee UUID", payeeErr,
			zap.String("journey", "createOrder"))
		errMessage := http_error.NewBadRequestError("Invalid Payee UUID")
		c.JSON(errMessage.Code, errMessage)
		return
	}

	payer, payerErr := uuid.Parse(orderRequest.Payer)
	if payerErr != nil {
		logger.Error("Error trying to parse Payer UUID", payerErr,
			zap.String("journey", "createOrder"))
		errMessage := http_error.NewBadRequestError("Invalid Payer UUID")
		c.JSON(errMessage.Code, errMessage)
		return
	}

	order := domain.NewOrderDomain(
		orderRequest.Amount,
		payee,
		payer,
	)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := oh.orderService.InsertOrderService(ctxTimeout, order)
	if err != nil {
		logger.Error(
			"Error trying to call InsertOrder service",
			err,
			zap.String("journey", "createOrder"))
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// FindOrderByIDHandler retrieves order information based on the provided order ID.
// @Summary Find Order by ID
// @Description Retrieves order details based on the order ID provided as a parameter.
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "ID of the order to be retrieved"
// @Success 200 {object} response.OrderResponse "User information retrieved successfully"
// @Failure 400 {object} http_error.HttpError "Error: Invalid order ID"
// @Failure 404 {object} http_error.HttpError "User not found"
// @Router /order/{id} [get]
func (oh *orderHandler) FindOrderByIDHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		logger.Error("Error trying to validate orderId",
			parseError,
			zap.String("journey", "findOrderByID"),
		)
		errorMessage := http_error.NewBadRequestError(
			"The ID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	order, err := oh.orderService.FindOrderByIDService(ctxTimeout, id)
	if err != nil {
		logger.Error("Error finding order by ID", err, zap.String("journey", "findOrderByID"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteUserHandler deletes a order with the specified ID.
// @Summary Delete Order
// @Description Deletes a order based on the ID provided as a parameter.
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "ID of the order to be deleted"
// @Success 200
// @Failure 400 {object} http_error.HttpError
// @Failure 500 {object} http_error.HttpError
// @Router /order/{id} [delete]
func (oh orderHandler) DeleteOrderHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		logger.Error("Error trying to validate orderId",
			parseError,
			zap.String("journey", "DeleteOrder"),
		)
		errorMessage := http_error.NewBadRequestError(
			"The ID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	serviceError := oh.orderService.DeleteOrderService(ctxTimeout, id)
	if serviceError != nil {
		logger.Error("Error trying to call deleteOrder service", serviceError, zap.String("journey", "DeleteOrder"))
		c.JSON(serviceError.Code, serviceError)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
