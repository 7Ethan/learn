#########################################################################
# File Name: redirect_to_file.sh
# Description: redirect_to_file.sh
# Author: Sergey Cheung
# mail: sergeychang@gmail.com
# Created Time: 2018-08-16 12:37:42
#########################################################################
#!/bin/bash
today=`date +%y%m%d`
ls /usr/local/bin/ -al >> test_output.$today
