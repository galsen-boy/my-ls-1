package t_flag

import (
	"my-ls/structures"
)

// Sorts files by time
func SortByTime(table []structures.FileData) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j].ModificationTime.Month.Month().String() > table[j+1].ModificationTime.Month.Month().String() {
				tempVar := table[j]
				table[j] = table[j+1]
				table[j+1] = tempVar
			} else if table[j].ModificationTime.Month.Month().String() == table[j+1].ModificationTime.Month.Month().String() {
				if table[j].ModificationTime.Day < table[j+1].ModificationTime.Day {
					tempVar := table[j]
					table[j] = table[j+1]
					table[j+1] = tempVar
				} else if table[j].ModificationTime.Day == table[j+1].ModificationTime.Day {
					if table[j].ModificationTime.Month.Hour() < table[j+1].ModificationTime.Month.Hour() {
						tempVar := table[j]
						table[j] = table[j+1]
						table[j+1] = tempVar
					} else if table[j].ModificationTime.Month.Hour() == table[j+1].ModificationTime.Month.Hour() {
						if table[j].ModificationTime.Month.Minute() < table[j+1].ModificationTime.Month.Minute() {
							tempVar := table[j]
							table[j] = table[j+1]
							table[j+1] = tempVar
						} else if table[j].ModificationTime.Month.Minute() == table[j+1].ModificationTime.Month.Minute() {
							if table[j].ModificationTime.Month.Second() < table[j+1].ModificationTime.Month.Second() {
								tempVar := table[j]
								table[j] = table[j+1]
								table[j+1] = tempVar
							} else if table[j].ModificationTime.Month.Second() == table[j+1].ModificationTime.Month.Second() {
								if table[j].ModificationTime.Month.Nanosecond() < table[j+1].ModificationTime.Month.Nanosecond() {
									tempVar := table[j]
									table[j] = table[j+1]
									table[j+1] = tempVar
								}
							}
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(table); i++ {
		if table[i].IsDirectory {
			SortByTime(table[i].SubFolder)
		}
	}
}
