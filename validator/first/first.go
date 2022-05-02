package first

import (
		"fmt"
		"os"
		"log"
		"bufio"
		"strings"
)

//This function takes in a file selected by the user and makes sure each line doesn't exceed 100 characters.
//Help for reading files from: https://www.dotnetperls.com/readdir-go
func FirstStandard() {
	fmt.Println("<--------------------------->")
	fmt.Println("Standard: 100 character rule.")
	fmt.Println("<--------------------------->")
	var fname string
	var totalerrors int = 0
	var currentline int = 0
	var lines []int
	directory, err := os.Getwd()
	directoryread, _ := os.Open(directory)
	directoryoutput, _ := directoryread.ReadDir(0)
	fmt.Println("Current files: ")
	for index := range(directoryoutput) {
		filehere := directoryoutput[index]
		filenamehere := filehere.Name()
		fmt.Println("\t", filenamehere)
	}
	fmt.Println()
	defer directoryread.Close()
	fmt.Println("Which file would you like to read?")
	if _, err := fmt.Scan(&fname); err != nil {
		fmt.Println(err)
	}
	content, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
        number := len(scanner.Text())
		currentline = currentline + 1
		if (number > 100) {
			totalerrors = totalerrors + 1
			lines = append(lines, currentline)
		}
    }
	if (totalerrors > 0) {
		var start string
		var choice string
		fmt.Println("The file violates the 100 character rule!")
		fmt.Println("Would you like further details? If so, input val detail when prompted.")
		fmt.Scan(&start)
		fmt.Scan(&choice)
		if (choice == "detail") {
			fmt.Println("It violated the standard", totalerrors, "times.")
			fmt.Println("The lines where it did so were: ", lines)
			fmt.Println()
			fmt.Println("Thanks for using this program!")
		} else {
			fmt.Println("Thanks for using this program!")
		}
	} else {
		fmt.Println("Standard passed! No lines have more than 100 characters.")
	}
}

func FirstStandardAuto() {
	fmt.Println("<--------------------------->")
	fmt.Println("Standard: 100 character rule.")
	fmt.Println("<--------------------------->")
	var totalerrors int = 0
	var currentline int = 0
	var lines []int
	var gofiles []string
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	directoryread, _ := os.Open(directory)
	directoryoutput, _ := directoryread.ReadDir(0)
	for index := range(directoryoutput) {
		filehere := directoryoutput[index]
		filenamehere := filehere.Name()
		if strings.HasSuffix(filenamehere, "go") {
			gofiles = append(gofiles, filenamehere)
		}
	}
	fmt.Println()
	defer directoryread.Close()
	tracker := len(gofiles)
	for (tracker > 0) {
		fmt.Println("Reading file: \n", gofiles[tracker - 1])
		content, err := os.Open(gofiles[tracker - 1])
    	if err != nil {
    		log.Fatal(err)
    	}
		defer content.Close()
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
	        number := len(scanner.Text())
			currentline = currentline + 1
			if (number > 100) {
				totalerrors = totalerrors + 1
				lines = append(lines, currentline)
			}
	    }
		if (totalerrors > 0) {
			fmt.Println("The file violates the 100 character rule!")
			fmt.Println("It violated the standard", totalerrors, "times.")
			fmt.Println("The lines where it did so were: ", lines)
			fmt.Println()
		} else {
			fmt.Println("Standard passed! No lines have more than 100 characters.")
		}
		tracker = tracker - 1
	}
}