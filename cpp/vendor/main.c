/*************************************************************************
	> File Name: main.c
	> Description: main.c
	> Author: Sergey Cheung
	> Mail: sergeychang@gmail.com 
	> Created Time: 2018-07-22 17:09:21
 ************************************************************************/

#include<stdio.h>
#include "vendor.h"
int add_ret(int a,int b){
    return a + b;
}

int main(){
    int sum = add(1,2,add_ret);
    printf("Sum is : %d\n",sum);
    return 0;
}


