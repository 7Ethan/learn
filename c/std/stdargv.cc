#include <stdio.h>
#include <string.h>
#include <stdarg.h>

void 
tLog(char *sInfo, ...){
    char *str   = NULL;
    char *str1 = NULL;
    char *str2 = NULL;

    va_list sArgv;          // 申请参数列表变量

    va_start(sArgv, sInfo); // 申明最后一个传递给函数的已知的固定参数
    /* 依次读取固定参数 sInfo 后的 参数 */
    str   = va_arg(sArgv, char*);
    str1 = va_arg(sArgv, char*);
    str2 = va_arg(sArgv, char*);
    va_end(sArgv);

    printf("运行输出：[%s,%s%s%s]\n",sInfo,str,str1,str2);
}

int 
main(int argc, char *argv[]){
    char str[4+1] = {0};
    
    strcpy(str,"INFO");
    tLog(str,argv[1],argv[2],argv[3]);

    return 0;
}


/**
$ ./a.out  Runoob Google Facebook
运行输出：[INFO,RunoobGoogleFacebook]
*/