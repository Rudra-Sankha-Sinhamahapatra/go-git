package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

/*
 Reads file + print size
 @HashObjectSizeRead
*/

func HashObjectSizeRead(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	fmt.Printf("read %d bytes from %s\n", len(data), path)

	return nil
}

/*
 Creates raw go git object (header + content)
 @CreateRawGoGitObject
*/

func CreateRawGoGitObject(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	size := len(data)

	header := fmt.Sprintf("blob %d\x00", size)

	raw := append([]byte(header), data...)

	fmt.Println("header: ", header)
	fmt.Println("raw object size: ", len(raw))

	return raw, nil
}

/*
  Computes the hash of a raw git object (header + content)
  @ComputeGoGitObjectHash
*/

func ComputeGoGitObectHash(raw []byte) (string, error) {
	if len(raw) == 0 {
		return "", fmt.Errorf("raw object is empty")
	}

	sum := sha1.Sum(raw)

	hash := hex.EncodeToString(sum[:])

	fmt.Println("object hash:", hash)

	return hash, nil
}
