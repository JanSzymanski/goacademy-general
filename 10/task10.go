//Task10
//Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bearbin/go-age"
	"github.com/goombaio/namegenerator"
)

type Pupil struct {
	id         int
	first_name string
	last_name  string
	dob        time.Time
}

func (p Pupil) GetAge() int {
	return age.Age(p.dob)
}
func (p Pupil) PPrint() {
	fmt.Printf("%d\t%s\t%s\t%v\t%d\n", p.id, p.first_name, p.last_name, p.dob.Format("2006-01-02"), p.GetAge())
}
func PrintTableHeader() {
	fmt.Println("Id\tFName\tLname\tDob\t\tAge")
}

// func AddPupil(store []Pupil) error {

// }
func NormaliseName(str string, cap int) string {
	str = strings.ToUpper(str[:1]) + str[1:] + "     "
	if len(str) > cap {
		return str[:cap-1]
	}
	return str
}

func InitStudents(how_many int) []Pupil {
	pupils := make([]Pupil, 0)
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	for i := range how_many {
		name := nameGenerator.Generate()
		nameSplt := strings.Split(name, "-")
		year := 1990 - i
		pupil := Pupil{
			id:         i,
			first_name: NormaliseName(nameSplt[0], 7),
			last_name:  NormaliseName(nameSplt[1], 7),
			dob:        time.Date(year, time.January, i, i, i, i, i, time.Local),
		}
		pupils = append(pupils, pupil)
	}
	return pupils
}

func main() {
	temp_pupils := InitStudents(10)
	PrintTableHeader()
	for _, pup := range temp_pupils {
		pup.PPrint()
	}
}
