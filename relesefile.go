package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	basedir = "/home/xxm/Gout/"
	pre = "C:\\code\\Gout\\"
)

func Realise(path string) {

	if path == "" {
		log.Fatal("path is empty")
		return
	}
	f, err1 := os.OpenFile(path, os.O_RDONLY, 0777)
	if err1 != nil {
		log.Fatalf(" read file%s err: %T", path, err1)
	}
	reader := bufio.NewReader(f)


	var newfile *os.File = nil
	var err error =nil
	for err == nil {
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
		readString, _,err := reader.ReadLine()
		if err != nil {
			log.Fatal("err :",err)
			break
		}
		str := string(readString)
		str = strings.Replace(str,"\\","/",-1)
		if strings.Contains(str,"nameend") {
			if newfile != nil {
				newfile.Close()
			}
			filename := str[:strings.LastIndex(str,"---")]
			short_name := filename[len(pre):]
			fullname  := basedir+short_name
			//fmt.Println(fullname)
			newfile,err = CreateFilewithDir(fullname)
			if err !=nil {
				log.Fatal(err)
				return
			}
			continue
		}
		newfile.WriteString(str+"\n")
	}

}



func CreateFilewithDir(path string) (*os.File, error) {
	path = strings.Replace(path,"\\","/",-1)
	fmt.Println(path)
	dirs := path[:strings.LastIndex(path, "/")]
	err := os.MkdirAll(dirs, os.ModePerm)
	if err != nil {
		log.Fatal("mkdir ", dirs, " err :", err)
		return nil, err
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil

}

func main() {
	Realise("/home/xxm/Gout.txt")
	//fmt.Println(makefilename("\\code\\servant-netease\\.git\\config"))
}
