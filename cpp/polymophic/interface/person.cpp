#include "person.h"
#include <iostream>
using namespace std;
// #include <QDebug>

Person::Person()
{
    personName = "";
}

void Person::Eat(){
    cout <<"eat bread";
}
void Person::Sleep(){
    cout <<"Sleep 1 hour";
}
void Person::Hobby(){
    cout <<"run 35 mins";
}
void Person::Speak(){
    cout <<"person speak a lot of language";
}
void Person::SetName(string pName){
    personName = pName;
    cout <<personName;
}
string Person::GetName(){
    return personName;
}