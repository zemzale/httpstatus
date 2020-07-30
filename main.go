package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func renderSummary() {
	headers := []string{"Code", "Summary"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	for _, code := range HTTPCodes {
		table.Append([]string{strconv.Itoa(code.Code), code.Name})
	}

	table.Render()
}

func renderCode(code HTTPCode) {
	headerTable := tablewriter.NewWriter(os.Stdout)
	headerTable.SetHeader([]string{"Code", "Summary"})
	headerTable.Append([]string{strconv.Itoa(code.Code), code.Name})
	headerTable.Render()

	//TODO @zemzale 30/07/20 Set table width to be the size of the terminal
	descriptionTable := tablewriter.NewWriter(os.Stdout)
	descriptionTable.SetHeader([]string{"Description"})
	descriptionTable.Append([]string{code.Description})
	descriptionTable.Render()
}

func main() {
	if len(os.Args) == 1 {
		renderSummary()
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
			renderCode(code)
			os.Exit(0)
		}
	}

	fmt.Println("No code with this number")
	os.Exit(1)
}
