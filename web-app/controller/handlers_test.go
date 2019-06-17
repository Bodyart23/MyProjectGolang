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
				t.Errorf("handler returned expected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	})
}
