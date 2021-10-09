package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karthik0309/insta_rest_api/controllers"
)
func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UserHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"_id":"6161ca67ad7cdad5b05ac2e4","name":"Karthik","email":"karthik@gmail.com"},{"_id":"6161caacad7cdad5b05ac2e5","name":"Arjun","email":"Arjun@gmail.com"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUserByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "6161ca67ad7cdad5b05ac2e4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UserHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"_id":"6161ca67ad7cdad5b05ac2e4","name":"Karthik","email":"karthik@gmail.com"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestCreateUser(t *testing.T) {

// 	var jsonStr = []byte(`{"name":"xyz","email_address":"xyz@pqr.com","password":"1234567890"}`)

// 	req, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.UserHandler)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":1***,"name":"xyz","email_address":"xyz@pqr.com"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }