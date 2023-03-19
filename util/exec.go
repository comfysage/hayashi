package util

func CloneCommand(url string, output string) []string {
	return []string{"git", "clone", "--filter=blob:none", url, output}
}
