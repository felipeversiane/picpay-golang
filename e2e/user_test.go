package e2e

import (
	"net/http"
	"testing"

	"github.com/felipeversiane/picpay-golang.git/internal/entity/request"
	"github.com/google/uuid"
)

func happyData() request.UserRequest {
	return request.UserRequest{
		Email:      "pelvess@example.com",
		Password:   "passwor8!F",
		FirstName:  "Pedro",
		LastName:   "Silva",
		Document:   "0234021102",
		Balance:    1000.00,
		IsMerchant: false,
	}
}

func TestInsertUser_ShouldReturnStatusBadRequest_WhenItHasInvalidData(t *testing.T) {
	t.Log("*** Test Insert User with Invalid Data")

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

func TestDeleteUser_ShouldReturnStatusNotFound_WhenUserIsNotOnDatabase(t *testing.T) {
	t.Log("*** Test Delete User when User is not on Database")

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

func TestFindUserByID_ShouldReturnStatusNotFound_WhenUserIsNotOnDatabase(t *testing.T) {
	t.Log("*** Test Find User by ID when User is not on Database")

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

func TestFindUserByEmail_ShouldReturnStatusNotFound_WhenUserIsNotOnDatabase(t *testing.T) {
	t.Log("*** Test Find User by Email when User is not on Database")

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

func TestFindUserByDocument_ShouldReturnStatusNotFound_WhenUserIsNotOnDatabase(t *testing.T) {
	t.Log("*** Test Find User by Document when User is not on Database")

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

func insertUserSuccessfully(user request.UserRequest, t *testing.T) string {
	t.Log("*** Insert User Sucessfully")

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

	if res["email"].(string) != user.Email {
		t.Fatal("Invalid Email")
	}

	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		t.Fatal("Invalid CreatedAt")
	}

	return id
}

func findUserSuccessfully(id string, t *testing.T) {
	t.Log("*** Find User Sucessfully")
	user := happyData()
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

	if res["email"].(string) != user.Email {
		t.Fatal("Invalid Email")
	}
}

func updateUserSuccessfully(id string, t *testing.T) {
	t.Log("*** Update User Sucessfully")
	api := NewApiClient()

	user := happyData()

	payload := map[string]interface{}{
		"email":       user.Email,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"is_merchant": user.IsMerchant,
		"balance":     user.Balance,
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

	if res["email"].(string) != user.Email {
		t.Fatal("Invalid Email")
	}

	if res["created_at"].(string) == "0001-01-01T00:00:00Z" {
		t.Fatal("Invalid CreatedAt")
	}
}

func deleteUserSuccessfully(id string, t *testing.T) {
	t.Log("*** Delete User Successfully")
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

func TestUserSuccessfully(t *testing.T) {
	t.Log("*** Start User Successfully")

	user := happyData()
	id := insertUserSuccessfully(user, t)
	findUserSuccessfully(id, t)
	updateUserSuccessfully(id, t)
	deleteUserSuccessfully(id, t)

	t.Log("*** End User Successful")
}
