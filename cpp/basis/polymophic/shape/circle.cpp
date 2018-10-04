#include "circle.h"

circle::circle(double r){
    cout << "New a circle object."<< endl;
    this->m_dR = r;
}

circle::~circle(){
    cout << "A circle object has been deleted ." << endl;
}

double circle::calcArea(){
    cout << "Calculating the circle's area."<<endl;
    return 3.14 * this->m_dR * this->m_dR; 
}

// #include "circle.h"

// Circle::Circle(double r)
// {
//     cout << "Circle()" << endl;
//     m_dR = r;
//  }

// Circle::~Circle()
// {
//     cout << "~Circle()" << endl;
// }
// double Circle::calcArea()
// {
//     cout << "Circle-->calcArea()" << endl;
//     return 3.14 * m_dR * m_dR;
// }