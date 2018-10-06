#!/bin/bash

#using a library file the wrong way

. ./lib/my_lib.sh

result=`addem 10 15`
echo "The result is $result"
