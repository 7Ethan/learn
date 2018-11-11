#include <stdio.h>

int main(){
    int counter = 5;
    do {
        printf("counter : %d \n",counter);
        counter--;
    } while(counter > 5);
    //特点： 先执行，后判断

    return -1;
}