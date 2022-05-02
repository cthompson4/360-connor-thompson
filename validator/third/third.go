package third

import (
	"fmt"
	"os"
)

func ThirdStandard() {
	fmt.Println("<------------------------------>")
	fmt.Println("Standard: included LICENSE file.")
	fmt.Println("<------------------------------>")
	var yes int = 0
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	directoryread, _ := os.Open(directory)
	directoryoutput, _ := directoryread.ReadDir(0)
	for index := range directoryoutput {
		filehere := directoryoutput[index]
		filenamehere := filehere.Name()
		if filenamehere == "LICENSE.lic" {
			yes = yes + 1
		}
	}
	defer directoryread.Close()
	if yes == 1 {
		fmt.Println("There is a license file included, so it passes.")
	} else {
		fmt.Println("There is no license file! Fail.")
	}
}
