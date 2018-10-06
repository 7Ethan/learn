#!/bin/bash
func1(){
    echo "Trying to display a non-existent file"
    ls -l non_existent_file
}

#由于最后一条命令未执行成功，返回的状态码非0
echo "testing the function"
func1
echo "The exit status is : $?"

echo "-------------------"

func2() {
	ls -l non_existent_file
	echo "Another test to display a non-existent file"
}

#由于最后一条命令echo执行成功，返回的状态码为0
echo "Another test:"
func2
echo "The exit status is : $?"