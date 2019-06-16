package controller

import (
	"encoding/json"
	"net/http"
	"somePriject/web-app/log"
	"somePriject/web-app/repository"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Info.Println("call GetUser function")
	var people = repository.Debtor{}
	users := people.GetAll()
	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Info.Println("call CreateUser function")
	var user = repository.Debtor{}
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	loan, err := strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Loan = loan
	if err != nil {
		log.Error.Println(err)
	}
	log.Debug.Println("fields for Insert:", user)
	user.Insert()
	response := "User was created"
	json.NewEncoder(w).Encode(response)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Info.Println("call UpdateUser function")
	var user = repository.Debtor{}
	user.Id = r.FormValue("id")
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	loan, err := strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Loan = loan
	if err != nil {
		log.Error.Println(err)
	}
	log.Debug.Println("fields for Update:", user)
	user.Update()
	response := "User was updated"

	json.NewEncoder(w).Encode(response)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Info.Println("call DeleteUser function")
	var user = repository.Debtor{}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Error.Println(err)
	}
	log.Debug.Println("fields for Delete:", user)
	user.Delete(id)
	response := "User was deleted"

	json.NewEncoder(w).Encode(response)
}

func UsersFilter(w http.ResponseWriter, r *http.Request) {
	log.Info.Println("call UserFilter function")
	name := r.FormValue("firstname")
	gender := r.FormValue("gender")
	firstDate := r.FormValue("firstDate")
	secondDate := r.FormValue("secondDate")
	log.Debug.Println("fields for Filter:", name, gender, firstDate, secondDate)
	users := repository.Filters(name, gender, firstDate, secondDate)
	json.NewEncoder(w).Encode(users)
}
