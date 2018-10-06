#!/bin/bash
#reading values from a file

file="for_in1.sh"

#更改字段分隔符，使其只能识别换行符
IFS=$'\n'

#处理长脚本是，在一个地方修改了该值，然后可能忘了修改过该值
#IFS.OLD=$IFS
#IFS=$'\n'
#具体代码
#IFS=$IFS.OLD

for state in `cat $file`
do
	echo "$state"
done

# #!/bin/bash
# # using a variable to hold the list
# list="Alabama Alaska Arizona"
# list=$list" Connecticut"
# for state in $list
# do
#         echo "Have you ever visited $state"
# done
