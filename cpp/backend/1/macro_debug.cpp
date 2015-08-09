#include <iostream>

using namespace std;

#define _DEBUG_

int
main(){
    int x = 10;
    
    #ifdef _DEBUG_
        cout << "File : " << __FILE__ << ",line : " << __LINE__ 
        << ",x : " << x << endl;
    #else
        printf("x = %d\n",x);
    #endif
    return 0;
}