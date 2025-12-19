package repo

import (
	"fmt"
	"time"
)

/*
  Builds raw commit object (no parent)
  @CreateCommitObject
*/

func CreateCommitObject(treeHash string, parent *string, message string) ([]byte, error) {
	author := "Rudra Sankha Sinhamahapatra <rudra@rudrasankha.com>"
	committer := author

	timestamp := time.Now().Unix()
	timezone := "+0530"

	var parentLine string
	if parent != nil {
		parentLine = fmt.Sprintf("parent %s\n", *parent)
	}

	content := fmt.Sprintf(
		"tree %s\n"+
			"%s"+
			"author %s %d %s\n"+
			"committer %s %d %s\n\n"+
			"%s\n",
		treeHash,
		parentLine,
		author, timestamp, timezone,
		committer, timestamp, timezone,
		message,
	)

	header := fmt.Sprintf("commit %d\x00", len(content))
	raw := append([]byte(header), []byte(content)...)

	return raw, nil
}

/*
  Writes commit object and returns commit hash
  @WriteCommitObject
*/

func WriteCommitObject(raw []byte, repoPath string) (string, error) {
	commitHash, err := ComputeGoGitObectHash(raw)
	if err != nil {
		return "", err
	}

	err = WriteGoGitObject(commitHash, raw, repoPath)
	if err != nil {
		return "", err
	}

	fmt.Printf("wrote commit object %s\n", commitHash)
	return commitHash, nil
}
