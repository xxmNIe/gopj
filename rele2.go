package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	relsease("/home/xxm/chronos.txt")
}


func relsease(path string){
	or, err1 := os.OpenFile(path, os.O_RDONLY, 0777)
	if err1 != nil {
		log.Fatal("openfile err:", err1)
	}
	reader := bufio.NewReader(or)

	var file *os.File = nil
	var err error =nil
	for err == nil {
		bytes,_, err := reader.ReadLine()
		readString := string(bytes)
		if err != nil {
			log.Fatal(err)
			break
		}
		if strings.Contains(readString,"nameend") {
			if file !=nil {
				file.Close()
			}
			filename := readString[:strings.Index(readString,"---")]
			filename = strings.Replace(filename,"\\code\\","\\release\\",1)
			file,err = createwithpath(filename)
			if err !=nil {
				log.Fatal("create file with dir err :",err)
				return
			}
			continue
		}
		file.WriteString(readString+"\r\n")

	}

}
func createwithpath(fullname string)(*os.File, error){
	var err error
	path := fullname[:strings.LastIndex(fullname,"\\")]
	if  err = os.MkdirAll(path,0777) ;err != nil {
		log.Fatal("create dir err :",err)
		return  nil,err
	}

	create, err := os.Create(fullname)
	return create,err

}