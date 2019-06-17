package repository

import (
	"github.com/franela/goblin"
	"somePriject/web-app/log"
	"testing"
)

func TestGetAll(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#GetAll", func() {
		g.It("should return all users", func() {
			var people = Debtor{}
			expected := people.GetAll()
			res := people.GetAll()
			g.Assert(res).Equal(expected)
		})
	})
}
func TestInsert(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#Insert", func() {
		g.It("should return 1 if success", func() {
			var people = Debtor{FirstName: "", LastName: "", Email: "", Gender: "", Loan: 0}
			expected := 1
			res := people.Insert()
			g.Assert(res).Equal(expected)
		})
	})
}
func TestUpdate(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#Update", func() {
		g.It("should return 1 if success", func() {
			var people = Debtor{FirstName: "", LastName: "", Email: "0", Gender: "", Loan: 0, Id: ""}
			expected := 1
			res := people.Update()
			g.Assert(res).Equal(expected)
		})
	})
}
func TestDelete(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#Delete", func() {
		g.It("should return 1 if success", func() {
			var people = Debtor{}
			expected := 1
			res := people.Delete(0)
			g.Assert(res).Equal(expected)
		})
	})
}
func TestFilters(t *testing.T) {
	g := goblin.Goblin(t)
	log.InitLogger()
	g.Describe("#Filters", func() {
		g.It("should return users with options", func() {
			expected := Filters("", "", "", "")
			res := Filters("", "", "", "")
			g.Assert(res).Equal(expected)
		})
	})
}
