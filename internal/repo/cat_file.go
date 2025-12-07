package repo

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

/*
  Decompresses a zlib-compressed git object
  @DecompressGoGitObject
*/

func DecompressesGoGitObject(compressed []byte) ([]byte, error) {
	if len(compressed) == 0 {
		return nil, fmt.Errorf("compressed data is empty")
	}

	b := bytes.NewReader(compressed)

	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, fmt.Errorf("failed to create zlib reader: %w", err)
	}
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read decompressed data: %w", err)
	}

	fmt.Println("decompressed size:", len(decompressed))

	return decompressed, nil
}

/*
  Parses a raw git object into (type, size, content)
  @ParseGoGitObjectHeader
*/

func ParseGoGitObjectHeader(raw []byte) (string, int, []byte, error) {
	nullIndex := bytes.IndexByte(raw, 0)
	if nullIndex == -1 {
		return "", 0, nil, fmt.Errorf("invalid git object: no header found")
	}

	header := string(raw[:nullIndex])

	parts := strings.SplitN(header, " ", 2)
	if len(parts) < 2 {
		return "", 0, nil, fmt.Errorf("invalid header format")
	}

	objectType := parts[0]

	size, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, nil, fmt.Errorf("invalid size in header : %w", err)
	}

	content := raw[nullIndex+1:]

	fmt.Println("object type:", objectType)
	fmt.Println("declared size:", size)
	fmt.Println("actual content size:", len(content))

	return objectType, size, content, nil
}
