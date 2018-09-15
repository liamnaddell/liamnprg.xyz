#!/bin/sh

for file in "static/code"/*
do 
    export DIRPATH=$(dirname $file)
    export  FILENAME=$(basename $file .scss)
    echo $DIRPATH $FILENAME $file
    pygmentize -O full,style=perldoc -f html -o "${DIRPATH}/${FILENAME}"{.html,}
done

