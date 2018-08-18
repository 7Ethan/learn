#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[])
{
    vector<int> vec;
    vec.push_back(1);
    vec.push_back(2);
    vec.push_back(3);
    
    for(int i = 0;i < vec.size();i++){
        cout << "vec[" << i << "]" << ":" << vec[i]<< endl;
    }
    vec.pop_back();////从尾部去除一个
    cout << "=================================" << endl;

    vector<int>::iterator itor = vec.begin();
    for (;itor != vec.end();itor++){
        cout << "using itor :" << *itor << endl;
    }

    // for(int i = 0;i < vec.size();i++){
    //     cout << "vec[" << i << "]" << ":" << vec[i]<< endl;
    // }
    return 0;
}
