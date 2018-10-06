#!/bin/bash

# redirecting the for output to a file
for file in /home/sergey/*
do
	if [ -d "$file" ]
	then
		echo "$file is a directory"
	else
		echo "$file is a file"
	fi
done > output.txt

# /home/sergey/25146808.jpeg is a file
# /home/sergey/anaconda3 is a directory
# /home/sergey/Desktop is a directory
# /home/sergey/Downloads is a directory
# /home/sergey/DTbyrSeUMAAiKIZ.jpg_large is a file
# /home/sergey/go is a directory
# /home/sergey/Music is a directory
# /home/sergey/node_modules is a directory
# /home/sergey/Pictures is a directory
# /home/sergey/Public is a directory
# /home/sergey/Templates is a directory
# /home/sergey/Videos is a directory
# /home/sergey/VirtualBox VMs is a directory
# /home/sergey/workspace is a directory
# /home/sergey/zdocuments is a directory

# piping a loop to another command

for state in "North Dakota" Connecticut
do
	echo "$state is next place to go"
done | sort
echo "This completes our travels"
