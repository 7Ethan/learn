#!/bin/sh

old_suffix="c"
new_suffix="cc"

dir=$(eval pwd)

for file in $(ls ${dir} | grep .${old_suffix})
    do
        # echo $file
        name=$(echo $file | cut -d . -f 1)
        # echo $name
        mv ${file} ${name}.${new_suffix}
    done