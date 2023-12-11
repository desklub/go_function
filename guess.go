package main

import (
	"fmt"
	"math/rand"
)

var initialScore int = 100
var point int = 20
var guess int
var correctCount int
var wrongCount int
var correct string = "Wow! Correct"
var wrong string = "You guessed wronged!"
var count int
var randomNumList []int
var guessNumList []int

func main() {
	fmt.Println("Welcome to the guess game! You have", initialScore, "Nitendo")
	fmt.Println("Each corrrect or wrong guess adds or removes", point, "from your Nitendo respectively")
	fmt.Println("Guess a Number between 1 to 10 or 0 to quit.")
	for count = 0; initialScore > 0; count++ {
	var	randomNumber = rand.Intn(11)

		fmt.Print("(", count+1, ") ", "Guess Number >>> \n ")
		fmt.Scan(&guess)
		if guess == randomNumber {
			initialScore += point
			fmt.Println(correct, "Your Nitendo is now:", initialScore)
			correctCount++
		}
		if guess != randomNumber {
			initialScore -= point
			fmt.Println(wrong, "Your Nitendo is now:", initialScore)
			wrongCount++
		}
		if guess == 0 {
			fmt.Println("Quiting....")
			break
		} 
		if guess == randomNumber && guess != randomNumber && guess <= 0 {
			fmt.Println("Please enter a number")
			continue
		}
		randomNumList = append(randomNumList, randomNumber)
		guessNumList = append(guessNumList, guess)
	}
	fmt.Printf("\nYour have guessed %d times\nCorrect guess - %d\nWrong guess - %d\nYour Final Credit is:%d\n", count, correctCount, wrongCount, initialScore)
	fmt.Println("Random Numbers  -", randomNumList)
	fmt.Println("Guessed Numbers -", guessNumList)
}