package main

import (
	"fmt"
	"os"

	"github.com/arvpyrna/sled/internal/db"
)

func showOptions() {
	options, err := os.ReadFile("res/options.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(options))
}

func getInput() {
	for {
		showOptions()
		var id int
		_, err := fmt.Scanf("%d", &id)
		if err != nil {
			fmt.Println("Error: Please enter a valid digit")
		}

		switch id {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
			performCleanup()
			os.Exit(0)
		}
	}
}

func performCleanup() {
	fmt.Println("Quitting...")
}

func main() {
	db.Init()
	getInput()
}
