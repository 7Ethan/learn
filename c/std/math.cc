#include <stdio.h>
#include <math.h>

int main ()
{
   int a, b;
   a = 1234;
   b = -344;
  
   printf("%d 的绝对值是 %lf\n", a, fabs(a));
   printf("%d 的绝对值是 %lf\n", b, fabs(b));
   
   return(0);
}