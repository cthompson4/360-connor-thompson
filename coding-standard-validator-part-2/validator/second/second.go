package second

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)


func SecondStandard() {
	fmt.Println("<----------------------------->")
	fmt.Println("Standard: included README file.")
	fmt.Println("<----------------------------->")
	var totalincluded int = 0
	content, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Project") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Email") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Credit") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Purpose") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Requirements") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Build instructions") {
			totalincluded = totalincluded + 1
		}
	}
	if totalincluded < 7 {
		var start string
		var choice string
		fmt.Println("The file violates the README.md guidelines!")
		fmt.Println("Would you like further details? If so, input val detail when prompted.")
		fmt.Scan(&start)
		fmt.Scan(&choice)
		if choice == "detail" {
			fmt.Println("It only had", totalincluded, "of the guidelines!")
			fmt.Println()
			fmt.Println("Thanks for using this program!")
		} else {
			fmt.Println("Thanks for using this program!")
		} 
	} else {
		fmt.Println("The README.md file passes the test!")
	}
}

func SecondStandardAuto() {
	fmt.Println("<----------------------------->")
	fmt.Println("Standard: included README file.")
	fmt.Println("<----------------------------->")
	var totalincluded int = 0
	content, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Project") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Email") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Credit") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Purpose") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Requirements") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Build instructions") {
			totalincluded = totalincluded + 1
		}
	}
	if totalincluded < 7 {
		fmt.Println("The file violates the README.md guidelines!")
		fmt.Println("It only had", totalincluded, "of the guidelines!")
		fmt.Println()
	} else {
		fmt.Println("The README.md file passes the test!")
	}
}

func SecondStandardTest() int {
	var totalincluded int = 0
	content, err := os.Open("test-statement-coverage/testREADME.md")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Project") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Email") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Credit") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Purpose") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Requirements") {
			totalincluded = totalincluded + 1
		}
		if strings.HasPrefix(scanner.Text(), "Build instructions") {
			totalincluded = totalincluded + 1
		}
	}
	if totalincluded < 7 {
		return 0
	} else {
		return 1
	}
}

