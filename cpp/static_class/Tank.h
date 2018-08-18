#ifndef TANK_H
#define TANK_H

class Tank{
    public:
    Tank(char code);
    virtual ~Tank();
    void fire();
    static int get_count();

    private:
    static int Count;
    char  Code;
};

#endif