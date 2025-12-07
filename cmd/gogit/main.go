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

		raw, err := repo.CreateRawGoGitObject(os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		hash, err := repo.ComputeGoGitObectHash(raw)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		err = repo.WriteGoGitObject(hash, raw, ".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

	case "cat-file":
		if len(os.Args) < 4 || os.Args[2] != "-p" {
			fmt.Println("usage: gogit cat-file -p <hash>")
			return
		}

		hash := os.Args[3]

		compressed, err := repo.ReadCompressedGoGitObject(hash, ".")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		decompressed, err := repo.DecompressesGoGitObject(compressed)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		_, _, content, err := repo.ParseGoGitObjectHeader(decompressed)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		repo.PrintGoGitObjectContent(content)

	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
