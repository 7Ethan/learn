#include <iostream>

using namespace std;

class Base{
    public:
        Base(){
            cout << "Base::Base()" << endl;
        }
        virtual ~Base(){
            cout << "~Base::Base()" << endl;
        }
};

class Derive:public Base(){
    
}