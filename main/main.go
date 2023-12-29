package main

import "github.com/BrokenMyth/go-utils/fileutil"

func main() {
	//err := fileutil.FileGenerateFromString("./main/src/pkg", "ax.txt", "aaa")

	err := fileutil.FileAppendString("./main/src/pkg", "a.txt", "aaa")
	if err != nil {
		return
	}
}
