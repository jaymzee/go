#!/bin/sh
#termset=$(stty -g)
#stty -icanon -echo

IFS=";"
printf '\033[s\033[999;999H\033[6n'
read -d "R" rowscols
printf '\033[u'
printf '\033[14t'
read -d "t" xyres
#stty $termset

rowscols="${rowscols//[^0-9;]/}"
rowscols=($rowscols)
xyres="${xyres//[^0-9;]/}"
xyres=($xyres)

echo ${rowscols[0]} ${rowscols[1]} ${xyres[2]} ${xyres[1]}
