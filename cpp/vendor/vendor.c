/*************************************************************************
	> File Name: vendor.c
	> Description: vendor.c
	> Author: Sergey Cheung
	> Mail: sergeychang@gmail.com 
	> Created Time: 2018-07-22 17:12:49
 ************************************************************************/

#include "vendor.h"

int add(int a,int b,int (*add_value)()){
    return (*add_value)(a,b);
}
