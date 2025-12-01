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

	case "hash-object":
		if len(os.Args) < 3 {
			fmt.Println("usage: gogit hash-object <file>")
			return

		}
		err := repo.HashObjectSizeRead(os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		_, err = repo.CreateRawGoGitObject(os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
