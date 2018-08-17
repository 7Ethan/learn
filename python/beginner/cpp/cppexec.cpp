#include <iostream>

using namespace std;

int add_func(int a,int b){
    return a + b;
}

int main(int argc, char const *argv[])
{
    int a,b;
    cout << "Please input two number : " << endl;
    cin >> a >> b;
    cout << "the c/c++ run result :" << endl;
    cout << add_func(a,b) << endl;
    return 0;
}
