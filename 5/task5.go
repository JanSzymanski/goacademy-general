//Task5
//Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
//Three single digit numbers from one method.
//Three double digit numbers from a second method.
//Three triple digit numbers from a third method.
//Finally sum all methods into a final sum in the main program.

package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func get_x_digits(digits int) int {
    var sum int
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Reading 3 numbers, each of %d digit(s)\n", digits)
    for idx := range 3 {
        var valid_num bool
        for !valid_num {
            fmt.Printf("Please provide number %d: ", idx + 1)
            scanner.Scan()
            if len(scanner.Text()) == digits{
                input, err := strconv.Atoi(scanner.Text())
                if err != nil {
                    fmt.Printf("Cannot convert %s to int, please try again.\n", scanner.Text())
                    continue
                }
                sum += input
                valid_num = true
            } else {
                fmt.Printf("Numbers must be %d digit(s) long. \n", digits)
            }
        }
    }
    return sum
}

func main(){
    var total int
    for i:= 1; i<=3; i++ {
        total += get_x_digits(i)
    }
    fmt.Printf("Total of 3 methods is %d.\n", total)

}
