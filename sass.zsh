#!/bin/zsh
for file in static/css/*.scss 
do 
    # apparently this substitution does not work in bash lol
    # sorry for the z shell dep
        sassc  ${${file}%.*}.{scss,css} 
done
