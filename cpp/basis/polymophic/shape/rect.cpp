#include "rect.h"

rect::rect(double w,double h){
    this->m_dWidth = w;
    this->m_dHeight = h;
    cout << "New a rect object,"<< " width->" << this->m_dWidth
    << ", height->" << this->m_dHeight << endl;
}

rect::~rect(){
    cout << "A rect object has been deleted ." <<endl;
}

double rect::calcArea(){
    cout << "calculating rect's area ." <<endl;
    return this->m_dHeight * this->m_dWidth;
}

// #include "rect.h"

// Rect::Rect(double m_dwidth, double m_dHeight)
// {
//     cout << "Rect()" << endl;
//     this->m_dHeight = m_dHeight;
//     this->m_dwidth = m_dwidth;
// }

// Rect::~Rect()
// {
//     cout << "~Rect()" << endl;
// }

// double Rect::calcArea()
// {
//     cout << "Rect::calcArea()"<< endl;
//     return m_dwidth * m_dHeight;
// }
