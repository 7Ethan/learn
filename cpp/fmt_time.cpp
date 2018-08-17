#include <iostream>
#include <ctime>
#include <stdlib.h>

using namespace std;

string current_date();

int main(int argc, char const *argv[])
{
    cout <<current_date().c_str() << endl;
    getwchar();
    return 0;
}

string current_date(){
    time_t now = time(NULL);
    char tmp[64];
    cout << "len(tmp):" << strlen(tmp) << endl;
    strftime(tmp,sizeof(tmp),"%Y-%m-%d",localtime(&now));
    return tmp;
}
