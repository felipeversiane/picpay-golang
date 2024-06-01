package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/response"
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type orderService struct {
	orderRepository repository.OrderRepository
	userService     UserService
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	userService UserService,
) OrderService {
	return &orderService{
		orderRepository, userService,
	}
}

type OrderService interface {
	InsertOrderService(ctx context.Context, order domain.OrderDomainInterface) (response.OrderResponse, *http_error.HttpError)
	ValidateAuthorization() bool
	FindOrderByIDService(ctx context.Context, id uuid.UUID) (response.OrderResponse, *http_error.HttpError)
}

func (oc *orderService) InsertOrderService(ctx context.Context, order domain.OrderDomainInterface) (response.OrderResponse, *http_error.HttpError) {
	payer, err := oc.userService.FindUserByIDService(order.GetPayer(), ctx)
	if err != nil {
		return response.OrderResponse{}, http_error.NewBadRequestError("Payer not found")
	}
	payee, err := oc.userService.FindUserByIDService(order.GetPayee(), ctx)
	if err != nil {
		return response.OrderResponse{}, http_error.NewBadRequestError("Payee not found")
	}

	if payer.Balance < order.GetAmount() {
		return response.OrderResponse{}, http_error.NewBadRequestError("Insufficient balance")
	}

	if payer.IsMerchant {
		return response.OrderResponse{}, http_error.NewBadRequestError("Merchants cannot send money")
	}

	_, err = oc.userService.FindUserByIDService(order.GetPayee(), ctx)
	if err != nil {
		return response.OrderResponse{}, http_error.NewBadRequestError("Payee not found")
	}

	if !oc.ValidateAuthorization() {
		return response.OrderResponse{}, http_error.NewBadRequestError("Order not authorized")
	}

	result, err := oc.orderRepository.InsertOrderRepository(ctx, order)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "InsertOrder"))
		return response.OrderResponse{}, err
	}

	payer.Balance -= order.GetAmount()
	payerUpdate := domain.NewUserUpdateDomain(payer.FirstName, payer.LastName, payer.Balance, payer.IsMerchant)
	if _, err := oc.userService.UpdateUserService(payer.ID, payerUpdate, ctx); err != nil {
		logger.Error("Error updating payer balance", err, zap.String("journey", "UpdatePayerBalance"))
		return response.OrderResponse{}, http_error.NewInternalServerError("Error updating payer balance")
	}

	payee.Balance += order.GetAmount()
	payeeUpdate := domain.NewUserUpdateDomain(payee.FirstName, payee.LastName, payee.Balance, payee.IsMerchant)
	if _, err := oc.userService.UpdateUserService(payee.ID, payeeUpdate, ctx); err != nil {
		logger.Error("Error updating payee balance", err, zap.String("journey", "UpdatePayeeBalance"))

		return response.OrderResponse{}, http_error.NewInternalServerError("Error updating payee balance")
	}

	return result, nil
}

func (oc *orderService) ValidateAuthorization() bool {
	url := os.Getenv("AUTHORIZATION_URL")

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		logger.Error("Error calling authorization service", err, zap.String("journey", "ValidateAuthorization"))
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", err, zap.String("journey", "ValidateAuthorization"))
		return false
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Error("Error unmarshalling JSON", err, zap.String("journey", "ValidateAuthorization"))
		return false
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return false
	}

	authorization, ok := data["authorization"].(bool)
	return ok && authorization
}

func (oc *orderService) FindOrderByIDService(ctx context.Context, id uuid.UUID) (response.OrderResponse, *http_error.HttpError) {
	result, err := oc.orderRepository.FindOrderByIDRepository(ctx, id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "FindOrderByID"))
		return response.OrderResponse{}, err
	}
	return result, nil
}
