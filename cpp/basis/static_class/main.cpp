#include "Tank.h"
#include <stdlib.h>
#include <iostream>
using namespace std;

int main()
{
    cout << Tank::get_count() << endl;
    //在类实例化之前就能使用
    Tank t1('A');
    cout << Tank::get_count() << endl;
    //初值10变成11
    cout << t1.get_count() << endl;

    // 堆上实例化自己管理内存
    Tank *p = new Tank('B');
    cout << Tank::get_count() << endl;
    Tank *q = new Tank('C');
    cout << q->get_count() << endl;

    delete p;
    delete q;

    cout << Tank::get_count() << endl;
    system("pause");
    return 0;
}