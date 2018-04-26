#!/bin/sh
for file in static/css/*.scss 
do 
    # apparently this substitution does not work in bash lol
    # sorry for the z shell dep
#        sassc  ${${file}%.*}.{scss,css} 

    export DIRPATH=$(dirname $file)
    export  FILENAME=$(basename $file .scss)
    echo $DIRPATH $FILENAME $file
    sassc "${DIRPATH}/${FILENAME}".{scss,css}
done

