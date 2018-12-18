#include <boost/thread/thread.hpp>
#include <iostream>

using namespace std;

// void 
// hello(string name){
//     cout << "Hello " << name << "^_^" <<endl;
// }
void 
hello(){
    cout << "Hello " << "^_^" <<endl;
}
int
main(int argc,char* argv[]){
    boost::thread child_thread(&hello);
    child_thread.join();

    return 0;
}

