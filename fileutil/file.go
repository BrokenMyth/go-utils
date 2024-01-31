package fileutil

import (
	"fmt"
	"os"
	"strings"
)

// FileGenerateFromString 通过字符串生成文件
func FileGenerateFromString(fileDir, fileName, content string) error {
	if fileDir == "" {
		fileDir = "."
	}
	//如果指定路径不存在则开始创建
	err := createNotExistDir(fileDir)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileDir+"/"+fileName, []byte(content), os.ModePerm)
	if err != nil {
		fmt.Println("Can't create txt.", err)
		return err
	}
	return nil
}

// GraduallyCreateDir 相对路径逐级创建目录
func GraduallyCreateDir(dir string) error {
	dirs := strings.Split(dir, "/")
	if len(dirs) == 0 || len(dirs) < 2 {
		return fmt.Errorf("dir cannot be empty or cannot be split")
	}
	//说明是相对路径
	if dirs[0] == "." {
		createDir := "."
		for i := 1; i < len(dirs); i++ {
			createDir += "/" + dirs[i]
			_, err := os.Stat(createDir)
			//如果 Err 说明不存在
			if err != nil && os.IsNotExist(err) {
				err = os.Mkdir(createDir, os.ModePerm)
				//如果创建失败，说明目录已经存在
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// createNotExistDir 创建不存在的路径，只支持相对路径创建
func createNotExistDir(fileDir string) error {
	_, err := os.Stat(fileDir)
	if err != nil && os.IsNotExist(err) {
		//尝试逐级创建
		e := GraduallyCreateDir(fileDir)
		if e != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

// FileAppendString 向文件追加一行 str
func FileAppendString(fileDir, fileName, content string) error {
	//如果指定路径不存在则开始创建
	err := createNotExistDir(fileDir)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(content + "\n"))
	if err != nil {
		return err
	}
	return nil
}
