#include <iostream>

using namespace std;

class Person {
    public:
    Person():m_Name("###"){};
    // Person();
    virtual ~Person(){
        cout << "A person onject has been deleted." <<endl;
    };
    virtual void  eat() = 0;
    virtual void sleep() = 0;
    private:
    string m_Name;
};

class worker:public Person{
    public:
    worker();
    ~worker();

    // 继承接口的方法,一定要显式声明!!!
    void  eat();
    void sleep();

    string get_name();
    private:
    string m_Name;
};

// worker::~Person(){
//     cout << "Delete person" << endl;
// }

// worker::Person(){

// }

worker::worker(){
    // m_Name = "";
}

worker::~worker(){
    cout << "delete a worker object" << endl;
}

void worker::eat(){
    cout << "working is eating..." << endl;
}

void worker::sleep(){
    cout << "worker is sleeping" << endl;
}

string worker::get_name(){
    return m_Name;
}

int main(int argc, char const *argv[])
{
    Person *p = new worker();
    worker *w = new worker();
    p->eat();
    p->sleep();
    cout << "================="<< endl;
    // cout << w->get_name();
    // w->sleep();

    cout << "================="<< endl;
    delete p;
    p = NULL;

    cout << ">>>>>>>>>>>>>>>>>>>>"<< endl;
    delete w;
    p = NULL;

    return 0;
}