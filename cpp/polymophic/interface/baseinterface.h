#ifndef BASEINTERFACE_H
#define BASEINTERFACE_H
//用于动物的接口类
class BaseInterface{
    //是动物都要吃东西，睡觉以及爱好
public:
    virtual void Eat()=0;
    virtual void Sleep()=0;
    virtual void Hobby()=0;

    virtual ~BaseInterface(){}
};

#endif // BASEINTERFACE_H