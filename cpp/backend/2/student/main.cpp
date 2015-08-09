#include <iostream>
#include "student.h"
using namespace std;

int
main(){
    Student* s = new Student(99,18);
    s->display();
}   