package e2e

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func happyData() (email, password, firstName, lastName string, isMerchant bool, document string, balance float64) {
	email = "pepplvess@example.com"
	password = "passwor8!F"
	firstName = "Pedro"
	lastName = "Silva"
	isMerchant = false
	document = "02340229102"
	balance = 1000.00
	return
}

func TestInsertUser_ShouldReturnStatusBadRequest_WhenItHasInvalidData(t *testing.T) {
	api := NewApiClient()
	params := []map[string]interface{}{
		nil,
		{},
		{"other": "value"},
		{"email": "invalid_email", "password": "passwordD!@3", "first_name": "John", "last_name": "Doe", "document": "12345678910", "balance": 200.00, "is_merchant": false},
		{"email": "jhondoe@jhondoe.com", "password": "passwordD12", "first_name": "John", "last_name": "Doe", "document": "12345678910", "balance": 200.00, "is_merchant": false},
		{"email": "jhondoe@jhondoe.com", "password": "passwordD12!", "first_name": "", "last_name": "Doe", "document": "12345678910", "balance": 200.00, "is_merchant": false},
		{"email": "jhondoe@jhondoe.com", "password": "passwordD12!", "first_name": "Jhon", "last_name": "", "document": "12345678910", "balance": 200.00, "is_merchant": false},
		{"email": "jhondoe@jhondoe.com", "password": "passwordD12!", "first_name": "Jhon", "last_name": "Doe", "document": "12344567810022", "balance": 200.00, "is_merchant": false},
		{"email": "jhondoe@jhondoe.com", "password": "passwordD12!", "first_name": "Jhon", "last_name": "Doe", "document": "12345678910", "balance": -200.00, "is_merchant": false},
	}

	for _, p := range params {
		resp, err := api.Post("/user", p)
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

func TestDeleteUser_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
	api := NewApiClient()
	id := uuid.NewString()

	resp, err := api.Delete("/user/" + id)
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

func TestGetUser_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
	api := NewApiClient()
	id := uuid.NewString()

	resp, err := api.Get("/user/" + id)
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

func TestGetUserByEmail_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
	api := NewApiClient()
	email := "jhondoe@xxx.com"

	resp, err := api.Get("/user/find_user_by_email/" + email)
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

func TestGetUserByDocument_ShouldReturnStatusNotFound_WhenPostIdIsNotOnDatabase(t *testing.T) {
	api := NewApiClient()
	document := "041906777777"

	resp, err := api.Get("/user/find_user_by_document/" + document)
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

func TestUserSucessfully(t *testing.T) {
	t.Log("*** Start User Successful")

	id := insertSuccessfully(t)
	findSuccessfully(id, t)
	updateSuccessfully(id, t)
	deleteSuccessfully(id, t)

	t.Log("*** End User Successful")
}

func insertSuccessfully(t *testing.T) string {
	t.Log("*** Insert User")

	api := NewApiClient()
	email, password, firstName, lastName, isMerchant, document, balance := happyData()

	payload := map[string]interface{}{
		"email":       email,
		"password":    password,
		"first_name":  firstName,
		"last_name":   lastName,
		"is_merchant": isMerchant,
		"balance":     balance,
		"document":    document,
	}

	resp, err := api.Post("/user", payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusCreated,
			resp.Status,
		)
	}

	res, err := api.ParseBody(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	id := res["id"].(string)

	if id == "" {
		t.Fatal("Invalid ID")
	}

	if res["email"].(string) != email {
		t.Fatal("Invalid Email")
	}

	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		t.Fatal("Invalid CreatedAt")
	}

	return id
}

func findSuccessfully(id string, t *testing.T) {
	t.Log("*** Find User")
	email, _, _, _, _, _, _ := happyData()
	api := NewApiClient()

	resp, err := api.Get("/user/" + id)
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

	if res["email"].(string) != email {
		t.Fatal("Invalid Email")
	}

}

func updateSuccessfully(id string, t *testing.T) {
	t.Log("*** Update User")
	api := NewApiClient()

	email, _, firstName, lastName, isMerchant, _, balance := happyData()

	payload := map[string]interface{}{
		"email":       email,
		"first_name":  firstName,
		"last_name":   lastName,
		"is_merchant": isMerchant,
		"balance":     balance,
	}

	resp, err := api.Put("/user/"+id, payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusCreated,
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

	if res["email"].(string) != email {
		t.Fatal("Invalid Email")
	}

	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		t.Fatal("Invalid CreatedAt")
	}
}

func deleteSuccessfully(id string, t *testing.T) {
	t.Log("*** Delete User")
	api := NewApiClient()

	resp, err := api.Delete("/user/" + id)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf(
			"Invalid Status Code. Expected Status \"%d\" and received \"%s\"",
			http.StatusNoContent,
			resp.Status,
		)
	}
}
