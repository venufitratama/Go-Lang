package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Admin"
	}
	registerUser("Admin", blacklist)
	registerUser("Venu", blacklist)

	//atau bisa langsung seperti ini
	registerUser("Root", func(name string) bool {
		return name == "Root"
	})
	registerUser("Admin", func(name string) bool {
		return name == "Admin"
	})
}
