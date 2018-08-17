#########################################################################
# File Name: redirect.sh
# Description: redirect.sh
# Author: Sergey Cheung
# mail: sergeychang@gmail.com
# Created Time: 2018-08-16 10:31:16
#########################################################################
#!/bin/zsh
if [ -e ./redirect.exe ] 
then
    echo "重定向.."
    # ./redirect.exe < input_number.txt | >> output.txt
    #  在.sh脚本文件中使用管道符号不成功..
    ./redirect.exe < input_number.txt >> output.txt
    echo "程序已成功执行 ^_^ ."
else
    clang redirect.c -o redirect.exe
    chmod +x ./redirect.exe
    echo "编译成功."
 fi
