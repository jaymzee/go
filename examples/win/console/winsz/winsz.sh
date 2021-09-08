#!/bin/sh
#termset=$(stty -g)
#stty -icanon -echo
printf '\033[6n'
read -d "R" rowscols
rowscols="${rowscols//[^0-9;]/}"
rowscols=("${rowscols//;/ }")
printf '(row %d, column %d)\n' ${rowscols[0]} ${rowscols[1]}

printf '\033[14t'
read -d "t" xyres
xyres="${xyres//[^0-9;]/}"
xyres=("${xyres//;/ }")
printf '(%d, yres %d, xres %d)\n' ${xyres[0]} ${xyres[1]} ${xyres[1]}

#stty $termset
