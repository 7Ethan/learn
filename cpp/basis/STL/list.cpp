#include <iostream>
#include <list>

using namespace std;

int main(int argc, char const *argv[])
{
    list<int> list1;
    list1.push_back(345);
    list1.push_back(-56);
    list1.push_back(4);
    list1.push_back(76);

    list<int>::iterator itor = list1.begin();
    for (;itor !=list1.end();itor++){
        cout << "Using itor : " << *itor << endl;
    }
    return 0;
}
