#include <iostream>

using namespace std;

int 
Min(int a,int b){
    if (a < b)
        return a;
    return b;
}

int
Max(int a,int b){
    if (a > b) 
        return a;
    return b;
}

int 
main(){
    int (*f)(int,int);
    int a = 10,b = 20;
    f = Min;
    cout << "Min : " << f(a,b) << endl;
    f = Max;
    cout << "Max : " << f(a,b) << endl;
}