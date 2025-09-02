package main

import (
	"cli_notes/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//cmd.Execute()
	id, err := service.Add("sleep", "sleep at 10 o`clock", "home")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(id)
}
