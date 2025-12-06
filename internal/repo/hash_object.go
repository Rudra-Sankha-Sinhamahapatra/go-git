package repo

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
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

/*
  Writes a compressed git object into .gogit/objects/ directory
  @WriteGoGitObject
*/

func WriteGoGitObject(hash string, raw []byte, repoPath string) error {
	if len(hash) < 2 {
		return fmt.Errorf("invalid hash")
	}

	dir := hash[:2]

	file := hash[2:]

	objectDir := filepath.Join(repoPath, "objects", dir)
	objectPath := filepath.Join(objectDir, file)

	if err := os.MkdirAll(objectDir, 0755); err != nil {
		return fmt.Errorf("cannot create object directory: %w", err)
	}

	f, err := os.Create(objectPath)
	if err != nil {
		return fmt.Errorf("cannot create object file: %w", err)
	}

	defer f.Close()

	w := zlib.NewWriter(f)

	_, err = w.Write(raw)

	if err != nil {
		return fmt.Errorf("cannot write compressed object: %w", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("cannot close compressed writer: %w", err)
	}

	fmt.Println("object written to:", objectPath)

	return nil
}
