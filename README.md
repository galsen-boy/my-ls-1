## Author

* Mouhamadou Fadilou Diop (https://learn.zone01dakar.sn/git/mouhamadoufadiop/)
* Abdou Khadre Wade (https://learn.zone01dakar.sn/git/serwade)
* Alimoudine Idrissou (https://learn.zone01dakar.sn/git/ialimoud)
<h1 align="center">MY-LS</h1>

## About The Project
My-ls is a project, that consists on creating your own ls command.

## Installation
```
git clone https://learn.zone01dakar.sn/git/serwade/my-ls-1
cd my-ls
```

## Usage
```
go run . [OPTIONS] [FILE|DIR]
```

## Examples
```
go run . -l == ls -l
go run . -r == ls -r
go run . folder/ == ls folder/
go run . -lraRt == ls -lraRt
```

## Available options

* <code>-l</code> - list with long format
* <code>-r</code> - list in reverse order
* <code>-a</code> - list all files including hidden file starting with '.'
* <code>-R</code> - list recursively directory tree
* <code>-t</code> - sort by time & date

## Additional information

Only standard go packages were in use.

