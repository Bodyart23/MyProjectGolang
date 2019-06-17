package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"somePriject/web-app/log"
	"time"
)

type Debtor struct {
	Id               string    `json:"id"`
	FirstName        string    `json:"name"`
	LastName         string    `json:"surname"`
	Email            string    `json:"email"`
	Gender           string    `json:"gender"`
	DateRegistration time.Time `json:"dateregistration"`
	Loan             float64   `json:"loan"`
}

type Crud interface {
	GetAll() []Debtor
	Insert() int
	Update() int
	Delete(id int) int
}

//var people []Debtor

func connection() *sql.DB {
	log.Info.Println("call connection function")
	connect, err := sql.Open("postgres", "postgres://postgres:bodyart79317@localhost/?sslmode=disable")
	if err != nil {
		log.Error.Println(err)
	}
	log.Debug.Println("return from connection function", connect)
	return connect
}

func (p Debtor) GetAll() []Debtor {
	var people []Debtor
	log.Info.Println("call GetAll method")
	db := connection()
	defer db.Close()
	rows, err := db.Query("select * from debtors")
	if err != nil {
		log.Error.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		d := Debtor{}
		err := rows.Scan(&d.Id, &d.FirstName, &d.LastName, &d.Email, &d.Gender, &d.DateRegistration, &d.Loan)
		if err != nil {
			log.Error.Println(err)
			continue
		}
		people = append(people, d)
	}
	log.Debug.Println("return from GetAll method", people)
	return people
}

func (p Debtor) Insert() int {
	log.Info.Println("call Insert method")
	db := connection()
	defer db.Close()
	result, err := db.Exec("insert into debtors (first_name,last_name,email,gender,date_registration,loan) values ($1,$2,$3,$4,$5,$6)",
		p.FirstName, p.LastName, p.Email, p.Gender, time.Now(), p.Loan)
	if err != nil {
		log.Error.Println(err)
	}
	count, _ := result.RowsAffected()
	log.Debug.Println("insert successful")
	return int(count)
}

func (p Debtor) Update() int {
	log.Info.Println("call Update method")
	db := connection()
	defer db.Close()
	result, err := db.Exec("update debtors set first_name = $1, last_name = $2, email = $3, gender = $4, loan =$5 where id = $6",
		p.FirstName, p.LastName, p.Email, p.Gender, p.Loan, p.Id)
	if err != nil {
		log.Error.Println(err)
	}
	count, _ := result.RowsAffected()
	log.Debug.Println("update successful")
	return int(count)
}

func (p Debtor) Delete(id int) int {
	log.Info.Println("call Delete method")
	db := connection()
	defer db.Close()
	result, err := db.Exec("delete from debtors where id = $1", id)
	if err != nil {
		log.Error.Println(err)
	}
	count, _ := result.RowsAffected()
	log.Debug.Println("delete successful")
	return int(count)
}

func Filters(name, gender, firstDate, secondDate string) []Debtor {
	log.Info.Println("call Filters function")
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
		startDate, err := time.Parse("1/2/2006", firstDate)
		if err != nil {
			log.Error.Println(err)
		}
		for _, value := range users {
			if startDate.Unix() == 0 {
				continue
			} else if startDate.Unix() <= value.DateRegistration.Unix() {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
			}
		}
		fallthrough
	case secondDate != "":
		var peopleFromFilter []Debtor
		endDate, err := time.Parse("1/2/2006", secondDate)
		if err != nil {
			log.Error.Println(err)
		}
		for _, value := range users {
			if endDate.Unix() == 0 {
				continue
			} else if endDate.Unix() >= value.DateRegistration.Unix() {
				peopleFromFilter = append(peopleFromFilter, value)
				users = peopleFromFilter
			}
		}
	}
	log.Debug.Println("return from Filters function:", users)
	return users
}
