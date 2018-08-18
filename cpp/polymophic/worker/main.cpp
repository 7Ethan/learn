#include <iostream>
#include "Person.h"
#include "Worker.h"
// #include <stdlib.h>
#include "Dustman.h"

int main(int argc, char const *argv[])
{
    Dustman d("老王",55);
    d.work();
    return 0;
}
