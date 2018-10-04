#ifndef SHAPE_H
#define SHAPE_H

#include <iostream>
using namespace std;

class shape{
    public:
    shape();
    // ~shape();
    virtual ~shape();   //防止内存泄漏!!!!
    virtual double calcArea();
    /**
     * ->普通函数不能是虚函数。 (必须是类的成员函数,不能是全局函数)
     * ->静态成员函数不能是虚函数。
     * -> 内联函数不能是虚函数
    */
};

#endif

// #ifndef SHAPE_H
// #define SHAPE_H

// #include <iostream>
// using namespace std;

// class Shape
// {
// public:
//     Shape();
//     virtual ~Shape();
//     virtual double calcArea();
// };

// #endif