#include "shape.h"

shape::shape(){
    cout << "New shape object ." <<endl;
}

shape::~shape(){
    cout << "A shape object has been deleted ." << endl; 
}

double shape::calcArea(){
    cout << "shape's calcArea method is running ." << endl;    
    return 0;
}

// #include "shape.h"

// Shape::Shape()
// {
//     cout << "Shape()" << endl;
// }

// Shape::~Shape()
// {
//     cout << "~Shape()" << endl;
// }

// double Shape::calcArea()
// {
//     cout << "Shape - > calcArea()" << endl;
//     return 0;
// }