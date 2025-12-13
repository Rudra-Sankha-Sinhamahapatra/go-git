package repo

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

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
  Builds a single tree entry in Git format:
  <mode> <name>\x00<20-byte binary hash>
  @BuildTreeEntry
*/

func BuildTreeEntry(mode string, name string, hexHash string) ([]byte, error) {
	hashBytes, err := hex.DecodeString(hexHash)
	if err != nil {
		return nil, fmt.Errorf("invalid hex hash: %w", err)
	}

	if len(hashBytes) != 20 {
		return nil, fmt.Errorf("invalid hash length: expected 20 bytes")
	}

	header := fmt.Sprintf("%s %s\x00", mode, name)

	entry := append([]byte(header), hashBytes...)

	return entry, nil
}

/*
  Sorts tree entries lexicographically by name
  @SortTreeEntries
*/

func SortTreeEntries(entries [][]byte) {
	sort.Slice(entries, func(i, j int) bool {
		return extractTreeEntryName(entries[i]) < extractTreeEntryName(entries[j])
	})
}

/*
  Extracts filename from a tree entry
  @extractTreeEntryName
*/

func extractTreeEntryName(entry []byte) string {
	spaceIndex := bytes.IndexByte(entry, ' ')
	if spaceIndex == -1 {
		return ""
	}

	nameStart := spaceIndex + 1

	nullIndex := bytes.IndexByte(entry[nameStart:], 0)
	if nullIndex == -1 {
		return ""
	}

	return string(entry[nameStart : nameStart+nullIndex])
}

// Debug helper
func DebugTreeEntryName(entry []byte) string {
	return extractTreeEntryName(entry)
}

/*
  Builds, hashes, and writes a tree object
  (non-recursive, single directory)
  @WriteTreeObject
*/

func WriteTreeObject(treeEntries [][]byte, repoPath string) (string, error) {
	var content []byte
	for _, entry := range treeEntries {
		content = append(content, entry...)
	}

	header := fmt.Sprintf("tree %d\x00", len(content))
	rawTree := append([]byte(header), content...)

	treeHash, err := ComputeGoGitObectHash(rawTree)
	if err != nil {
		return "", err
	}

	err = WriteGoGitObject(treeHash, rawTree, repoPath)
	if err != nil {
		return "", err
	}

	fmt.Println("tree written with hash:", treeHash)
	return treeHash, nil
}

/*
  Recursively builds tree objects and returns tree hash
  @WriteTree
*/

func WriteTree(path string, repoPath string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("cannot read directory %s: %w", path, err)
	}

	var treeEntries [][]byte

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(path, name)

		if name == ".git" || name == ".gogit" {
			continue
		}

		if entry.IsDir() {
			subTreeHash, err := WriteTree(fullPath, repoPath)
			if err != nil {
				return "", err
			}

			treeEntry, err := BuildTreeEntry("40000", name, subTreeHash)
			if err != nil {
				return "", err
			}

			treeEntries = append(treeEntries, treeEntry)
		} else {
			blobHash, err := ComputeBlobHashForFile(fullPath)
			if err != nil {
				return "", err
			}

			treeEntry, err := BuildTreeEntry("100644", name, blobHash)
			if err != nil {
				return "", err
			}

			treeEntries = append(treeEntries, treeEntry)
		}
	}

	SortTreeEntries(treeEntries)

	return WriteTreeObject(treeEntries, repoPath)
}
