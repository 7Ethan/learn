#include <stdio.h>

int 
add_static_val(){
    static int val = 1;
    val++;
    return val;
}

int 
main(int argc,char* argv[]){
    printf("%d\n",add_static_val());
    printf("%d\n",add_static_val());
    printf("%d\n",add_static_val());

    return 0;
}
