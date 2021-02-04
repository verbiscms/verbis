#!/bin/bash

find . -type f -name '*.go' | while read i;
do
	echo $i
 	if ! grep -q Copyright $i
  	then
    cat copyright.txt $i >$i.new && mv $i.new $i
  	fi
done