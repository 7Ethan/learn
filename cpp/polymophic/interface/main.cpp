#include <iostream>
#include "baseinterface.h"
#include "person.h"

using namespace std;

int main(int argc, char const *argv[])
{
    BaseInterface *pBaseInterface = new Person();
    pBaseInterface->Sleep();
    return 0;
}
