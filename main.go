package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "debug" {
		debug()
		os.Exit(0)
	}

	inputCode, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid code")
		os.Exit(1)
	}

	if inputCode < 99 || 600 < inputCode {
		fmt.Println("No code with this number")
		os.Exit(1)
	}

	for _, code := range HTTPCodes {
		if code.Code == inputCode {
			fmt.Printf("Code : %d\n", code.Code)
			fmt.Printf("Name : %s\n", code.Name)
			fmt.Printf("Description : %v\n", code.Description)
			os.Exit(0)
		}
	}

	fmt.Println("No code with this number")
	os.Exit(1)
}

func debug() {
	f, err := os.Create("codeStatus.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, code := range HTTPCodes {
		fmt.Printf("Code : %d\n", code.Code)
		fmt.Printf("Name : %s\n", code.Name)
		fmt.Printf("Description : %v\n", code.Description)
	READAGAIN:
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		switch char {
		case 'y':
			f.WriteString(fmt.Sprintf("%d : good\n", code.Code))
		case 'n':
			f.WriteString(fmt.Sprintf("%d : bad\n", code.Code))
		default:
			goto READAGAIN
		}
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}
