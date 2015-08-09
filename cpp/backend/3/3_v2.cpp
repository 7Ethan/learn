#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

typedef struct rect{
    int id;
    int length;
    int width;
    bool operator< (const rect& a) const{
        if (id != a.id){
            return id < a.id;
        } else{
            if (length != a.length){
                return length < a.length;
            } else {
                return width < a.width;
            }
        }
    }
}Rect;

int
main(){
    vector<Rect> vec;
    Rect rect;
   
    {
        rect.id = 01;
        rect.length = 2;
        rect.width = 3;
    }
    vec.push_back(rect);

     {
        rect.id = 2;
        rect.length = 3;
        rect.width = 4;
    }
    vec.push_back(rect);

    // declare a iter
    vector<Rect>::iterator iter = vec.begin();
    cout<<(*iter).id<<' '<<(*iter).length<<' '<<(*iter).width<<endl;  
    cout << " -------------------- " << endl;

    sort(vec.begin(),vec.end());
    cout<<(*iter).id<<' '<<(*iter).length<<' '<<(*iter).width<<endl;  

    return 0;
}