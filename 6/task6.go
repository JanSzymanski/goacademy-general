//Task6
//Create a program that calculates the age of a person given their date of birth. [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
//(Use the github.com/bearbin/go-age to aid in the creation of this app. Also review unit testing applied to the application age.go within the imported package.)




package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strconv"
    "log"
    "github.com/bearbin/go-age"
)


func main() {
    var pf = fmt.Printf
    var pl = fmt.Println
    scanner := bufio.NewScanner(os.Stdin)
    var year, month, day int = -1, -1, -1

    pl("Please provide your date of birth")
    for year == -1 {
        fmt.Print("Year: ")
        scanner.Scan()
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            pl("Error converting inputed year to number")
            continue
        }
        year = num
    }
    for month == -1 {
        fmt.Print("Month: ")
        scanner.Scan()
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            pl("Error converting inputed month to number")
            continue
        }
        if num < 1 || num > 12 {
            pl("Month has to be between 1 and 12")
            continue
        }
        month = num
    }
    for day == -1 {
        fmt.Print("Day: ")
        scanner.Scan()
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            pl("Error converting inputed day to number")
            continue
        }
        //following day range protection is not ideal but OK-ish.
        //Would be best to check how many days for each month + leap year Feb
        if num < 1 || num > 31 {
            pl("Day has to be between 1 and 31")
            continue
        }
        day = num
    }
    dateString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
    bod, err := time.Parse("2006-01-02", dateString)
    if err != nil {
        pl(err)
        log.Fatal("There was an error while parsing date")
    }
    pf("You are %v years old.\n", age.Age(bod))



}
