#include <iostream>
#include "Tank.h"

using namespace std;

int Tank::Count = 10; //构造函数之外，单独初始化

Tank::Tank(char code){
    cout << "New a Tank object." << endl;
    Code = code;    
    Count++;    //千万别忘了对静态成员操作
}

Tank::~Tank(){
    --Count;    ////千万别忘了对静态成员操作
    cout << "Delete a tank object ." << endl;
}

void Tank::fire(){
    cout << "Tank ->fire()." <<endl;
}

int Tank::get_count(){
    cout << "get_cout() 方法被调用" << endl;
    // return this->Count; 只有非静态方法才能试试 this 
    return Count;
}

