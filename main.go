package main

import (
	"fmt"
	"os"
	"strconv"
)

var emptyDirs []string
var emptyFiles []string
var errors []string
var depthMap = make(map[int]bool)

func generateDashes(dashes int) string {
	temp := ""
	for i := 0; i < dashes; i++ {
		temp += "━"
	}
	return temp
}

func generateSpaces(spaces int) string {
	temp := ""
	for i := 0; i < spaces; i++ {
		if i%2 == 0 && depthMap[i] {
			temp += "┃"
		} else {
			temp += " "
		}
	}
	return temp
}

func isDirEmpty(path string) bool {
	entries, err := os.ReadDir(path)
	return err != nil || len(entries) == 0
}

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func listEntries(path string, initialDepth int, maxDepth int) float64 {
	if maxDepth > -1 {
		if initialDepth/2 >= maxDepth {
			return 0
		}
	}

	path = path + "/"

	entries, err := os.ReadDir(path)

	var size float64 = 0
	var s float64

	depthMap[initialDepth] = false
	if len(entries) > 1 {
		depthMap[initialDepth] = true
	}

	if err != nil {
		errors = append(errors, err.Error())
	}

	if initialDepth == 0 {
		fmt.Println(".")
	}

	for i, e := range entries {
		if e.IsDir() {
			if i == len(entries)-1 {
				depthMap[initialDepth] = false
				if isDirEmpty(path+e.Name()) || initialDepth/2 == maxDepth-1 {
					fmt.Printf("%s┗%s /%s\n", generateSpaces(initialDepth), generateDashes(2), e.Name())
				} else {
					fmt.Printf("%s┗%s┓ /%s\n", generateSpaces(initialDepth), generateDashes(1), e.Name())
				}
			} else {
				if isDirEmpty(path+e.Name()) || initialDepth/2 == maxDepth-1 {
					fmt.Printf("%s┣%s /%s\n", generateSpaces(initialDepth), generateDashes(2), e.Name())
				} else {
					fmt.Printf("%s┣%s┓ /%s\n", generateSpaces(initialDepth), generateDashes(1), e.Name())
				}
			}

			size += listEntries(path+e.Name(), initialDepth+2, maxDepth)
		} else {
			info, _ := e.Info()
			s = float64(info.Size())
			size += s

			if size == 0 {
				emptyFiles = append(emptyFiles, path+info.Name())
			}

			if i == len(entries)-1 {
				fmt.Printf("%s┗%s %s [%.1f KB]\n", generateSpaces(initialDepth), generateDashes(2), e.Name(), s/1024)
			} else {
				fmt.Printf("%s┣%s %s [%.1f KB]\n", generateSpaces(initialDepth), generateDashes(2), e.Name(), s/1024)
			}
		}
	}

	if size == 0 {
		emptyDirs = append(emptyDirs, path)
	}

	return size
}

func main() {
	fmt.Println(os.Args)
	if !isAdmin() {
		fmt.Println("Please run the program as an administrator for better results")
	}

	path, _ := os.Getwd()
	maxDepth := -1

	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if len(os.Args) > 2 {
		maxDepth, _ = strconv.Atoi(os.Args[2])
	}

	totalSize := listEntries(path, 0, maxDepth)

	if maxDepth == -1 {
		fmt.Printf("\nsize of \"%s\" is %.2f MB\n", path, totalSize/1024000)
	}
	fmt.Printf("\nfound %d empty directories\n", len(emptyDirs))
	fmt.Printf("\nfound %d empty files\n", len(emptyFiles))
	fmt.Printf("\n%d errors encountered\n", len(errors))

	fmt.Print("\npress enter to exit...")
	fmt.Scanln()
}
