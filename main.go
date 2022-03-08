package main

import (
	"fmt"
	"os"
	"time"

	"github.com/arvpyrna/sled/internal/db"
	"github.com/schollz/progressbar/v3"
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
			startTimer(3)
		case 3:
		case 4:
		case 5:
		case 6:
		case 9:
			performCleanup()
			os.Exit(0)
		}
	}
}

// Timer with interactive progress bar
// Gets input in minute but shows progress bar in the unit of seconds
func startTimer(minutes int) {
	durationInSec := minutes * 60
	fmt.Println(fmt.Sprintf("Timer starting for %d minutes", minutes))
	bar := progressbar.Default(int64(durationInSec))
	for i := 0; i < durationInSec; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}

func performCleanup() {
	fmt.Println("Quitting...")
}

func main() {
	db.Init()
	getInput()
}
