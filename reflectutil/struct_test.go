package reflectutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStructTagsValue(t *testing.T) {
	s := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Id: "1xx", Name: "xiaoming", Age: 18}
	value := GetStructTagsValue(s, "json")

	assert.Equal(t, []string{"id", "name", "age"}, value)
}

func TestGetStructFieldValue(t *testing.T) {
	s := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Id: "1xx", Name: "xiaoming", Age: 18}
	value := GetStructFieldValue(s)

	fmt.Println(value)
}

func TestGetStructSliceFieldValue(t *testing.T) {
	s := []*struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{Id: "1xx", Name: "xiaoming", Age: 18},
		{Id: "2xx", Name: "xiaoxxxx", Age: 19},
	}
	value, _ := GetStructSliceFieldValue(s)

	fmt.Println(value)
}

func TestGetStructFieldName(t *testing.T) {
	s := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Id: "1xx", Name: "xiaoming", Age: 18}
	value := GetStructFieldName(&s)

	fmt.Println(value)
}
