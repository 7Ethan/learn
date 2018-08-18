// #include <iostream>
// #include "Coordinate.h"

// using namespace std;

// int main(int argc, char const *argv[])
// {
//     Coordinate c(1,-200);
//     cout << "c's X:" << c.getX() << ",and c's Y : " << c.getY() << endl;
//     cout << " =====  Overload ======= " << endl;
//     -c;
//     cout << "c's X:" << c.getX() << ",and c's Y : " << c.getY() << endl;

// /**
//  * New a Coordinate object.
// c's X:1,and c's Y : -200
//  =====  Overload ======= 
// c's X:-1,and c's Y : 200
// A Coordinate object has been deleted . 
// */

//     return 0;
// }



#include "Coordinate.h"
#include <iostream>
#include <stdlib.h>
using namespace std;

int main()
{
    Coordinate coor1(1, 3);
    Coordinate coor2(2, 4);
    Coordinate coor3(0, 0);

    coor3 = coor1 + coor2;

    //cout << coor3.getX() << "," << coor3.getY() << endl;

    cout << coor3[0] <<endl;
    cout << coor3[1] << endl;
    

    system("pause");
    return 0;
}