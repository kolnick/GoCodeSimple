package io

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	path1 := "../../res/微信截图_20230615134853.png"
	path2 := "E:\\workspace\\GoCodeSimple\\src\\res\\微信截图_20230615134853.png"
	fmt.Println(filepath.IsAbs(path1))
	// 获取相对路径
	rel, _ := filepath.Rel("E:\\workspace", path2)
	fmt.Println(rel)
	// 获取绝对路径
	fmt.Println(filepath.Abs(path2))

	fmt.Println(path.Join(path1, ".."))
	fmt.Println(path.Join(path1, "."))
	//fmt.Println(path.Join("d:/", path))
}
