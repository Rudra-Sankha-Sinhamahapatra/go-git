package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
  Reads the compressed object from .gogit/objects/<dir>/<file>
  @ReadCompressedGoGitObject
*/

func ReadCompressedGoGitObject(hash string, repoPath string) ([]byte, error) {
	if len(hash) < 2 {
		return nil, fmt.Errorf("invalid hash")
	}

	dir := hash[:2]
	file := hash[2:]

	objectPath := filepath.Join(repoPath, ".gogit", "objects", dir, file)

	compressedData, err := os.ReadFile(objectPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read compressed object %s: %w", hash, err)
	}

	fmt.Println("compressed object size:", len(compressedData))
	fmt.Println("object found at:", objectPath)

	return compressedData, nil

}
