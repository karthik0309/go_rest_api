package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karthik0309/insta_rest_api/controllers"
)

func TestGetPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"_id":"6161cb92ad7cdad5b05ac2e6","UID":"6161ca67ad7cdad5b05ac2e4","caption":"test1","image_url":"http://test.img","created_at":"time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"},{"_id":"6161cb9ead7cdad5b05ac2e7","UID":"6161ca67ad7cdad5b05ac2e4","caption":"test2","image_url":"http://test2.img","created_at":"time.Date(2021, time.October, 9, 22, 34, 30, 362991000, time.Local)"},{"_id":"6161cbc5ad7cdad5b05ac2e8","UID":"6161caacad7cdad5b05ac2e5","caption":"test3","image_url":"http://test3.img","created_at":"time.Date(2021, time.October, 9, 22, 35, 9, 536000000, time.Local)"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetPostByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "6161cb92ad7cdad5b05ac2e6")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"_id":"6161cb92ad7cdad5b05ac2e6","UID":"6161ca67ad7cdad5b05ac2e4","caption":"test1","image_url":"http://test.img","created_at":"time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetPostByUserID(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "6161ca67ad7cdad5b05ac2e4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPostsByUserId)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"_id":"6161cb92ad7cdad5b05ac2e6","UID":"6161ca67ad7cdad5b05ac2e4","caption":"test1","image_url":"http://test.img","created_at":"time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"},{"_id":"6161cb9ead7cdad5b05ac2e7","UID":"6161ca67ad7cdad5b05ac2e4","caption":"test2","image_url":"http://test2.img","created_at":"time.Date(2021, time.October, 9, 22, 34, 30, 362991000, time.Local)"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestCreatePost(t *testing.T) {

// 	var jsonStr = []byte(`{"caption":"xyz","image_url":"http://img.com","UID":"1234567890"}`)

// 	req, err := http.NewRequest("POST", "/posts/", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.PostHandler)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":1***,"caption":"xyz","image_url":"http://img.com","UID":"1234567890"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }