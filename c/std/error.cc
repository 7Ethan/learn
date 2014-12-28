#include <stdio.h>
#include <errno.h>
#include <math.h>

int 
main(int argc,char* argv[]){
    double val;

    errno = 0;
    val = sqrt(-10);
    if(errno == EDOM){
        /**
         * C 库宏 EDOM 表示一个域错误，
         * 它在输入参数超出数学函数定义的域时发生，errno 被设置为 EDOM。
         */
        printf("Invalid value \n");
    } else {
        printf("Valid value\n");
    }

    errno = 0;
    val = sqrt(10);
    if(errno == EDOM){
        printf("Invalid value\n");
    } else {
        printf("Valid value\n");
    }

    return(0);
}