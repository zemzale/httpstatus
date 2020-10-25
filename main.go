package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	TABLE_PADDING int = 4
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

func renderCode(code HTTPCode, width int) {
	headerTable := tablewriter.NewWriter(os.Stdout)
	headerTable.SetHeader([]string{"Code", "Summary"})
	headerTable.Append([]string{strconv.Itoa(code.Code), code.Name})
	headerTable.Render()

	descriptionTable := tablewriter.NewWriter(os.Stdout)
	descriptionTable.SetColWidth(width)
	descriptionTable.SetHeader([]string{"Description"})
	descriptionTable.Append([]string{code.Description})
	descriptionTable.Render()
}

func terminalWidth() (int, error) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return 0, err
	}
	return width, nil
}

func tableWidth() (int, error) {
	fullWidth, err := terminalWidth()
	if err != nil {
		return 0, err
	}
	return fullWidth - TABLE_PADDING, nil
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
			width, err := tableWidth()
			if err != nil {
				width = 80
			}
			renderCode(code, width)
			os.Exit(0)
		}
	}

	fmt.Println("No code with this number")
	os.Exit(1)
}
