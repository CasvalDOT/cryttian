package main

import (
	"cryttian/configurator"
	"fmt"
	"os"
)

func listThemes() {
	configurator, err := configurator.NewConfigurator()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	files, err := configurator.ListThemes()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for index, file := range files {
		fmt.Printf("[%d] %s\n", index+1, file)
	}
}

func applyTheme() {
	configurator, err := configurator.NewConfigurator()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	themeName := ""
	if len(os.Args) <= 2 {
		themeName, err = configurator.SelectTheme()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		themeName = os.Args[2]
	}

	err = configurator.ApplyTheme(themeName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Please provide a valid verb")
		fmt.Println("1. list")
		fmt.Println("2. apply <theme>")
		os.Exit(1)
	}

	if args[1] == "list" {
		listThemes()
	} else if args[1] == "apply" {
		applyTheme()
	} else {
		fmt.Println("Invalid verb")
		os.Exit(1)
	}
}
