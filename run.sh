#!/bin/bash

i=1
while [ $i -le 100 ]
do
  curl localhost:8080/test
  i=$(( $i + 1 ))
  echo ""
done