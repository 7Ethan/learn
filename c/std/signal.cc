/**
 * C 库函数 
 *      void (*signal(int sig, void (*func)(int)))(int) 
 * 设置一个函数来处理信号，即带有 sig 参数的信号处理程序。
 * */

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>

void 
sighandler(int signum){
   printf("捕获信号 %d，跳出...\n", signum);
   exit(1);
}

int 
main(int argc,char* argv[])
{
   signal(SIGINT, sighandler);

   while(1) {
      printf("开始休眠...（请发出中断信号停止休眠）\n");
      sleep(1);
   }

   return(0);
}

