package controller

import (
	"encoding/json"
	"net/http"
	"somePriject/web-app/repository"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var people = repository.Debtor{}
	users := people.GetAll()
	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = repository.Debtor{}
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	user.Loan, _ = strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Insert()
	response := "User was created"

	json.NewEncoder(w).Encode(response)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = repository.Debtor{}
	user.Id = r.FormValue("id")
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	user.Loan, _ = strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Update()
	response := "User was updated"

	json.NewEncoder(w).Encode(response)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user = repository.Debtor{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	user.Delete(id)
	response := "User was deleted"

	json.NewEncoder(w).Encode(response)
}

func UsersFilter(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("firstname")
	gender := r.FormValue("gender")
	firstDate := r.FormValue("firstDate")
	secondDate := r.FormValue("secondDate")
	users := repository.Filters(name, gender, firstDate, secondDate)
	json.NewEncoder(w).Encode(users)
}
