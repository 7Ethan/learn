#include <stdio.h>
#include <setjmp.h>

static jmp_buf buf;

void 
second(void) {
    printf("second\n");         // 打印
    longjmp(buf,1);             // 跳回setjmp的调用处 - 使得setjmp返回值为1
}

void 
first(void) {
    second();
    printf("first\n");          // 不可能执行到此行
}

int 
main(int argc,char* argv[]) {   
    /**
     * setjmp()
     * 这个宏把当前环境保存在变量 environment 中，以便函数 longjmp() 后续使用。
     * 如果这个宏直接从宏调用中返回，则它会返回零，
     * 但是如果它从 longjmp() 函数调用中返回，则它会返回一个非零值。
    */
    if ( ! setjmp(buf) ) {
        first();                // 进入此行前，setjmp返回0
    } else {                    // 当longjmp跳转回，setjmp返回1，因此进入此行
        printf("main\n");       // 打印
    }

    return 0;
}

/***
 * 
 * second
 * main
*/