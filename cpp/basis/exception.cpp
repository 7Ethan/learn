#include <iostream>

using namespace std;

double division(double dividend,double divisor){
    if (0 == divisor ){
        throw string("除数不能为0");
    } 
    else {
        return dividend / divisor;
    }
    // return ;
}

int main(int argc, char const *argv[])
{
    int a = 0;
    int b = 0;
    cin >> a;
    cin >> b;
    
    try
    {
        division(a,b);
    }
    catch(string &str)
    {
        cout << str << endl;
    }
    
    return 0;
}
