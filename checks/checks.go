package checks

import (
	"my-ls/structures"
	"os"
)

// Checks if path exists
func CheckPath(path string) bool {
	err := os.Chdir(path)
	if err != nil {
		os.Chdir(structures.STARTDIR)
		return false
	}

	return true
}

// Checks if folder or file is hidden
func IsHidden(filename string) bool {
	if filename[0:1] == "." {
		return true
	}
	return false
}
