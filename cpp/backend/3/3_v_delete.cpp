#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int 
main(){
    vector<int> vec;
    vec.push_back(1);
    vec.push_back(2);
    vec.push_back(3);
    vec.push_back(4);
    vec.push_back(5);
    vec.push_back(3);
    vec.push_back(3);
    vec.push_back(3);
    vec.push_back(30);
    vec.push_back(3);
    vec.push_back(23);

    // cout << vec.pop_back() << endl;

	// vector<int>::iterator iter=vec.begin();    
    for(vector<int>::iterator iter = vec.begin();iter != vec.end();){ //如果在 for 循环里进行 iter++，那么会有连续的 3 删除不了。
        if (*iter == 3){
            iter = vec.erase(iter);
        } else {
            ++iter;
        }
    }

    for(vector<int>::iterator iter=vec.begin();iter!=vec.end();iter++){
		cout<<*iter<<" ";
	}
    cout << endl;

    cout<< "------------------------" << endl;
    // int a = vec.pop_back(); //ERROR!!!!!
    vec.pop_back();
    cout << "Locate at 0:" << vec[0] <<endl; 
    return 0;
}