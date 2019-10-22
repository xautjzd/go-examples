package version

import "fmt"

var (
	gitBranch string
	gitCommit string
	built     string
	goVersion string
)

func Version() string {
	return fmt.Sprintf("branch: %s, commit: %s, built: %s, go: %s", gitBranch, gitCommit, built, goVersion)
}
