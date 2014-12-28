#include <stdio.h>
#include <string.h> //memmove()

int 
main (int argc,char* argv[]){
   const char dest[] = "oldstring";
   const char src[]  = "newstring";

   printf("Before memmove dest = %s, src = %s\n", dest, src);
   memmove(dest, src, 9); //要注意段错误，防止越界！！
   printf("After memmove dest = %s, src = %s\n", dest, src);

   return(0);
}