#include <stdio.h>

int 
main(int argc,char* argv[]){
    int a[10] = {10 * 1};   //只初始化第一个
    int b = 88;
    a[2] = b;

    printf("a[2] = %d\n",a[2]);
    printf("a[2] = %d\n",a[1]);
    printf("a[2] = %d\n",a[0]);
    
    return 0;
}