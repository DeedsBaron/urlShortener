#!/bin/bash
if [[ $SOLUTION == 0 ]]; then
  ./urlShortener
elif [[ $SOLUTION == 1 ]]; then
  ./urlShortener -mem psql
fi
while :
do
done