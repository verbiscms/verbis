#!/bin/bash

find . -type f -name '*.go' | while read i;
do
	echo "{"
		echo "file:$i"
		contents=$( cat $i )
		echo "contents:"
		echo "\"$contents\""
	echo "}"
done
