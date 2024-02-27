package data_interaction

import (
	"fmt"
	"log"
	"my-ls/calculations"
	"my-ls/checks"
	"my-ls/flags"
	"my-ls/sorts"
	"my-ls/structures"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Reads a certain directory to collect file names
func ReadDir(path string, content []structures.FileData, skipHidden bool, RFlag bool) []structures.FileData {
	var fileList []structures.FileData

	os.Chdir(path)
	saveDirPath := path

	var pathToWorkWith string
	var name string
	UpperPath := GetUpperPath(path)

	for m := 0; m < 2; m++ {
		var DotsFolderData structures.FileData
		if m == 0 {
			pathToWorkWith = path
			name = "."
		} else {
			pathToWorkWith = UpperPath
			name = ".."
		}

		AppendData(&DotsFolderData, pathToWorkWith, saveDirPath)
		DotsFolderData.Name = name
		DotsFolderData.IsHidden = true
		fileList = append(fileList, DotsFolderData)
	}

	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	sorts.SortWordArr(list)

	for _, name := range list {
		if checks.IsHidden(name) && !skipHidden {
			continue
		}

		var dataToAppend structures.FileData
		AppendData(&dataToAppend, name, saveDirPath)

		if dataToAppend.Name == "" {
			return content
		}

		if dataToAppend.IsDirectory && RFlag {
			subFolderPath := path + "/" + dataToAppend.Name
			dataToAppend.SubFolder = ReadDir(subFolderPath, content, skipHidden, RFlag)
		}

		os.Chdir(saveDirPath)
		fileList = append(fileList, dataToAppend)
		file.Close()
	}

	return fileList
}

// Gets data from all files to save it into array
func AppendData(dataToAppend *structures.FileData, name string, saveDirPath string) {
	fileInfo, err := os.Lstat(name)

	if err != nil {
		return
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		dataToAppend.SymLinkPath, _ = os.Readlink(name)
	}

	timeToAppend := fmt.Sprintf("%+03d:%+03d", fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
	timeToAppend = strings.Replace(timeToAppend, "+", "", -1)

	dataToAppend.IsDirectory = fileInfo.IsDir()
	dataToAppend.IsHidden = checks.IsHidden(name)
	perm := permString(fileInfo)
	switch perm {
	case "c":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = "c" + fileInfo.Mode().Perm().String()[1:]
		//Name of file given
		dataToAppend.Name = fileInfo.Name()
	case "d":
		dataToAppend.IsDirectory = true
		dataToAppend.Permission = strings.ToLower(fmt.Sprintf("%v", fileInfo.Mode()))
		//Name of file given

		dataToAppend.Name = fileInfo.Name()

	case "b":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = "b" + fileInfo.Mode().Perm().String()[1:]
		//Name of file given
		dataToAppend.Name = fileInfo.Name()
	case "l":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = "l" + fileInfo.Mode().Perm().String()[1:]
		//Name of file given
	case "-":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = strings.ToLower(fmt.Sprintf("%v", fileInfo.Mode()))
		//Name of file given

		dataToAppend.Name = fileInfo.Name()

	case "p":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = "p" + fileInfo.Mode().Perm().String()[1:]
	case "s":
		dataToAppend.IsDirectory = false
		dataToAppend.Permission = "s" + fileInfo.Mode().Perm().String()[1:]
	}

	dataToAppend.Name = fileInfo.Name()
	//dataToAppend.Size = fileInfo.Size()
	if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		// Cette partie vérifie si les informations spécifiques au système sur le fichier peuvent être converties en syscall.Stat_t
		// fileInfo est une structure contenant des informations sur le fichier
		// syscall.Stat_t est une structure contenant des informations spécifiques au système sur le fichier
		devValue := stat.Rdev
		// stat.Rdev représente le numéro de périphérique (numéros majeur et mineur combinés)
		major := uint64((devValue >> 8) & 0xfff)
		// Les bits 12 à 31 représentent le numéro majeur du périphérique
		minor := uint64(devValue & 0xff)
		// Les bits 0 à 7 représentent le numéro mineur du périphérique
		if devValue == 0 {
			// Si le numéro de périphérique est 0, cela indique un fichier ordinaire (pas un périphérique)
			// Dans ce cas, définissez la taille du fichier sur la taille de bloc du système de fichiers
			dataToAppend.Size = stat.Size
		} else {
			// Si le numéro de périphérique n'est pas 0, cela indique un fichier de périphérique
			// Définissez les numéros majeur et mineur du périphérique
			dataToAppend.Major = major
			dataToAppend.Minor = minor
		}
	}
	dataToAppend.ModificationTime.Day = fileInfo.ModTime().Day()
	dataToAppend.ModificationTime.Month = fileInfo.ModTime()
	dataToAppend.ModificationTime.Time = timeToAppend
	dataToAppend.ModificationTime.FullTime = fileInfo.ModTime()
	dataToAppend.Path = saveDirPath

	if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		UID, err := user.LookupId(fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Uid))
		if err != nil {
			UID = &user.User{Username: fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Uid)}
			return
		}
		dataToAppend.Owner = UID.Username
		group, err := user.LookupGroupId(fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Gid))
		if err != nil {
			group = &user.Group{Gid: fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Gid)}
			return
		}
		dataToAppend.Group = group.Name

		dataToAppend.Hardlinks = int(stat.Nlink)
		dataToAppend.SizeKB = int(stat.Blocks / 2)

	} else {
		panic("permission denied")
	}
}

// Gets data from given files in START directory
func DataFromMainDir(files []string, contentList []structures.FileData, flagsToUse structures.Flags, fs *[]structures.FolderContent) {
	if len(files) != 0 {
		var seekingContent []structures.FileData
		sorts.SortWordArr(files)

		for i := 0; i < len(files); i++ {
			for k := 0; k < len(contentList); k++ {
				if contentList[k].Name == files[i] {
					seekingContent = append(seekingContent, contentList[k])
					break
				}

				if k == len(contentList)-1 {
					fmt.Printf("ls: cannot access '%s': No such file or directory\n", files[i])
				}
			}
		}

		seekingContent = flags.ApplyFlags(flagsToUse, seekingContent)
		if len(seekingContent) != 0 {
			CollectFiles(seekingContent, seekingContent[0].Path, flagsToUse, fs, true)
		}
	} else {
		CollectFiles(contentList, contentList[0].Path, flagsToUse, fs, false)
	}
}

// Gets data from files in given directory
func DataFromDifferentDir(paths []string, flagsToUse structures.Flags, fs *[]structures.FolderContent) {
	for i := 0; i < len(paths); i++ {
		var tempVar []structures.FileData
		tempVar = ReadDir(paths[i], tempVar, flagsToUse.Flag_a, flagsToUse.Flag_R)
		tempVar = flags.ApplyFlags(flagsToUse, tempVar)
		CollectFiles(tempVar, tempVar[0].Path, flagsToUse, fs, false)
	}
}

// Structures all collected folders into list with files data
func CollectFiles(content []structures.FileData, path string, flagsToUse structures.Flags, res *[]structures.FolderContent, certainFiles bool) {
	var dataToAppend structures.FolderContent
	var totalCalculated bool
	longestOwnerName := 0
	longestGroupName := 0
	maxLinksLen := 0
	maxSizeLen := 0
	maxPerLen := 0
	for _, v := range content {
		longestOwnerName = max(longestOwnerName, len(v.Owner))
		longestGroupName = max(longestGroupName, len(v.Group))
		maxLinksLen = max(maxLinksLen, len(strconv.Itoa(v.Hardlinks)))
		maxSizeLen = max(maxSizeLen, len(strconv.Itoa(int(v.Size))))
		maxPerLen = max(maxPerLen, len(v.Permission))
	}

	if certainFiles {
		dataToAppend.Path = "null"
	} else {
		dataToAppend.Path = path + ":"
	}

	for i := 0; i < len(content); i++ {
		if !flagsToUse.Flag_a && content[i].IsHidden {
			if content[i].Name == "." {
				dataToAppend.MainData = content[i]
			}
			continue
		}

		if flagsToUse.Flag_l {
			if !totalCalculated {
				if certainFiles {
					dataToAppend.Total = -1
				} else {
					dataToAppend.Total = calculations.CalculateBlocks(content)
				}
				totalCalculated = true
			}

			var currentTime time.Time = time.Now()
			var yearOrTime string
			diff := currentTime.Sub(content[i].ModificationTime.FullTime)

			if diff.Hours() >= 4380 {
				yearOrTime = strconv.Itoa(content[i].ModificationTime.FullTime.Year())
			} else {
				yearOrTime = content[i].ModificationTime.Time
			}
			var fileName string
			//timerUnix := content[i]. + content[i].Name + "\n"
			format := fmt.Sprintf("%%-%ds %%-%dd %%-%ds %%-%ds  %%-%dd  %%-%ds %%s\n",
				maxPerLen, maxLinksLen, longestOwnerName, longestGroupName, maxSizeLen, 10)
			if content[i].Major == 0 && content[i].Minor == 0 {
				// Si les numéros majeur et mineur sont nuls, affichez la taille du fichier
				fileName = fmt.Sprintf(format,
					content[i].Permission,
					content[i].Hardlinks,
					content[i].Owner,
					content[i].Group,
					content[i].Size,
					content[i].ModificationTime.Month.UTC().Format("Jan")+" "+strconv.Itoa(content[i].ModificationTime.Day)+" "+yearOrTime+" ",
					content[i].Name)
			} else {
				// Sinon, affichez les numéros majeur et mineur du périphérique
				format = fmt.Sprintf("%%-%ds %%-%dd %%-%ds %%-%ds  %%-%dd, %%-%dd  %%-%ds %%s\n",
					maxPerLen, maxLinksLen, longestOwnerName, longestGroupName, 2, 3, 12)
				fileName = fmt.Sprintf(format,
					content[i].Permission,
					content[i].Hardlinks,
					content[i].Owner,
					content[i].Group,
					content[i].Major,
					content[i].Minor,
					content[i].ModificationTime.Month.UTC().Format("Jan")+" "+strconv.Itoa(content[i].ModificationTime.Day)+" "+yearOrTime+" ",
					content[i].Name)
			}

			if content[i].SymLinkPath != "" {
				fileName = fileName[:len(fileName)-1]
				fileName = fileName + " -> " + content[i].SymLinkPath + "\n"
			}

			dataToAppend.FileNames = append(dataToAppend.FileNames, fileName)

		} else {
			dataToAppend.FileNames = append(dataToAppend.FileNames, content[i].Name+" ")
		}
	}

	*res = append(*res, dataToAppend)

	if flagsToUse.Flag_R {
		for i := 0; i < len(content); i++ {
			if content[i].IsDirectory && content[i].Name != "." && content[i].Name != ".." {
				CollectFiles(content[i].SubFolder, path+"/"+content[i].Name, flagsToUse, res, false)
			}
		}
	}
}

// Prints collected data
func PrintData(fs []structures.FolderContent, flagsToUse structures.Flags) {
	for i := 0; i < len(fs); i++ {
		if len(fs) > 1 && !flagsToUse.Flag_R && fs[i].Path != "null" {
			fmt.Println(fs[i].Path)
		}

		if len(fs[i].FileNames) == 0 {
			if flagsToUse.Flag_R && fs[i].Path != "null" {
				fmt.Println(fs[i].Path)
			}
			if flagsToUse.Flag_l && fs[i].Total != -1 {
				fmt.Println("total:", fs[i].Total)
			}
		}

		for k := 0; k < len(fs[i].FileNames); k++ {
			if flagsToUse.Flag_R && k == 0 && fs[i].Path != "null" {
				fmt.Println(fs[i].Path)
			}
			if flagsToUse.Flag_l && k == 0 && fs[i].Total != -1 {
				fmt.Println("total:", fs[i].Total)
			}
			fmt.Print(fs[i].FileNames[k])
		}

		if len(fs) == 1 && !flagsToUse.Flag_l && len(fs[i].FileNames) != 0 {
			fmt.Println()
			continue
		}

		if len(fs) > 1 && !flagsToUse.Flag_l {
			fmt.Println()
		}

		if i != len(fs)-1 && len(fs[i].FileNames) != 0 {
			fmt.Println()
		}
	}
}

// Gets parent dataToAppend path of START directory
func GetUpperPath(path string) string {
	pathInRune := []rune(path)
	var dashCounter int = 0

	for k := 0; k < len(pathInRune); k++ {
		if pathInRune[k] == '/' {
			dashCounter++
		}
	}

	if dashCounter == 1 {
		return "/"
	}

	for k := 0; k < len(path); k++ {
		if path[len(path)-1:] == "/" {
			path = path[:len(path)-1]
			break
		} else {
			path = path[:len(path)-1]
		}
	}

	return path
}
func permString(info os.FileInfo) string {
	// info.Mode().String() does not produce the same output as `ls`, so we must build that string manually
	mode := info.Mode()
	// this "type" is not the file extension, but type as far as the OS is concerned
	filetype := "-"
	if mode&os.ModeDir != 0 {
		filetype = "d"
	} else if mode&os.ModeSymlink != 0 {
		filetype = "l"
	} else if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice == 0 {
			filetype = "b" // block device
		} else {
			filetype = "c" // character device
		}
	} else if mode&os.ModeNamedPipe != 0 {
		filetype = "p"
	} else if mode&os.ModeSocket != 0 {
		filetype = "s"
	}
	return filetype
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func IsExecutableScript(info os.FileInfo) bool {
	if info.Mode()&0111 != 0 && info.Mode().IsRegular() {
		return true
	}
	return false
}
