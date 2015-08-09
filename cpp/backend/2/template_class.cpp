#include <iostream>

using namespace std;

template <class T>
class Operation{
    public:
        Operation(T a,T b) {
            x = a;
            y = b;
        }
        T add(){
            return x + y;
        }

        T sub(){
            return x - y;
        }
    private:
        T x,y;
};

int
main(){
    // Operation<int> instance = Operation<int>(1,2);
    Operation<int> instance(1,2);
    cout << "Sum : " << instance.add() << endl;
    cout << "Sub : " << instance.sub() << endl; 
}