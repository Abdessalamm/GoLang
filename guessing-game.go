package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Game: Guess a number between 0 and 10")
	fmt.Println("You have 3 tries")

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretnumber := randomizer.Intn(10)

	var guess int

	for try :=1; try <= 3; try++{
		fmt.Println("Trial %d\n", try)
		fmt.Println("Please enter your number")
		fmt.Scan(&guess)
		if guess < secretnumber {
			fmt.Println("Sorry, wrong guess; number is too small \n ")
		} else if guess < secretnumber{
			fmt.Println("Sorry, wrong guess; number is too Large \n ")

		} else {
			fmt.Printf("You win! \n")
			break
		}
		if try == 3 {
			fmt.Printf(
			"Game Over!! \n")
			fmt.Printf("The Correct number is %d\n", secretnumber)
			break
		}
	}

}
