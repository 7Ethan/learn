#ifndef DUSTMSN_H
#define DUSTMSN_H

#include "Worker.h"

class Dustman:public Worker{
    public:
    Dustman(string name,int age);
    virtual ~Dustman();
    virtual void work();
};

#endif