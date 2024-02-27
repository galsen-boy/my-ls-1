package flags

import (
	"fmt"
	"my-ls/checks"
	"my-ls/r_flag"
	"my-ls/structures"
	"my-ls/t_flag"
	"os"
	"strings"
)

// Collects all given by user flags, filenames, folders and saves them into slices
func CollectAllAgruments(arguments []string) (structures.Flags, []string, []string, []string, bool) {
	var flagsToUse structures.Flags
	var paths []string
	var files []string
	var folders []string

	var wasPath bool

	inputArgs := arguments[1:]

	for i := 0; i < len(inputArgs); i++ {
		if inputArgs[i][:1] == "/" {
			if checks.CheckPath(inputArgs[i]) {
				paths = append(paths, inputArgs[i])
				continue
			} else {
				fmt.Printf("ls: cannot access '%s': No such file or directory\n", inputArgs[i])
				wasPath = true
				continue
			}
		}

		if strings.Contains(inputArgs[i], "/") {
			if checks.CheckPath(structures.STARTDIR + "/" + inputArgs[i]) {
				folders = append(folders, structures.STARTDIR+"/"+inputArgs[i])
				continue
			} else {
				fmt.Printf("ls: cannot access '%s': Not a directory\n", inputArgs[i])
				wasPath = true
				continue
			}
		}

		if !strings.Contains(inputArgs[i], "-") || inputArgs[i] == "-" {
			fileInfo, err := os.Lstat(inputArgs[i])

			if err == nil {
				if fileInfo.Mode()&os.ModeSymlink == 0 {
					if checks.CheckPath(structures.STARTDIR + "/" + inputArgs[i]) {
						folders = append(folders, structures.STARTDIR+"/"+inputArgs[i])
						continue
					}
				}
			}

			files = append(files, inputArgs[i])
			continue
		}

		for _, k := range inputArgs[i] {
			if k == '-' {
				flagsToUse = DetectFlag(inputArgs[i], flagsToUse)
				break
			} else {
				paths = append(paths, inputArgs[i])
				break
			}
		}
	}

	return flagsToUse, paths, files, folders, wasPath
}

// Detects if flags from user are right
func DetectFlag(flagToCheck string, flagsToUse structures.Flags) structures.Flags {
	dashFound := false
	for _, i := range flagToCheck {
		if i == 'R' && flagsToUse.Flag_R == false {
			flagsToUse.Flag_R = true
		} else if i == 'a' && flagsToUse.Flag_a == false {
			flagsToUse.Flag_a = true
		} else if i == 'r' && flagsToUse.Flag_r == false {
			flagsToUse.Flag_r = true
		} else if i == 'l' && flagsToUse.Flag_l == false {
			flagsToUse.Flag_l = true
		} else if i == 't' && flagsToUse.Flag_t == false {
			flagsToUse.Flag_t = true
		} else if i == '-' && !dashFound {
			dashFound = true
		} else {
			fmt.Println("ERROR")
			os.Exit(0)
		}
	}

	return flagsToUse
}

// Applies flag to collected data
func ApplyFlags(flagsToUse structures.Flags, contentList []structures.FileData) []structures.FileData {
	if flagsToUse.Flag_t == true {
		t_flag.SortByTime(contentList)
	}
	if flagsToUse.Flag_r == true {
		contentList = r_flag.ReverseList(contentList)
	}
	return contentList
}
