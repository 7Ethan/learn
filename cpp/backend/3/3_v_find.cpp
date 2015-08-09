#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(){
    vector<int> vec;
    vec.push_back(1);
    vec.push_back(2);
    vec.push_back(3);
    vec.push_back(4);
    vec.push_back(5);

    vector<int>::iterator iter = find(vec.begin(),vec.end(),100);
    if (iter != vec.end()){
        cout << "Found! " << endl;
    } else {
        cout << "Not found ! " << endl;
    }
    return 0;
}