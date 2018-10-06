#include <iostream>
#include <thread>

using namespace std;

void my_print(){
    cout << "my print function ." << endl;
}

int main(int argc, char const *argv[])
{
    /* clang++ -std=c++11 function_thread.cpp -pthread */
    thread t1(my_print);
    t1.join();
    cout << "main thread >>> " << endl;
    return 0;
}
