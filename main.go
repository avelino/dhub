package main

import (
	"fmt"

	"github.com/avelino/dhub/repositories"
	"github.com/avelino/dhub/users"
)

func main() {
	a := users.Auth{
		Username: "",
		Password: "",
	}
	fmt.Println(repositories.BuildHistory(a, "nuveo", "runners"))
}
