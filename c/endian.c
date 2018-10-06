#include <stdio.h>

typedef union{
    short a;
    char b[sizeof(short)];
} MODULE; // MODULE is an alia

int 
main(int argc, char const *argv[])
{
    MODULE m;
    m.a = 0x0102;
    if (m.b[0] == 0x01 && m.b[1] == 0x02) {
        printf("big endian .\n");
    } else if(m.b[0] == 0x02 && m.b[1] == 0x01) {
        printf("samll endian.\n");
    } else {
        printf("Unknown.\n");
    }
    return 0;
}
