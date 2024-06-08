package e2e

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/google/uuid"
)

func firstUser() request.UserRequest {
	return request.UserRequest{
		Email:      "olisvsss@example.com",
		Password:   "passwor8!F",
		FirstName:  "Oliveira",
		LastName:   "Silva",
		Document:   "992375832",
		Balance:    1000.00,
		IsMerchant: true,
	}
}

func secondUser() request.UserRequest {
	return request.UserRequest{
		Email:      "pedrnllxss@example.com",
		Password:   "passwor8!F",
		FirstName:  "Pedro",
		LastName:   "Silva",
		Document:   "323467222",
		Balance:    200.00,
		IsMerchant: false,
	}
}
func TestInsertOrder_ShouldReturnStatusBadRequest_WhenItHasInvalidData(t *testing.T) {
	t.Log("*** Test Insert Order with Invalid Data")

	api := NewApiClient()
	params := []map[string]interface{}{
		nil,
		{},
		{"other": "value"},
		{"payee": uuid.NewString(), "payer": "not", "amount": 100.00},
		{"payee": "not", "payer": uuid.NewString(), "amount": 100.00},
		{"payee": uuid.NewString(), "payer": uuid.NewString(), "amount": -100.00},
	}

	for _, p := range params {
		resp, err := api.Post("/order", p)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer resp.Body.Close()
		assertStatusCode(t, resp, http.StatusBadRequest)
	}
}

func TestFindOrder_ShouldReturnStatusNotFound_WhenOrderIdIsNotOnDatabase(t *testing.T) {
	t.Log("*** Test Find Order when Order is not on Database")

	api := NewApiClient()
	id := uuid.NewString()

	resp, err := api.Get("/order/" + id)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	assertStatusCode(t, resp, http.StatusNotFound)
}

func InsertOrder_ShouldReturnStatusBadRequest_InsufficientBalance(payer string, payee string, t *testing.T) {
	t.Log("*** Test Insert Order with Insufficient Balance")

	api := NewApiClient()

	payload := map[string]interface{}{
		"amount": 201.00,
		"payee":  payee,
		"payer":  payer,
	}
	resp, err := api.Post("/order", payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	assertStatusCode(t, resp, http.StatusBadRequest)
}

func InsertOrder_ShouldReturnStatusBadRequest_MerchantCannotSendMoney(payer string, payee string, t *testing.T) {
	t.Log("*** Test Insert Order with Payer as Merchant")

	api := NewApiClient()

	payload := map[string]interface{}{
		"amount": 200.00,
		"payer":  payer,
		"payee":  payee,
	}
	resp, err := api.Post("/order", payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	assertStatusCode(t, resp, http.StatusBadRequest)
}

func insertOrderSuccessfully(payer string, payee string, t *testing.T) string {
	t.Log("*** Insert Order Successfully")

	api := NewApiClient()

	payload := map[string]interface{}{
		"amount": 100.00,
		"payer":  payer,
		"payee":  payee,
	}

	for {
		resp, err := api.Post("/order", payload)
		if err != nil {
			t.Fatal(err.Error())
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			res, err := api.ParseBody(resp)
			if err != nil {
				t.Fatal(err.Error())
			}

			id := res["id"].(string)

			if id == "" {
				t.Fatal("Invalid ID")
			}

			if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
				t.Fatal("Invalid CreatedAt")
			}

			if res["payer"].(string) != payer {
				t.Fatal("Invalid Payer")
			}

			if res["payee"].(string) != payee {
				t.Fatal("Invalid Payee")
			}

			return id
		}

		if resp.StatusCode == http.StatusBadRequest {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err.Error())
			}

			var bodyJson map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &bodyJson); err != nil {
				t.Fatal(err.Error())
			}

			if bodyJson["message"] == "Order not authorized" {
				t.Log("Order not authorized, retrying...")
				return insertOrderSuccessfully(payer, payee, t)
			} else {
				t.Fatalf("Bad Request: %s", bodyJson["message"])
			}
		}

		t.Fatalf(
			"Invalid Status Code. Expected Stat9us \"%d\" and received \"%s\"",
			http.StatusCreated,
			resp.Status,
		)
	}
}

func findOrderSuccessfully(id string, t *testing.T) {
	t.Log("*** Find Order Successfully")
	api := NewApiClient()

	resp, err := api.Get("/order/" + id)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	assertStatusCode(t, resp, http.StatusOK)

	res, err := api.ParseBody(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	if res["id"].(string) != id {
		t.Fatal("Invalid ID")
	}
}

func deleteOrderUserSuccessfully(id string, t *testing.T) {
	t.Log("*** Delete Order User Successfully")
	api := NewApiClient()

	resp, err := api.Delete("/user/" + id)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assertStatusCode(t, resp, http.StatusNoContent)
}

func insertOrderUserSuccessfully(user request.UserRequest, t *testing.T) string {
	t.Log("*** Insert Order User Successfully")

	api := NewApiClient()

	payload := map[string]interface{}{
		"email":       user.Email,
		"password":    user.Password,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"document":    user.Document,
		"balance":     user.Balance,
		"is_merchant": user.IsMerchant,
	}

	resp, err := api.Post("/user", payload)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assertStatusCode(t, resp, http.StatusCreated)

	res, err := api.ParseBody(resp)
	if err != nil {
		t.Fatal(err)
	}

	id := res["id"].(string)
	if id == "" {
		t.Fatal("Invalid ID")
	}
	if res["email"].(string) != user.Email {
		t.Fatal("Invalid Email")
	}
	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		t.Fatal("Invalid CreatedAt")
	}

	return id
}

func TestOrderFlow(t *testing.T) {
	t.Log("*** Start Order Flow")

	firstUser := firstUser()
	secondUser := secondUser()

	firstID := insertOrderUserSuccessfully(firstUser, t)
	secondID := insertOrderUserSuccessfully(secondUser, t)
	InsertOrder_ShouldReturnStatusBadRequest_InsufficientBalance(secondID, firstID, t)
	InsertOrder_ShouldReturnStatusBadRequest_MerchantCannotSendMoney(firstID, secondID, t)
	orderID := insertOrderSuccessfully(secondID, firstID, t)
	findOrderSuccessfully(orderID, t)
	deleteOrderUserSuccessfully(firstID, t)
	deleteOrderUserSuccessfully(secondID, t)

	t.Log("*** End Order Flow Successful")
}
