//Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
//  7 and 11 are to be called NATURAL
//  2 is called SNAKE-EYES-CRAPS
//  3 and 12 is called LOSS-CRAPS
//  Any other combination is called NEUTRAL.


package main

import(
    "fmt"
    "math/rand"
    "time"
)
func assign_dices(result int) string {
    switch result {
    case 7,11:
        return "NATURAL"
    case 2:
        return "SNAKE-EYED-CRAPS"
    case 3,12:
        return "LOSS-CRAPS"
    default:
        return "NEUTRAL"
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 6

    for i:=0; i<50; i++ {
        dice1 := rand.Intn(max-min+1) + min
        dice2 := rand.Intn(max-min+1) + min

        fmt.Printf("Dices rolled: %d and %d, it gives: %d\t %s\n", dice1, dice2, dice1+dice2, assign_dices(dice1+dice2)) 
    }

}
