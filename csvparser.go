package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

type Debtor struct {
	Id               int
	FirstName        string
	LastName         string
	Email            string
	Gender           string
	DateRegistration time.Time
	Loan             float64
}

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "gender"}[g]
}

var people []Debtor

func main() {
}

func openParser() {
	csvFile, _ := os.Open("C:/Users/Администратор/go/web-app/awesomeProject/MOCK_DATA.csv")
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for i := 0; i <= 100; i++ {
		line, _ := reader.Read()
		if line[0] != "id" && line[1] != "first_name" && line[2] != "last_name" && line[3] != "email" && line[4] != "gender" && line[5] != "date_registration" && line[6] != "loan" {
			id, _ := strconv.Atoi(line[0])
			dateReg, _ := time.Parse("1/2/2006", line[5])
			loan, _ := strconv.ParseFloat(line[6], 64)
			people = append(people, Debtor{
				Id:               id,
				FirstName:        line[1],
				LastName:         line[2],
				Email:            line[3],
				Gender:           line[4],
				DateRegistration: dateReg,
				Loan:             loan,
			})
		}
	}
}

func checkinDebtors(dateStart, dateEnd time.Time, people []Debtor) {
	for _, value := range people {
		if value.DateRegistration.Unix() >= dateStart.Unix() && value.DateRegistration.Unix() <= dateEnd.Unix() {
			fmt.Println(value)
		}
	}
}

func genderPercent() {
	countMale := 0
	countFemale := 0
	var g Gender
	for i := 0; i <= len(people)-1; i++ {

		if people[i].Gender == "Male" {
			g = Male
		} else if people[i].Gender == "Female" {
			g = Female
		}
		switch g {
		case Male:
			countMale++
		case Female:
			countFemale++
		}
	}
	fmt.Println("Number of men", countMale)
	fmt.Println("Number of women", countFemale)
	percentMale := (countMale * (len(people))) / 100
	percentFemale := (countFemale * (len(people))) / 100
	fmt.Println("Percent of men", percentMale)
	fmt.Println("Percent of women", percentFemale)
}

func getDebtors(firstLoan, secondLoan float64, people []Debtor) {
	for _, value := range people {
		if value.Loan >= firstLoan && value.Loan <= secondLoan {
			fmt.Println(value)
		}
	}
}

func sortByLoan() {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Loan < people[j].Loan
	})
	for _, value := range people {
		fmt.Println(value)
	}
}
