package main

//Created by Connor Thompson
import (
	"fmt"
	"example.com/validator/first"
	"example.com/validator/fourth"
	"example.com/validator/second"
	"example.com/validator/third"
)

//Main validator function, which serves as the users directory.
func main() {
	var start string
	var choice string
	fmt.Println("<-------------------------------------------------------------------------------------------->")
	fmt.Println("Welcome to the Go Validator!")
	fmt.Println("To see the commands, type val help when prompted.")
	fmt.Println("Otherwise, type val start to execute all standards at once, as well as scanning each .go file.")
	fmt.Println("<--------------------------------------------------------------------------------------------->")
	if _, err := fmt.Scan(&start); err != nil {
		fmt.Println("Invalid selection!")
	}
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Invalid selection!")
	}
	if choice == "help" {
		//Directs to each of the go functions depending on user input
		fmt.Println()
		fmt.Println("To execute the first coding standard, 100 characters per line max, type val first when prompted.")
		fmt.Println("To execute the second coding standard, README.md correct formatting, type val second when prompted.")
		fmt.Println("To execute the third coding standard, presence of license, type val third when prompted.")
		fmt.Println("To execute the fourth coding standard, included comments, type val fourth when prompted.")
		fmt.Println()
		if _, err := fmt.Scan(&start); err != nil {
			fmt.Println("Invalid selection!")
		}
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid selection!")
		}
		if choice == "first" {
			first.FirstStandard()
		} else if choice == "second" {
			second.SecondStandard()
		} else if choice == "third" {
			third.ThirdStandard()
		} else if choice == "fourth" {
			fourth.FourthStandard()
		} else {
			fmt.Println("Something went wrong!")
		}
	} else if choice == "start" {
		//Does all of the standard checks in an accelerrated manner, specifically watching for go files.
		first.FirstStandardAuto()
		second.SecondStandardAuto()
		third.ThirdStandard()
		fourth.FourthStandardAuto()
	} else {
		fmt.Println("Choice not valid!")
	}
}
