package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	rootpath := "/etc"

	root := FileNode{"projects", rootpath, []*FileNode{}}
	fileInfo, _ := os.Lstat(rootpath)

	walk(rootpath, fileInfo, &root)

	data, _ := json.Marshal(root)

	fmt.Printf("%s", data)
}

type FileNode struct {
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	FileNodes []*FileNode `json:"children"`
}

func walk(path string, info os.FileInfo, node *FileNode) {
	// 列出当前目录下的所有目录、文件
	files := listFiles(path)

	// 遍历这些文件
	for _, filename := range files {
		// 拼接全路径
		fpath := filepath.Join(path, filename)

		// 构造文件结构
		fio, _ := os.Lstat(fpath)

		// 将当前文件作为子节点添加到目录下
		child := FileNode{filename, fpath, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)

		// 如果遍历的当前文件是个目录，则进入该目录进行递归
		if fio.IsDir() {
			walk(fpath, fio, &child)
		}
	}

	return
}

func listFiles(dirname string) []string {
	f, _ := os.Open(dirname)

	names, _ := f.Readdirnames(-1)
	f.Close()

	sort.Strings(names)

	return names
}
