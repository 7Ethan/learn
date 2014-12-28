/**
 * locale.h 头文件定义了特定地域的设置，比如日期格式和货币符号。
*/
#include <locale.h>
#include <stdio.h>

int main (int argc,char* argv[]){
   struct lconv * lc;

   setlocale(LC_MONETARY, "it_IT");
   lc = localeconv();
   printf("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf("International Currency Symbol: %s\n",lc->int_curr_symbol);

   setlocale(LC_MONETARY, "en_US");
   lc = localeconv();
   printf("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf("International Currency Symbol: %s\n",lc->int_curr_symbol);

   setlocale(LC_MONETARY, "en_GB");
   lc = localeconv();
   printf ("Local Currency Symbol: %s\n",lc->currency_symbol);
   printf ("International Currency Symbol: %s\n",lc->int_curr_symbol);

   printf("Decimal Point = %s\n", lc->decimal_point);

   return 0;
}

/**
 * 
Local Currency Symbol: EUR
International Currency Symbol: EUR
Local Currency Symbol: $
International Currency Symbol: USD
Local Currency Symbol: £
International Currency Symbol: GBP
Decimal Point = .
*/