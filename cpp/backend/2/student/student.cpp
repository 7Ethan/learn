#include "student.h"
#include <iostream>

using namespace std;

void 
Student::display(){
    cout << "name : " << name << endl;
    cout << "id : " << id << endl;
    cout << "age : " << age << endl;
}

Student::Student(){
    cout << "A student object has been created ! " << endl;
}

Student::Student(int id,int age){
    this->id = id;
    this->age = age;
    // this->name = name;
}

Student::~Student(){
    cout << "An student object has been deleted." << endl;
}