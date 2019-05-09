package main

import (
	"sort"
	"fmt"
	"io"
	"os"
)

func dirTree(writer io.Writer, filePath string, printFiles bool) error {
	graphics := [3]string {"├───","└───","│"}
	return dirTreeInner(writer, filePath, printFiles, 0, graphics, 0)
}

func dirTreeInner(writer io.Writer, filePath string, printFiles bool, level int, graphics [3]string, leftPrint int) error {
	f, err := os.Open(filePath)

	if err != nil {
		return err
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}

	if !printFiles {
		var listInner []os.FileInfo

		for idx, file := range list {
			if file.IsDir() {
				listInner = append(listInner, list[idx])
			}
		}

		list = listInner
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })

	for idx, file := range list {

		for idxInner := 0; idxInner < level; idxInner++ {
			if idxInner < level - leftPrint {
				fmt.Fprint(writer, graphics[2]+"\t")
			} else {
				fmt.Fprint(writer, "\t")
			}
		}

		if idx == len(list) - 1 {
			fmt.Fprint(writer, graphics[1])
			leftPrint++
		} else {
			fmt.Fprint(writer, graphics[0])
		}
		fmt.Fprint(writer, file.Name())

		if file.IsDir() {
			fmt.Fprint(writer, "\n")
			dirTreeInner(writer, filePath + "/" + file.Name(), printFiles, level + 1, graphics, leftPrint)
		} else {
			if size := file.Size(); size != 0 {
				fmt.Fprint(writer, " (", size, "b)\n")
			} else {
				fmt.Fprint(writer, " (empty)\n")
			}
		}
	}
	return nil
}