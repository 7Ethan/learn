#ifndef PERSON_H
#define PERSON_H
// #include <QString>
#include <iostream>
using namespace std;
#include "baseinterface.h"
//继承动物的接口类
class Person : public BaseInterface
{

public:
    Person();
//继承动物接口后一定要实现接口类中的接口
    void Eat();
    void Sleep();
    void Hobby();
//但是人类作为动物有自己的行为和爱好
private:
    string personName;
public:
    void Speak();
    void SetName(string pName);
    string GetName();
};

#endif // PERSON_H