#include <iostream>
using namespace std;

class Animal{
    public:
    Animal(){
        cout << "New a animal object" << endl;
    }
    virtual ~Animal(){
        cout << "Delete a animal obejct" << endl;
    }

    void virtual eat(){
        cout << "Animal -> eat()" << endl;
    }

    void virtual move(){
        cout << "Animal -> move()" << endl;
    }
};

class Dog:public Animal{
    public:
    Dog(){
        cout << "New a dog object" << endl;
    }

    virtual ~Dog(){
        cout<< "Deleted a dog object " << endl;
    }

    void virtual eat(){
        cout <<"Dog -> eat()" << endl;
    }

    void virtual move(){
        cout << "Dog -> move" << endl;
    }
};


int main(int argc, char const *argv[])
{
    Animal *a = new Dog();
    if(NULL == a) {
        cout << "Error accur when new a Animal object" << endl;
        return -1;
    }
    a->eat();
    a->move();

    delete a;
    a = NULL;
    return 0;
}
