#include "Dustman.h"
#include <iostream>

using namespace std;

// void Dustman::work(){
//     cout << "Dustman -> work()" << endl;
// }

Dustman::Dustman(string name,int age):Worker(name,age){
    cout << "New a dustman object ." << endl;
}

Dustman::~Dustman(){
    cout << "A dustman object has been deleted ." << endl;
}

void Dustman::work(){
    cout << "清洁工人扫地." << endl;
}