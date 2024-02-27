package calculations

import "my-ls/structures"

// Calculate disk usage of listed files
func CalculateBlocks(filesData []structures.FileData) int {
	var sum int
	for i, _ := range filesData {
		sum = sum + filesData[i].SizeKB
	}
	return sum
}
