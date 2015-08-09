#include <iostream>
#include <vector>
#include <algorithm> // used by for_each

using namespace std;

void
print_num(int n){
    cout << n << " ";
}

int
main(){
    int a[7]={1,2,3,4,5,6,7};
    vector<int> i_vector(a,a+7);

	i_vector[5]=99;
	cout<< i_vector[5] << endl << "size : " << i_vector.size()<<endl << 
    "-------" << endl ;

    vector<int>::iterator iter; // define a iter
    for(iter = i_vector.begin();iter != i_vector.end();iter++){
        cout << *iter << ", ";
    }

    cout << endl << "another way : " << endl;

    
    for(int i = 0; i < i_vector.size(); i++){
        cout << i_vector[i] << ", ";
    }
    cout << endl;
    
    cout << endl << "another way : " << endl;
    for_each(i_vector.begin(),i_vector.end(),print_num);
    cout <<endl;

    return 0;
}