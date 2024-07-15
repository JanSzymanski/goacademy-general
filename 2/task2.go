package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Task: Create a program that lets the user input a first name,
//Task: middle name and last name. Display the person's full name on one line. [Keyboard input]
//Observation: the for (while) loop didn't work because ReadString returns new line
//We can potentially also use fmt.Scanf - to be tested!

func main() {
	reader := bufio.NewReader(os.Stdin)
	var f_name, m_name, l_name string
	fmt.Println("===== Made with Reader =====")
	for f_name == "" {
		fmt.Print("Please enter your first name: ")
		string_read, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from keyboard")
			continue
		}
		f_name = string_read
	}
	for m_name == "" {
		fmt.Print("Please enter your middle name: ")
		string_read, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from keyboard")
			continue
		}
		m_name = string_read
	}
	for l_name == "" {
		fmt.Print("Please enter your last name: ")
		string_read, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from keyboard")
			continue
		}
		l_name = string_read
	}
	combined_name := fmt.Sprintf("Entered name was: %s %s %s\n",
		strings.Trim(f_name, "\n"),
		strings.Trim(m_name, "\n"),
		strings.Trim(l_name, "\n"))
	combined_name = strings.ReplaceAll(combined_name, "  ", " ")

	fmt.Println(combined_name)

	fmt.Println("===== Made with Scanner =====")
	fmt.Println("==== With aliased Prints ====")

	var pf = fmt.Printf

	instructions := [3]string{"first", "middle", "last"}
	inputs := [3]string{}
	scanner := bufio.NewScanner(os.Stdin)

	for idx, val := range instructions {
		pf("Please enter %q name: ", val)
		scanner.Scan()
		input := scanner.Text()
		inputs[idx] = input
	}

	pf("Entered name is: %s %s %s\n", inputs[0], inputs[1], inputs[2])

}
