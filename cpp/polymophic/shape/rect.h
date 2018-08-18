#ifndef RECT_H
#define RECT_H

#include "shape.h"
// using namespace std;

class rect:public shape{
    public:
    rect(double width,double height);
    // ~rect();
    virtual ~rect();
    /**
     * 推荐把子类的析构函数也加上virtual。
     * 因为子类也有可能成为其他类的父类。
    */
    virtual double calcArea();

    private:
    double m_dWidth;
    double m_dHeight;
};

#endif

// #ifndef RECT_H
// #define RECT_H

// #include "shape.h"
// class Rect : public Shape
// {
// public:
//     Rect(double width,double height);
//     ~Rect();
//     double calcArea();

// protected:
//     double m_dwidth;
//     double m_dHeight;
// };

// #endif // RECT_H