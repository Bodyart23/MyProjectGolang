package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Debtor struct {
	Id               string    `json:"id"`
	FirstName        string    `json:"firstname"`
	LastName         string    `json:"lastname"`
	Email            string    `json:"email"`
	Gender           string    `json:"gender"`
	DateRegistration time.Time `json:"dateregistration"`
	Loan             float64   `json:"loan"`
}

type Crud interface {
	GetAll() []Debtor
	Insert()
	Update()
	Delete(id int)
}

var people []Debtor

func connection() *sql.DB {
	connect, err := sql.Open("postgres", "postgres://postgres:bodyart79317@localhost/?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = connect.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return connect
}

func (p Debtor) GetAll() []Debtor {
	db := connection()
	defer db.Close()
	rows, err := db.Query("select * from debtors")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		d := Debtor{}
		err := rows.Scan(&d.Id, &d.FirstName, &d.LastName, &d.Email, &d.Gender, &d.DateRegistration, &d.Loan)
		if err != nil {
			fmt.Println(err)
			continue
		}
		people = append(people, d)
	}
	return people
}

func (p Debtor) Insert() {
	db := connection()
	defer db.Close()
	result, err := db.Exec("insert into debtors (first_name,last_name,email,gender,date_registration,loan) values ($1,$2,$3,$4,$5,$6)",
		p.FirstName, p.LastName, p.Email, p.Gender, time.Now(), p.Loan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.RowsAffected())
}

func (p Debtor) Update() {
	db := connection()
	defer db.Close()
	result, err := db.Exec("update debtors set first_name = $1, last_name = $2, email = $3, gender = $4, loan =$5 where id = $6",
		p.FirstName, p.LastName, p.Email, p.Gender, p.Loan, p.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}

func (p Debtor) Delete(id int) {
	db := connection()
	defer db.Close()
	result, err := db.Exec("delete from debtors where id = $1", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.RowsAffected())
}

func Filters(name, gender, firstDate, secondDate string) []Debtor {
	var user = Debtor{}
	var users []Debtor
	users = user.GetAll()
	switch {
	case name != "":
		var peopleFromFilter []Debtor
		for _, value := range users {
			if name == value.FirstName {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
			}
		}
		fallthrough
	case gender != "":
		var peopleFromFilter []Debtor
		for _, value := range users {
			if gender == value.Gender {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
			}
		}
		fallthrough
	case firstDate != "":
		var peopleFromFilter []Debtor
		startDate, _ := time.Parse("1/2/2006", firstDate)
		for _, value := range users {
			if startDate.Unix() == 0 {
				continue
			} else if startDate.Unix() <= value.DateRegistration.Unix() {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
				//users = nil
			}
		}
		fallthrough
	case secondDate != "":
		var peopleFromFilter []Debtor
		endDate, _ := time.Parse("1/2/2006", secondDate)
		for _, value := range users {
			if endDate.Unix() == 0 {
				continue
			} else if endDate.Unix() >= value.DateRegistration.Unix() {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
			}
		}
	}

	return users
}
