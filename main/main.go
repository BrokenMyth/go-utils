package main

import "github.com/BrokenMyth/go-utils/fileutil"

func main() {
	//err := fileutil.GenerateFileFromString("./main/src/pkg", "ax.txt", "aaa")

	err := fileutil.FileAppend("./main/src/pkg", "a.txt", "aaa")
	if err != nil {
		return
	}
}
