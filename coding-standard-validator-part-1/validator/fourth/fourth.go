package fourth

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FourthStandard() {
	fmt.Println("<------------------------>")
	fmt.Println("Standard: enough comments.")
	fmt.Println("<------------------------>")
	var fname string
	var commentcount int = 0
	directory, err := os.Getwd()
	if err != nil {
        log.Fatal(err)
    }
	directoryread, _ := os.Open(directory)
	directoryoutput, _ := directoryread.ReadDir(0)
	fmt.Println("Current files: ")
	for index := range directoryoutput {
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
		trimmedtext := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(trimmedtext, "//") {
			commentcount = commentcount + 1
		}
	}
	if commentcount < 1 {
		var start string
		var choice string
		fmt.Println("The file violates the comments rule!")
		fmt.Println("Would you like further details? If so, input val detail when prompted.")
		fmt.Scan(&start)
		fmt.Scan(&choice)
		if choice == "detail" {
			fmt.Println("It had no comments!")
			fmt.Println()
			fmt.Println("Thanks for using this program!")
		} else {
			fmt.Println("Thanks for using this program!")
		}
	} else if commentcount < 3 {
		var start string
		var choice string
		fmt.Println("The file violates the comments rule!")
		fmt.Println("Would you like further details? If so, input val detail when prompted.")
		fmt.Scan(&start)
		fmt.Scan(&choice)
		if choice == "detail" {
			fmt.Println("It had only", commentcount, "comment(s)!")
			fmt.Println()
			fmt.Println("Thanks for using this program!")
		} else {
			fmt.Println("Thanks for using this program!")
		}
	} else {
		fmt.Println("Standard passed! There are enough comments.")
	}
}

func FourthStandardAuto() {
	fmt.Println("<------------------------>")
	fmt.Println("Standard: enough comments.")
	fmt.Println("<------------------------>")
	var commentcount int = 0
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
			trimmedtext := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(trimmedtext, "//") {
			commentcount = commentcount + 1
		}
	}
	if commentcount < 1 {
		fmt.Println("The file violates the comments rule!")
		fmt.Println("It had no comments!")
	} else if commentcount < 3 {
		fmt.Println("The file violates the comments rule!")
		fmt.Println("It had only", commentcount, "comment(s)!")
	} else {
		fmt.Println("Standard passed! There are enough comments.")
	}
	tracker = tracker - 1
	}
}