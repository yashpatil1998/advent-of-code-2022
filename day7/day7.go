package day7

import (
	"strconv"
	"strings"
)

func Day7Logic() (int64, int64) {
	input := strings.Split(input7, "\n")
	root := buildFileSystem(input)

	unusedDiskSpace := part2TotalDiskSpace - root.Size
	spaceToFreeUp := part2SpaceNeeded - unusedDiskSpace

	part1, part2 := day7Logic(&root, spaceToFreeUp, root.Size)

	return part1, part2
}

type directory struct {
	Name        string
	Directories []*directory
	Files       []file
	Parent      *directory
	Size        int64
}

type file struct {
	Name string
	Size int64
}

const part1DirMaxLimit = 100000
const part2TotalDiskSpace = 70000000
const part2SpaceNeeded = 30000000

func day7Logic(dir *directory, spaceNeeded int64, spaceToDelete int64) (int64, int64) {

	var part1 int64
	var part2 = spaceToDelete

	if dir.Size <= part1DirMaxLimit {
		part1 += dir.Size
	}

	if dir.Size >= spaceNeeded && dir.Size < part2 {
		part2 = dir.Size
	}

	for _, d := range dir.Directories {
		temp1, temp2 := day7Logic(d, spaceNeeded, part2)

		part1 += temp1
		part2 = temp2
	}

	return part1, part2
}

func buildFileSystem(commands []string) directory {
	var root directory
	var currentDir = &root

	for _, command := range commands {
		cmd := strings.Split(command, " ")
		if strings.HasPrefix(command, "$") {
			// if here, means input is a command
			switch cmd[1] {
			case "cd":
				if cmd[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, dir := range currentDir.Directories {
						if dir.Name == cmd[2] {
							currentDir = dir
						}
					}
				}
			case "ls":
			}
		} else {
			// if here, means input is a file or directory
			switch cmd[0] {
			case "dir":
				dir := directory{Name: cmd[1], Parent: currentDir}
				currentDir.Directories = append(currentDir.Directories, &dir)
			default:
				size, _ := strconv.ParseInt(cmd[0], 10, 64)
				newFile := file{Name: cmd[1], Size: size}

				currentDir.Files = append(currentDir.Files, newFile)
				increaseDirSize(size, currentDir)
			}
		}
	}
	return root
}

func increaseDirSize(size int64, dir *directory) {
	dir.Size += size
	if dir.Parent != nil {
		increaseDirSize(size, dir.Parent)
	}
}
