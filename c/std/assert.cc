#include <assert.h>
#include <stdio.h>

int
main(int argc,char* argv[]){
    int a;
    char str[50];

    printf("Please input an integer : \n");
    scanf("%d",&a);
    assert(a >= 10);
    printf("The input number is : %d\n",a);

    printf("Please input a string :\n");
    // scanf("%s",&str);
    scanf("%s",str);
    assert(str != NULL);
    printf("The input string is : %s\n",str);
    return 0;
}