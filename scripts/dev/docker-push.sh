#!/bin/bash
push() {
	REPO=$1
	FILE=$2
	DIR=$3
	docker build -t $REPO -f $FILE $DIR
	docker push $REPO
}
