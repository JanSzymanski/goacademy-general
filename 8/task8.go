//Task8
// Create a program that: [Write File][Read File][I/O Package][I/O]
// Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// Reads a list of cities from the newly created file.
// Displays the list of cities in alphabetical order.



package main

import (
    "fmt"
    "os"
    "strings"
    "slices"
)

func main() {
    initial_cities := [8]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

    f, err := os.Create("cities")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()
    fmt.Println("Initiating a file")
    for _, v := range initial_cities {
        l, err := f.WriteString(v + "\n")
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Printf("Successfully wrote %d bytes.\n", l)
    }
    fmt.Println("Initiation complete")
    fmt.Println("======================")
    fmt.Println("Now, reading that file back")
    fmt.Println("And sorting")

    read_f, err := os.ReadFile("cities")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    read_fs := strings.Trim(string(read_f), "\n")
    sliced_cities := strings.Split(read_fs, "\n")
    slices.Sort(sliced_cities)
    for _, city := range sliced_cities {
        fmt.Println(city)
    }
}
