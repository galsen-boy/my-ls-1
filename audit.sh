#!/bin/bash
RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"
Continue() {
  echo -e "\nPress any key to continue"
  while [ true ]; do
    read -t 3 -n 1
    if [ $? = 0 ]; then
      clear
      return
    fi
  done
}
clear

echo -e "Creating folders and files for audit\n"
echo -e "Creating folder test1..."
mkdir test1
echo -e "Creating folder test2..."
mkdir test2
echo -e "Creating folder test3..."
mkdir test3
echo -e "Creating file file1 inside test1 folder..."
touch ./test1/file1
echo -e "Creating file file2 inside test1 folder..."
touch ./test1/file2
echo -e "Creating file file3 inside test1 folder..."
touch ./test1/file3
echo -e "Creating file 1.txt inside test2 folder..."
touch ./test2/1.txt
echo -e "Creating file 2.txt inside test2 folder..."
touch ./test2/2.txt
echo -e "Creating file 3.txt inside test2 folder..."
touch ./test2/3.txt
echo -e "Creating folder folder1 inside test1 folder..."
mkdir ./test1/folder1
echo -e "Creating folder folder2 inside test1 folder..."
mkdir ./test1/folder2
echo -e "Creating folder folder1 inside test2 folder..."
mkdir ./test2/folder1
echo -e "Creating file 1.txt inside test1/folder1 folder..."
touch ./test1/folder1/1.txt
echo -e "Creating file 2.txt inside test2/folder1 folder..."
touch ./test2/folder1/2.txt
echo -e "Creating hidden folder..."
mkdir .hiddenfolder
echo -e "Creating hidden file..."
touch .hiddenfile
echo -e "Creating symlink for file..."
ln -s ./test2/3.txt ./symboliclinktofile
echo -e "Creating symlink for folder..."
ln -s ./test2/ ./symboliclinktofolder
Continue

echo -e "Run both my-ls-1 and the system command ls with no arguments.\n"
echo -e "${RED}System command ls${ENDCOLOR}"
ls
echo -e "\n${GREEN}my-ls-1${ENDCOLOR}"
go run .
Continue

echo -e "Run both my-ls-1 and the system command ls with the arguments: '<file name>'.\n"
echo -e "${RED}System command ls main.go${ENDCOLOR}"
ls main.go
echo -e "\n${GREEN}my-ls-1 main.go${ENDCOLOR}"
go run . "main.go"
Continue

echo -e "Run both my-ls-1 and the system command ls with the arguments: '<directory name>'.\n"
echo -e "${RED}System command ls sorts${ENDCOLOR}"
ls sorts
echo -e "\n${GREEN}my-ls-1 sorts${ENDCOLOR}"
go run . "sorts"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l'\n"
echo -e "${RED}System command ls -l${ENDCOLOR}"
ls -l
echo -e "\n${GREEN}my-ls-1 -l${ENDCOLOR}"
go run . "-l"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l <file name>'\n"
echo -e "${RED}System command ls -l main.go${ENDCOLOR}"
ls -l main.go
echo -e "\n${GREEN}my-ls-1 -l main.go${ENDCOLOR}"
go run . "-l" "main.go"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l <directory name>'\n"
echo -e "${RED}System command ls -l test1${ENDCOLOR}"
ls -l test1
echo -e "\n${GREEN}my-ls-1 -l test1${ENDCOLOR}"
go run . "-l" "test1"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l /usr/bin'\n"
echo -e "${RED}System command ls -l /usr/bin${ENDCOLOR}"
Continue
ls -l /usr/bin
Continue
echo -e "\n${GREEN}my-ls-1 -l /usr/bin${ENDCOLOR}"
Continue
go run . "-l" "/usr/bin"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-R', in a directory with folders in it.\n"
echo -e "${RED}System command ls -R test1${ENDCOLOR}"
ls -R test1
echo -e "\n${GREEN}my-ls-1 -R test1${ENDCOLOR}"
go run . "-R" "test1"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-a'\n"
echo -e "${RED}System command ls -a${ENDCOLOR}"
ls -a
echo -e "\n${GREEN}my-ls-1 -a${ENDCOLOR}"
go run . "-a"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-r'\n"
echo -e "${RED}System command ls -r${ENDCOLOR}"
ls -r
echo -e "\n${GREEN}my-ls-1 -r${ENDCOLOR}"
go run . "-r"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-t'\n"
echo -e "${RED}System command ls -t${ENDCOLOR}"
ls -t
echo -e "\n${GREEN}my-ls-1 -t${ENDCOLOR}"
go run . "-t"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-la'\n"
echo -e "${RED}System command ls -la${ENDCOLOR}"
ls -la
echo -e "\n${GREEN}my-ls-1 -la${ENDCOLOR}"
go run . "-la"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l -t <directory name>'\n"
echo -e "${RED}System command ls -l -t test1${ENDCOLOR}"
ls -l -t test1
echo -e "\n${GREEN}my-ls-1 -l -t test1${ENDCOLOR}"
go run . "-l" "-t" "test1"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-lRr <directory name>'\n"
echo -e "${RED}System command ls -lRr test1${ENDCOLOR}"
ls -lRr test1
echo -e "\n${GREEN}my-ls-1 -lRr test1${ENDCOLOR}"
go run . "-lRr" "test1"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-l <directory name> -a <file name>'\n"
echo -e "${RED}System command ls -l test1 -a main.go${ENDCOLOR}"
ls -l test1 -a main.go
echo -e "\n${GREEN}my-ls-1 -l test1 -a main.go${ENDCOLOR}"
go run . "-l" "test1" "-a" "main.go"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-lR <directory name>///<sub directory name>/// <directory name>/<sub directory name>/'\n"
echo -e "${RED}System command ls -lR test1///folder1/// test2/folder1/${ENDCOLOR}"
ls -lR test1///folder1/// test2/folder1/
echo -e "\n${GREEN}my-ls-1 -lR test1///folder1/// test2/folder1/${ENDCOLOR}"
go run . "-lR" "test1///folder1///" "test2/folder1/"
Continue

echo -e "Run both my-ls-1 and the system command ls with the flag: '-la /dev'\n"
echo -e "${RED}System command ls -la /dev${ENDCOLOR}"
Continue
ls -la /dev
Continue
echo -e "\n${GREEN}my-ls-1 -la /dev${ENDCOLOR}"
Continue
go run . "-la" "/dev"
Continue

echo -e "Create directory with - name and run both my-ls-1 and the system command ls with the arguments: '-'"
echo -e "Creating - folder..."
mkdir -
Continue
echo -e "${RED}System command ls -${ENDCOLOR}"
ls -
echo -e "\n${GREEN}my-ls-1 -${ENDCOLOR}"
go run . "-"
Continue

echo -e "Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: '-l <symlink file>/'"
echo -e "${RED}System command ls -l symboliclinktofile/${ENDCOLOR}"
ls -l symboliclinktofile/
echo -e "\n${GREEN}my-ls-1 -l symboliclinktofile/${ENDCOLOR}"
go run . "-l" "symboliclinktofile/"
Continue

echo -e "Create file and link for this file and run both my-ls-1 and the system command ls with the arguments: '-l <symlink file>'"
echo -e "${RED}System command ls -l symboliclinktofile${ENDCOLOR}"
ls -l symboliclinktofile
echo -e "\n${GREEN}my-ls-1 -l symboliclinktofile${ENDCOLOR}"
go run . "-l" "symboliclinktofile"
Continue

echo -e "Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: '-l <symlink dir>/'"
echo -e "${RED}System command ls -l symboliclinktofolder/${ENDCOLOR}"
ls -l symboliclinktofolder/
echo -e "\n${GREEN}my-ls-1 -l symboliclinktofolder/${ENDCOLOR}"
go run . "-l" "symboliclinktofolder/"
Continue

echo -e "Create directory that contains files and link for this directory and run both my-ls-1 and the system command ls with the arguments: '-l <symlink dir>'"
echo -e "${RED}System command ls -l symboliclinktofolder${ENDCOLOR}"
ls -l symboliclinktofolder
echo -e "\n${GREEN}my-ls-1 -l symboliclinktofolder${ENDCOLOR}"
go run . "-l" "symboliclinktofolder"
Continue

echo -e "Try running the program with '-R ~' and with the command time before the program name (ex: 'time ./my-ls-1 -R ~')."
time go run . "-R" ~
Continue