#include "Worker.h"
#include <iostream>
using namespace std;

Worker::Worker(string name,int age):Person(name){
    this->age = age;
    cout << "New a worker object .Name is : " << 
        name << ",Age is : " << this->age <<endl;
}

Worker::~Worker(){
    cout << "A worker obejct has been deleted." << endl;
}

