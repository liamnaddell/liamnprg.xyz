#!/bin/sh
for file in static/css/*.scss 
do 
    # apparently this substitution does not work in bash lol
    # sorry for the z shell dep
#        sassc  ${${file}%.*}.{scss,css} 

    export DIRPATH=$(dirname $file)
    export  FILENAME=$(basename $file .scss)
    echo $DIRPATH $FILENAME $file
    if [ ${PREFIX} == "" ]; then
	    sassc "${DIRPATH}/${FILENAME}".{scss,css}
    else
	    sassc "${PREFIX}/${DIRPATH}/${FILENAME}".scss "${PREFIX}/${DIRPATH}/${FILENAME}".css
    fi 
done

