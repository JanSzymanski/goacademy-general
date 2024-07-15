package main


//Task4
//Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// *Display the array content in ascending sequential order 1 to 10.
// *Display the array content in descending sequential order 10 to 1.
// *Count even numbers and odd numbers in increasing and decreasing sequential order.
// *Display the even and odd count sequences to screen.
import (
    "fmt"
)

func main() {
    array := [10]int {1,2,3,4,5,6,7,8,9,10}

    fmt.Println("Array in asc sequential order")
    for i:=0; i<len(array); i++{
        fmt.Printf("%d ", array[i])
    }
    fmt.Println("\nArray in desc sequential order")
    for i:=len(array)-1; i>=0; i--{
        fmt.Printf("%d ", array[i])
    }
    fmt.Println("\nEven numbers")
    var sum_of_even int
    var sum_of_odd int
    for i:=0; i<len(array); i++{
        num := array[i]
        if num%2==0{
            fmt.Printf("%d ", num)
            sum_of_even+=num
        }
    }
    fmt.Println()
    for i:=len(array)-1; i>=0; i--{
        num := array[i]
        if num%2==0{
            fmt.Printf("%d ", num)
        }
   }
    fmt.Printf(" Sum: %d\n", sum_of_even)
    fmt.Println("Odd numbers")
    for i:=0; i<len(array); i++{
        num := array[i]
        if num%2==1{
            fmt.Printf("%d ", num)
            sum_of_odd+=num
        }
    }
    fmt.Println()
    for i:=len(array)-1; i>=0; i--{
        num := array[i]
        if num%2==1{
            fmt.Printf("%d ", num)
        }
   }
    fmt.Printf(" Sum: %d\n", sum_of_odd)
}
