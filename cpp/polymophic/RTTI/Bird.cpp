#include "Bird.h"
#include <iostream>
using namespace std;

void Bird::takeoff(){
    cout << "Bird ->takeoff()" << endl;
}

void Bird::land(){
    cout << "Bird -> land()" <<endl;
}

void Bird::foraging(){
    cout << "Bird -> foraging()" <<endl;
}

Bird::~Bird(){
    cout << "A (bird) object has been deleted ." << endl;
}