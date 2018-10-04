#include <stdio.h>

char *ip = "192.168.001.001";

void int_to_ipv4(int tmp, char *ip)
{
    unsigned short ip1, ip2, ip3, ip4;
    ip1 = tmp & 0xff;
    ip2 = (tmp >> 8) & 0xff;
    ip3 = (tmp >> 16) & 0xff;
    ip4 = (tmp >> 24) & 0xff;
    sprintf(ip, "%03d.%03d.%03d.%03d", (int)ip1, (int)ip2, (int)ip3, (int)ip4);
}

int ipv4_to_int(char *ip)
{
    int tmp = 0;
    char ip1, ip2, ip3, ip4;

    ip1 = atoi(ip);
    ip = strchr(ip, '.');
    if(!ip)
        return -1;
    ip2 = atoi(++ip);
    ip = strchr(ip, '.');
    if(!ip)
        return -1;
    ip3 = atoi(++ip);
    ip = strchr(ip, '.');
    if(!ip)
        return -1;
    ip4 = atoi(++ip);

    tmp |= ip4 & 0xff;
    tmp = (tmp << 8) | (ip3 & 0xff);
    tmp = (tmp << 8) | (ip2 & 0xff);
    tmp = (tmp << 8) | (ip1 & 0xff);

    return tmp;
}

int main(){
    int_to_ipv4(16820416,ip);
    // int result = ipv4_to_int(ip);
    // printf("%x\n",result);
    return 0;
}