#include <iostream>
#include <string>

using namespace std;

int 
main(){
    char str[] = "I am a programmer." ;//str 是一个字符数组
    char * str1="abc"; //str1是一个字符指针变量，可以指向一个字符串
    // 不能通过 str1 指针改变它指向的字符串常量。 “abc” 存放在常量区

    char * str2[] = {"hello world","good bye"};
    string str3 = "I am a programmer, too.";

    cout<<"str: "<<str<<endl;
    cout<<"str1: "<<str1<<endl;   

    str2[0] = "hello cpp!>>>>>>>>>>>>>>>>";
    cout<<"str2[0]: "<<str2[0]<<endl;
    cout<<"str3: "<<str3<<endl;
    return 0;
}