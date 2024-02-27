package r_flag

import "my-ls/structures"

// Returns reverted file array
func ReverseList(table []structures.FileData) []structures.FileData {
	var result []structures.FileData

	for i := len(table) - 1; 0 <= i; i-- {
		result = append(result, table[i])
	}

	for i := 0; i < len(table); i++ {
		if result[i].IsDirectory {
			result[i].SubFolder = ReverseList(result[i].SubFolder)
		}
	}

	return result
}
