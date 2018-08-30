#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
    int* arr = new int[10];
    if (arr == NULL){
        cout << "Error in new an object ! " << endl;
        return -1;
    }

    for (int i = 0; i <= 15;i++){
        cout << arr[i] << "  ";
    }

    delete arr;
    arr = NULL;
    
    return 0;
}
