package main

import (
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("failed load .env")
	}
}

func main(){
	
}
