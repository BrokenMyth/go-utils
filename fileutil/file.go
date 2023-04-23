package fileutil

import (
	"fmt"
	"os"
	"strings"
)

// GenerateFileFromString 通过字符串生成文件
func GenerateFileFromString(fileDir, fileName, content string) error {
	if fileName == "" {
		return fmt.Errorf("fileName cannot be empty")
	}
	if fileDir == "" {
		return fmt.Errorf("fileDir cannot be empty")
	}
	b := []byte(content)
	filePath := fileDir + "/" + fileName
	_, err := os.Stat(fileDir)
	if err != nil {
		if os.IsNotExist(err) {
			//尝试逐级创建
			e := GraduallyCreateDir(fileDir)
			if e != nil {
				return err
			}
		} else {
			return err
		}
	}
	err = os.WriteFile(filePath, b, os.ModePerm)
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
