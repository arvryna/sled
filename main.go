package main

import (
	"fmt"

	"github.com/arvpyrna/sled/internal/db"
)

func showOptions() {

}

func main() {
	showOptions()
	fmt.Println("Starting sled ...")
	db.Init()
}
