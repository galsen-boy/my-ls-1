package main

import (
	//"fmt"
	"my-ls/data_interaction"
	"my-ls/flags"
	"my-ls/structures"
	"os"
)

func main() {
	// if len(os.Args) < 2 || os.Args[1] != "my-ls" {
	// 	fmt.Print("Usage: go run . my-ls [FLAGS]\n\nEX: go run . my-ls -l\n")
	// 	return
	// }
	var contentList []structures.FileData
	var foldersStucture []structures.FolderContent

	structures.STARTDIR, _ = os.Getwd()

	flagsToUse, paths, files, folders, wasPath := flags.CollectAllAgruments(os.Args)

	contentList = data_interaction.ReadDir(structures.STARTDIR, contentList, flagsToUse.Flag_a, flagsToUse.Flag_R)
	contentList = flags.ApplyFlags(flagsToUse, contentList)

	if (len(paths) == 0 && len(folders) == 0 && !wasPath) || len(files) != 0 {
		data_interaction.DataFromMainDir(files, contentList, flagsToUse, &foldersStucture)
	}

	data_interaction.DataFromDifferentDir(paths, flagsToUse, &foldersStucture)
	data_interaction.DataFromDifferentDir(folders, flagsToUse, &foldersStucture)
	data_interaction.PrintData(foldersStucture, flagsToUse)
}
