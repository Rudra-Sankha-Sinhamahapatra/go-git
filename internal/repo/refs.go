package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
  Resolves HEAD to a ref path
  Example: "ref: refs/heads/master" -> "refs/heads/master"
  @ResolveHEAD
*/

func ResolveHEAD(repoPath string) (string, error) {
	headPath := repoPath + "/HEAD"

	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", fmt.Errorf("cannot read HEAD: %w", err)
	}

	content := strings.TrimSpace(string(data))

	if !strings.HasPrefix(content, "ref: ") {
		return "", fmt.Errorf("detached HEAD not supported yet")
	}

	ref := strings.TrimPrefix(content, "ref: ")
	return ref, nil
}

/*
  Reads current commit hash from a ref
  @ReadCurrentCommit
*/

func ReadCurrentCommit(repoPath string, ref string) (string, bool, error) {
	refPath := repoPath + "/" + ref

	data, err := os.ReadFile(refPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false, nil
		}
		return "", false, fmt.Errorf("cannot read ref: %w", err)
	}

	commitHash := strings.TrimSpace(string(data))
	return commitHash, true, nil
}

/*
  Updates a ref to point to a commit hash
  Example: refs/heads/main -> <commitHash>
  @UpdateRef
*/

func UpdateRef(repoPath string, ref string, commitHash string) error {
	refPath := repoPath + "/" + ref

	if err := os.MkdirAll(filepath.Dir(refPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(refPath, []byte(commitHash+"\n"), 0644)
}
