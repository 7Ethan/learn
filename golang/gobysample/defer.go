package main

import (
	"fmt"
	"os"
)

func createfile(p string) (*os.File,error){
	fmt.Println("Creating ..")
	f,err := os.Create(p)
	if err != nil{
		return nil,err
	}
	return f,nil
}

func writefile(f *os.File) {
	fmt.Println("writing files.")
	fmt.Fprintln(f,"data")
}

func closefile(f *os.File) {
	fmt.Println("Closing file..")
	f.Close()
}

func main(){
	file,_ := createfile("./file.txt")
	defer closefile(file)
	writefile(file)
}

