/**
 * Run-Time Type Identification
*/

/**
 * dynamic_cast注意事项：
 * -> 只能应用于指针和引用的转换
 * -> 要转换的类型中必须包含虚函数
 * -> 转换成功返回子类的地址，失败返回NULL
 * 
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * 
 * typeid注意事项：
 * -> type_id返回一个type_info对象的引用
 * -> 如果想通过基类的指针获得派生类的数据类型,基类必须带有虚函数
 * -> 只能获取对象的实际类型。(不能判断当前指针是基类还是子类)
*/

/**
 * 继承关系不是RTTI的充分条件，只是必要条件，所以存在继承关系的类不一定可以用RTTI技术
 * => RTTI的含义是运行时类型识别
 * => RTTI技术可以通过父类指针识别其所指向对象的真实数据类型
 * => 运行时类型别必须建立在虚函数的基础上，否则无需RTTI技术
 * 
*/

#include <iostream>
#include "Bird.h"
#include "Plane.h"
using namespace std;
#include <stdlib.h>
#include <typeinfo>

void doSomething(Flyable *obj)
{
    cout << "Print object type info:" << typeid(*obj).name() << endl;
    obj->takeoff();
    if (typeid(*obj) == typeid(Bird)){
        Bird *b = dynamic_cast<Bird *>(obj);
        b->foraging();
    }  
    if(typeid(*obj) == typeid(Plane)){
        Plane *p = dynamic_cast<Plane *>(obj);
        p->carry();
    }

    obj->land();
}

int main()
{
    Bird *b  = new Bird();
    if (NULL == b) {
        cout << "Error ." << endl;
        return -1;
    }
    doSomething(b);

    // !!!!!!!!!! 必须手动释放内存!!!!
    delete b;
    b = NULL;

    cout << "==================================" << endl;
    Plane p;
    doSomething(&p);
    system("pause");
    return 0;
}

