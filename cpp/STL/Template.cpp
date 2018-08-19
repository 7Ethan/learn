#include <iostream>
using namespace std;

template<typename T>
void display(T t){
    cout << "template<typename T> : " << t <<endl; 
}

template<typename T,int ksize>
void display(T t){
    cout << "template<typename T,int ksize>" << endl;
    for (int i = 0;i < ksize;i++){
        cout << t << endl;
    }
}

template<typename T,class C>
void display(T t,C c){
    cout << " template<typename T,class C> : t ->" << t 
        << ",c->"  << c << endl;
}   

int main(int argc, char const *argv[])
{
    display(100);
    display("你好",3);
    display(true,"boolean");
    return 0;
}
