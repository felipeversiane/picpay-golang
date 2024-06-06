package e2e

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/google/uuid"
)

func firstUser() request.UserRequest {
	return request.UserRequest{
		Email:      "olikvess@example.com",
		Password:   "passwor8!F",
		FirstName:  "Oliveira",
		LastName:   "Silva",
		Document:   "2222222",
		Balance:    1000.00,
		IsMerchant: true,
	}
}

func secondUser() request.UserRequest {
	return request.UserRequest{
		Email:      "pedrinls@example.com",
		Password:   "passwor8!F",
		FirstName:  "Pedro",
		LastName:   "Silva",
		Document:   "3333333",
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
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf(
				"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
				http.StatusBadRequest,
				resp.Status,
			)
		}
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
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusNotFound,
			resp.Status,
		)
	}
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
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusBadRequest,
			resp.Status,
		)
	}
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
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusBadRequest,
			resp.Status,
		)
	}
}

func InsertOrderSuccessfully(payer string, payee string, t *testing.T) string {
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
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err.Error())
			}

			bodyString := string(bodyBytes)
			if bodyString == "Order not authorized" {
				t.Log("Order not authorized, retrying...")
				continue
			}
		}

		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
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

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusOK,
			resp.Status,
		)
	}

	res, err := api.ParseBody(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	if res["id"].(string) != id {
		t.Fatal("Invalid ID")
	}
}

func TestFlow(t *testing.T) {
	t.Log("*** Start Flow")

	firstUser := firstUser()
	secondUser := secondUser()

	firstID := insertUserSuccessfully(firstUser, t)
	secondID := insertUserSuccessfully(secondUser, t)
	InsertOrder_ShouldReturnStatusBadRequest_InsufficientBalance(secondID, firstID, t)
	InsertOrder_ShouldReturnStatusBadRequest_MerchantCannotSendMoney(firstID, secondID, t)
	orderID := InsertOrderSuccessfully(secondID, firstID, t)
	findOrderSuccessfully(orderID, t)
	deleteUserSuccessfully(firstID, t)
	deleteUserSuccessfully(secondID, t)

	t.Log("*** End Flow Successful")
}
