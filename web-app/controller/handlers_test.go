package controller

import (
	"bytes"
	"github.com/franela/goblin"
	"net/http"
	"net/http/httptest"
	"somePriject/web-app/log"
	"testing"
)

func TestGetUsers(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#GetUsers", func() {
		g.It("should return all users", func() {
			req, err := http.NewRequest("GET", "/users", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(GetUsers)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := `[{"id":"43","name":"Bogdan","surname":"Degtar","email":"someEmail","gender":"Male","dateregistration":"2019-06-17T00:00:00Z","loan":123.23},{"id":"44","name":"Bob","surname":"Stinger","email":"emailSome","gender":"Male","dateregistration":"2019-06-17T00:00:00Z","loan":321.32}]`
			if rr.Body.String() == expected {
				t.Errorf("handler returned expected body: got %v want %v",
					rr.Body.String(), expected)
				//g.Assert(rr.Body.String()).Equal(expected)
			}
		})
	})
}

func TestCreateUser(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#CreateUser", func() {
		g.It("should return 'User was created' ", func() {
			var jsonStr = []byte(`{"name":"Dobby","surname":"Potter","email":"dob@pot.com","gender":"Male","loan":"123.321"}`)

			req, err := http.NewRequest("POST", "/user/", bytes.NewBuffer(jsonStr))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreateUser)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			expected := "User was created"
			if rr.Body.String() == expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	})
}
func TestUpdateUser(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#UpdateUser", func() {
		g.It("should return 'User was updated' ", func() {
			var jsonStr = []byte(`{"id":44,"name":"Dobby","surname":"Potter","email":"dob@pot.com","gender":"Male","loan":"123.321"}`)

			req, err := http.NewRequest("PUT", "/user/", bytes.NewBuffer(jsonStr))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UpdateUser)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			expected := "User was updated"
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	})
}
func TestDeleteUser(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#DeleteUser", func() {
		g.It("should return 'User was deleted' ", func() {
			req, err := http.NewRequest("DELETE", "/user/", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("id", "43")
			req.URL.RawQuery = q.Encode()
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(DeleteUser)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			expected := "User was deleted"
			if rr.Body.String() == expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	})
}
func TestUsersFilter(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#UsersFilter", func() {
		g.It("should return filtering users ", func() {
			req, err := http.NewRequest("DELETE", "/user/", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("name", "Bogdan")
			req.URL.RawQuery = q.Encode()
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UsersFilter)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			expected := `{"id":"43","name":"Bogdan","surname":"Degtar","email":"someEmail","gender":"Male","dateregistration":"2019-06-17T00:00:00Z","loan":123.23}`
			if rr.Body.String() == expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	})
}
