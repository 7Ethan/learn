#include "Person.h"
#include <iostream>

Person::Person(string name){
    this->m_strName = name;
    cout << "New a person object ,name is : " << this->m_strName << endl;
}

Person::~Person(){
    cout << "Object " << this->m_strName << "has been deleted ." << endl;
}