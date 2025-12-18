package repo

import (
	"fmt"
	"os"
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
