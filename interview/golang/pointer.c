#include <stdio.h>

int main()
{
    void *p = malloc(100);
    printf("%lu" ,sizeof(p));
}