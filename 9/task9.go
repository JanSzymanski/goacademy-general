//Task9
//Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("===== Made with Scanner =====")
	fmt.Println("==== With aliased Prints ====")

	var pf = fmt.Printf
	var pl = fmt.Println

	instructions := [3]string{"first", "middle", "last"}
	inputs := [3]string{}
	scanner := bufio.NewScanner(os.Stdin)

	for idx, val := range instructions {
		pf("Please enter %q name: ", val)
		scanner.Scan()
		input := scanner.Text()
		inputs[idx] = input
	}

    pl("Entered name is: ")
    //below the conversion array -> slice is not necessary but just for the purpose of the exercise.
    for _, name := range inputs[:] {
        pl(name)

    }

}
