#include <stdarg.h>
#include <stdio.h>

int 
mul(int num_args, ...){
   int val = 1;
   va_list ap;
   /**
    * 这是一个适用于 va_start()、va_arg() 和 va_end() 这三个宏存储信息的类型。
   */

   int i;

   va_start(ap, num_args);
   for(i = 0; i < num_args; i++) {
      val *= va_arg(ap, int);
   }
   va_end(ap);

   return val;
}

int 
main(int argc,char* argv[]){
   printf("15 * 12 = %d\n",  mul(2, 15, 12) );
   
   return 0;
}
