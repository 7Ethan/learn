#include <iostream>

using namespace std;

#define Foo(x) do{\
    printf("Hello ");\
    printf("world!\n");\
}while(0)

int
main(){
    return 0;
}