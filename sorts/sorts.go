package sorts

// Sorts name array
func SortWordArr(table []string) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				tempVar := table[j]
				table[j] = table[j+1]
				table[j+1] = tempVar
			}
		}
	}
}
