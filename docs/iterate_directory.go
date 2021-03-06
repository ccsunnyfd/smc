package main
 
import (
	"fmt"
	"io/ioutil"
)
 
func GetAllFile(pathname string, count int) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}
	for _, fi := range rd {
		for i := 0; i < count; i++ {
			fmt.Print("  ")
		}
		if fi.IsDir() {
			fmt.Println("* " + fi.Name())
			count++
			fullDir := pathname + "/" + fi.Name()
			GetAllFile(fullDir, count)
			count--
			if err != nil {
				fmt.Println("read dir fail:", err)
			}
		} else {
			fileName := fi.Name()
			if fileName[len(fileName)-3:] == ".md" {
				fullName := pathname + "/" + fileName
				fmt.Println("* [" + fileName[:len(fileName)-3] + "](" + fullName[3:len(fullName)-3] + ")")
			}
		}
	}
}
 
func main() {
	//遍历打印所有的文件名
	GetAllFile("./", 0)
 
	fmt.Printf("done")
}