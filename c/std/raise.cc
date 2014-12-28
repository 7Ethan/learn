/**
 * C 库函数 int raise(int sig) 会促使生成信号 sig。sig 参数与 SIG 宏兼容。
*/

#include <signal.h>
#include <stdio.h>
#include <stdlib.h>

void 
signal_catchfunc(int signal){
    printf("!! 信号捕获 !!\n");
}

int 
main(int argc,char* argv[]){
    // int ret;
    __sighandler_t ret;

    ret = signal(SIGINT, signal_catchfunc);

    if( ret == SIG_ERR) 
    {
        printf("错误：不能设置信号处理程序。\n");
        exit(0);
    }

    printf("开始生成一个信号\n");
    // ret = raise(SIGINT);
    /**
     * warning: incompatible integer to pointer conversion assigning to '__sighandler_t' (aka 'void (*)(int)') from 'int' [-Wint-conversion]
    ret = raise(SIGINT);
    */
    int new_ret = raise(SIGINT);


    if( new_ret !=0 ) 
    {
        printf("错误：不能生成 SIGINT 信号。\n");
        exit(0);
    }

    printf("退出...\n");
    return(0);
}

/**
 * 
开始生成一个信号
!! 信号捕获 !!
退出...
*/

