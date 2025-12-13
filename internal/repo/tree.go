package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
  Lists all files and directories in the current folder.
  (First micro-step of write-tree)
  @ListDirectoryEntries
*/

func ListDirectoryEntries(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("cannot read directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(path, name)

		if name == ".gogit" || name == ".git" {
			continue
		}

		if entry.IsDir() {
			fmt.Println("dir:", name)
		} else {
			hash, err := ComputeBlobHashForFile(fullPath)
			if err != nil {
				return err
			}
			fmt.Printf("file: %s -> blob %s\n", name, hash)
		}

	}

	return nil
}

/*
  Computes blob hash for a file WITHOUT writing object
  @ComputeBlobHashForFile
*/

func ComputeBlobHashForFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read file: %w", err)
	}

	size := len(data)
	header := fmt.Sprintf("blob %d\x00", size)
	raw := append([]byte(header), data...)

	hash, err := ComputeGoGitObectHash(raw)
	if err != nil {
		return "", err
	}

	return hash, nil
}

/*
  Recursively walks the working directory
  (Write-tree micro-step: recursion only)
  @WalkWorkingDirectory
*/

func WalkWorkingDirectory(path string, depth int) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("cannot read directory %s: %w", path, err)
	}

	prefix := strings.Repeat(" ", depth)

	for _, entry := range entries {
		name := entry.Name()

		if name == ".git" || name == ".gogit" {
			continue
		}

		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			fmt.Println(prefix+"dir:", name)

			err := WalkWorkingDirectory(fullPath, depth+1)
			if err != nil {
				return err
			}
		} else {
			hash, err := ComputeBlobHashForFile(fullPath)
			if err != nil {
				return err
			}
			fmt.Println(prefix+"file:", name, "-> blob", hash)
		}
	}
	return nil
}
