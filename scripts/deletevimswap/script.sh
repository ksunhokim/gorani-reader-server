#!/bin/bash
FILES=`find ../.. -name "*.swp"`
for f in $FILES
do
	echo "$f"
	rm "$f"
done

