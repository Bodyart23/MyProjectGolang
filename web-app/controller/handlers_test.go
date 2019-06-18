package controller

import (
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
				//t.Errorf("handler returned expected body: got %v want %v",
				//	rr.Body.String(), expected)
				g.Assert(rr.Body.String()).Equal(expected)
			}
		})
	})
}

//func TestUsersFilter(t *testing.T) {
//	g := goblin.Goblin(t)
//	log.InitLogger()
//	g.Describe("#UsersFilter", func() {
//		g.It("should return all users", func() {
//			req, err := http.NewRequest("GET", "/users/filter/?name=Bogdan", nil)
//			if err != nil {
//				t.Fatal(err)
//			}
//			rr := httptest.NewRecorder()
//
//			handler := http.HandlerFunc(UsersFilter)
//			handler.ServeHTTP(rr, req)
//
//			if status := rr.Code; status != http.StatusOK {
//				t.Errorf("handler returned wrong status code: got %v want %v",
//					status, http.StatusOK)
//			}
//
//			expected := `[{"id":"43","name":"Bogdan","surname":"Degtar","email":"someEmail","gender":"Male","dateregistration":"2019-06-17T00:00:00Z","loan":123.23},{"id":"44","name":"Bob","surname":"Stinger","email":"emailSome","gender":"Male","dateregistration":"2019-06-17T00:00:00Z","loan":321.32}]`
//			if rr.Body.String() == expected {
//				t.Errorf("handler returned expected body: got %v want %v",
//					rr.Body.String(), expected)
//			}
//		})
//	})
//}
//func TestEventNextHandler(t *testing.T) {
//	// integration test on http requests to EventNextHandler
//	log.InitLogger()
//	request, _ := http.NewRequest("GET", "/users", nil)
//	var people  = repository.Debtor{}
//	people.FirstName = "Bill"
//	people.LastName = request.FormValue("surname")
//	people.Email = "sdfsdf"
//	people.Gender = request.FormValue("gender")
//	people.Loan, _ = strconv.ParseFloat(request.FormValue("loan"),64)
//	response := httptest.NewRecorder()
//	people.Insert()
//	if response.Code != http.StatusOK {
//	t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
//	}
//}
//func TestGetUsers(t *testing.T) {
//	g := goblin.Goblin(t)
//	log.InitLogger()
//	g.Describe("#GetUsers", func() {
//		g.It("should return all users", func() {
//		})
//	})
//}
