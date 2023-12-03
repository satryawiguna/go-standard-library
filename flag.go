package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your hostname")
	username := flag.String("u", "root", "Put your username")
	password := flag.String("pwd", "", "Put your password")
	port := flag.Int("p", 0, "Put your database port")

	flag.Parse()

	fmt.Printf("Your access credential %s %s %s", *host, *username, *password, *port)
	fmt.Println("hostname", host)
	fmt.Println("username", username)
	fmt.Println("password", password)
	fmt.Println("port", port)
}
