package main

import (
	"fmt"
	"os"
	"walk-the-camino/app"
)

func main() {
	if err := app.Start(); err != nil {
		fmt.Println("Error while starting application ", err.Error())
		os.Exit(1)
	}
}
