#include <iostream>
#include <map>

using namespace std;

int main(int argc, char const *argv[])
{
    map<int,string> m;
    pair<int,string> p1(3,"hello"); //pair原型是: template<class _T1, class _T2>
    pair<int,string> p2(6,"world");

    m.insert(p1);
    m.insert(p2);
    // cout << m[3] << endl;
    // cout << m[6] << endl;

    map<int,string>::iterator itor = m.begin();
    for (;itor != m.end();itor++)
    {
        //cout << *itor << endl;//错误，每个都包含keyvalue
        cout << itor->first << endl;
        cout << itor->second << endl;
    }
    return 0;
}
