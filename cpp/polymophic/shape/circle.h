#ifndef CIRCLE_H
#define CIRCLE_H

// #include <iostream>
#include "shape.h"

// using namespace std;

class circle:public shape{
    public:
    circle(double r);
    // ~circle();
    virtual ~circle();
    /**
     * 推荐把子类的析构函数也加上virtual。
     * 因为子类也有可能成为其他类的父类。
    */
    virtual double calcArea();

    private:
    double m_dR;
};

#endif

// #ifndef CIRCLE_H
// #define CIRCLE_H

// #include "shape.h"
// class Circle:public Shape
// {
// public:
//     Circle(double r);
//     ~Circle();
//     double calcArea(); // 同名且参数返回值一致
// protected:
//     double m_dR;
// };

// #endif