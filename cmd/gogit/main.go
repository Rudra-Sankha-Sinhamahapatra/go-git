package main

import (
	"fmt"
	"os"

	"github.com/Rudra-Sankha-Sinhamahapatra/go-git/internal/repo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gogit <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		err := repo.InitRepo(".")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("Initialized empty gogit repository in .gogit/")
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
