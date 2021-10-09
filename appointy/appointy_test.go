package appointy_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/appointy"
)

// To test the CreateUser HTTP Request Handler
func Test_CreateUser_1(t *testing.T) {
	var jsonStr = []byte(`{"Name":"test 0","Email":"test01@gmail.com","Password":"testingPassword"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appointy.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: ot %v want %v",
			status, http.StatusOK)
	}
	response := rr.Body.String()
	fmt.Printf(response)
}

// To test the CreatePost HTTP Request Handler
func Test_CreatePost_1(t *testing.T) {
	var jsonStr = []byte(`{"Caption":"Image1","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","UserID":{"$oid":"616170b26888c9a871be8efa"}}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appointy.CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: ot %v want %v",
			status, http.StatusOK)
	}
	response := rr.Body.String()
	fmt.Printf(response)
}

// To test the GetUserById HTTP Request Handler
func Test_GetUserById_1(t *testing.T) {

	req, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "6161a15b0f8772368d20603f")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appointy.GetUserById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"_id":{"$oid":"6161a15b0f8772368d20603f"},"Name":"test3","Email":"test3@gmail.com","Password":"\r��\u0012d,��?t�v�t��s�)(#�1;�׊�|ݏr#Z��S�&y~x�N��\u0000/��\u0007K\u0006m��\u0011N2�"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetPostById_1(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "616123f89760661304a9a605")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appointy.GetPostById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"_id":{"$oid":"616123f89760661304a9a605"},"Caption":"testImage1","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","UserID":{"$oid":"61611e9ac421e095e2660e41"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// To test the ListPostsByUser HTTP Request Handler
func Test_ListPostsByUser_1(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "616170b26888c9a871be8efa")
	q.Add("page", "2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appointy.ListPostsByUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Caption":"Image5","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","Timestamp":{"T":1633783435,"I":0},"UserID":"616170b26888c9a871be8efa","_id":"61618e8b29cd22ab9d2a7b6d"}
	{"Caption":"Image6","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","Timestamp":{"T":1633783439,"I":0},"UserID":"616170b26888c9a871be8efa","_id":"61618e8f29cd22ab9d2a7b6e"}
	{"Caption":"Image7","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","Timestamp":{"T":1633783442,"I":0},"UserID":"616170b26888c9a871be8efa","_id":"61618e9229cd22ab9d2a7b6f"}
	{"Caption":"Image8","ImageURL":"https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_204af4585e7c679cdd1bc0842ff6c82b/appointy.png","Timestamp":{"T":1633783446,"I":0},"UserID":"616170b26888c9a871be8efa","_id":"61618e9629cd22ab9d2a7b70"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
