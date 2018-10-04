#include <iostream>
#include <string>

using namespace std;

class worker {
    public:
    worker();
    worker(string city,int id);
    string getcity(void);
    ~worker();

    private:
    int job_id;
    string city;
};

class farmer{
    public:
    farmer();
    farmer(string name);
    string getname(void);
    ~farmer();

    private:
    string name;
};

worker::worker(string city,int id){
    this->job_id = id;
    this->city = city;
    cout << "New a worker," << "id:" << this->job_id << "." << endl;
}

worker::~worker(){
    cout << "A worker object has been deleted ." << endl;
}

string worker::getcity(){
    return this->city;
}

farmer::farmer(string name){
    this->name = name;
    cout << "New a farmer ,name : " << this->name << endl;
}

string farmer::getname(){
    return this->name;
}

farmer::~farmer(){
    cout << "A farmer object has been deleted."<< endl;
}


class migrantworker :public worker, virtual public farmer{
    public:
    migrantworker(string name,string city, int id);
    ~migrantworker();
    void display();
};

migrantworker::migrantworker(string name,string city,int id):worker(city,id),farmer(name){
    cout << "New a migrantworker ,name : " << this->getname() <<endl;
}

migrantworker::~migrantworker(){
    cout << this->getname() << " object has been deleted !"<< endl;
}

void migrantworker::display(){
    cout << "migrantworker's display function ~";
    cout << "Info : name-> "<< this->getname() << ",city->" << this->getcity() << endl;
}


int main(int argc, char const *argv[])
{
    migrantworker *object = new migrantworker("xiaozhang","zhanjiang",007);
    if (NULL == object){
        cout << "Error accur when new a object ~ "<< endl;
        return -1;
    }

    object->display();
    delete object;
    object = NULL;

    // worker *w = new migrantworker("aaaaaaa","ccccc",999);
    // if (NULL == w) {
    //     cout <<"Error accur when new worker ." << endl;
    //     return -1;
    // }

// 问题代码,子类是多继承,而且没有考虑用虚函数.

    // delete w;
    // w = NULL;


    return 0;
}

