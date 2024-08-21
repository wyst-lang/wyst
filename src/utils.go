package main

import (
	"fmt"
	"os"
	"path"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

type Module struct {
	Code  string
	Name  string
	Path  string
	Out   string
	Inner []Module
}

func writeFile(file_name, file_content string) {
	os.MkdirAll(path.Dir(file_name), os.ModePerm)
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Printf("Error creating file (%s): %v\n", file_name, err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(file_content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func readFile(file_name string) string {
	out, _ := os.ReadFile(file_name)
	return string(out)
}

func NewRoot(file_path string) Module {
	root := Module{}

	root.Path = file_path
	root.Out = "out/main.go"
	root.Name = "root"
	root.Code = root.Transpile(readFile(file_path))

	return root
}

func (root *Module) ImportModule(modname string) string {
	module := Module{Name: modname}
	module.Path = path.Join(path.Dir(root.Path), module.Name+".wyst")
	if FileExists(module.Path) {
		module.Code = "package " + fmt.Sprintf("_%#x", modname) + "\n" + module.Transpile(readFile(module.Path))
	} else {
		module.Path = path.Join(path.Dir(root.Path), module.Name, "mod.wyst")
		if FileExists(module.Path) {
			module.Code = "package " + fmt.Sprintf("_%#x", modname) + "\n" + module.Transpile(readFile(module.Path))
		}
	}
	module.Out = fmt.Sprintf("%s/I%#x/%s.go", path.Dir(root.Out), module.Name, module.Name)
	root.Inner = append(root.Inner, module)
	return fmt.Sprintf("%s/I%#x", path.Dir(root.Out), module.Name)
}

func WriteModules(root Module) {
	for _, module := range root.Inner {
		writeFile(module.Out, module.Code)
		WriteModules(module)
	}
}

func WriteRoot(root Module) {
	writeFile("out/main.go", "package main\n"+root.Code)
	writeFile("out/go.mod", "module out\n")
	WriteModules(root)
}
