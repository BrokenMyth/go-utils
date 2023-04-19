package strings

import (
	"strings"
	"unicode"
)

// UnderlineToHump 下划线写法转为驼峰写法
func UnderlineToHump(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// CapitalizeFirst 首字母大写
func CapitalizeFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// HumpToUnderline 驼峰式写法转为下划线写法
func HumpToUnderline(name string) string {
	buffer := ""
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer += "_"
			}
			buffer += string(unicode.ToLower(r))
		} else {
			buffer += string(r)
		}
	}
	return buffer
}

// PojoFirst 首字母小写
func PojoFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
