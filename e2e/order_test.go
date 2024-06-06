package e2e

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func firstUser() (email, password, firstName, lastName string, isMerchant bool, document string, balance float64) {
	email = "ollvess@example.com"
	password = "passwor8!F"
	firstName = "Oliveira"
	lastName = "Silva"
	isMerchant = false
	document = "02340239102"
	balance = 1000.00
	return
}

func secondUser() (email, password, firstName, lastName string, isMerchant bool, document string, balance float64) {
	email = "pplvess@example.com"
	password = "passwor8!F"
	firstName = "Pedro"
	lastName = "Silva"
	isMerchant = false
	document = "02341129102"
	balance = 200.00
	return
}

func TestInsertOrder_ShouldReturnStatusBadRequest_WhenItHasInvalidData(t *testing.T) {
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

func TestGetOrder_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
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

func TestDeleteOrder_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
	api := NewApiClient()
	id := uuid.NewString()

	resp, err := api.Delete("/order/" + id)
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
