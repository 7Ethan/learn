package main

import (
	_ "fmt"
	"io/ioutil"
	"log"
	"os"
	_ "reflect"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("get path failed")
	}
	//println("---++++++++++++")
	//fmt.Println(reflect.TypeOf(path).Name())
	path = path + "/log"
	os.Chdir("~")
	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|
		os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln("open file failed : ", err)
	}

	Trace = log.New(ioutil.Discard, "Trace : ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning : ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "Error : ", log.Ldate|log.Ltime|log.Lshortfile) //级别最高，需要持久化到文件
}

func main() {
	//path ,err := os.Getwd()
	//if err != nil {
	//	log.Fatalln("get path failed")
	//}
	//println("---++++++++++++")
	////fmt.Println(reflect.TypeOf(path).Name())
	//path = path + "/log"
	//fmt.Println(path)
	Trace.Println("Trace ---------")
	Info.Println("Info --------")
	Warning.Println("Warning _________")
	Error.Println("Error ------------")
}
