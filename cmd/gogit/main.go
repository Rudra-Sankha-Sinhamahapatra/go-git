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

	case "write-tree":
		treeHash, err := repo.WriteTree(".", ".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("root tree hash:", treeHash)

	case "commit-tree":
		if len(os.Args) < 4 || os.Args[2] != "-m" {
			fmt.Println("usage: gogit commit-tree -m <message>")
			return
		}

		message := os.Args[3]

		ref, err := repo.ResolveHEAD(".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		parentHash, exists, err := repo.ReadCurrentCommit(".gogit/", ref)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		treeHash, err := repo.WriteTree(".", ".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		var parentPtr *string
		if exists {
			parentPtr = &parentHash
		}

		rawCommit, err := repo.CreateCommitObject(treeHash, parentPtr, message)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		commitHash, err := repo.WriteCommitObject(rawCommit, ".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if err := repo.UpdateRef(".gogit/", ref, commitHash); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("commit hash:", commitHash)

	case "head":
		ref, err := repo.ResolveHEAD(".gogit/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("HEAD points to:", ref)

		hash, exists, err := repo.ReadCurrentCommit(".gogit/", ref)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if exists {
			fmt.Println("current commit:", hash)
		} else {
			fmt.Println("no commits yet")
		}

	case "log":
		err := repo.LogCommits(".gogit", ".")
		if err != nil {
			fmt.Println("error:", err)
		}

	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
