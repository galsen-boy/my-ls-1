package structures

import "time"

type FileData struct {
	IsDirectory      bool
	IsHidden         bool
	Permission       string
	Name             string
	Hardlinks        int
	Owner            string
	Group            string
	Size             int64
	SizeKB           int
	Major            uint64
	Minor            uint64
	SymLinkPath      string
	ModificationTime Date
	SubFolder        []FileData // If it's folder, here we save all children files data
	Path             string
}

type Date struct {
	FullTime time.Time
	Month    time.Time
	Day      int
	Time     string
	Year     int
}

type Flags struct {
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
}

type FolderContent struct {
	Path      string
	Total     int
	FileNames []string
	MainData  FileData
}

var STARTDIR string // starting folder
