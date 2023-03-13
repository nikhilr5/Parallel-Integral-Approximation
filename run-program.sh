#!/bin/bash
if [[ ("$#" -ne 2) && ("$#" -ne 3)]]; then
    echo "Must enter either two or three parameters"
    echo "If you enter two the program will run sequentially and use arguments as the bounds"
    echo "If enter three the program will run parallel and use first argument as threadcount and other two as bounds for integral"
    exit 2
fi
cd integral-approximation

if [[ "$#" -eq 3 ]];  then
    go run parallel/parallel.go  $1 $2 $3
    exit 2
fi

go run sequential/sequential.go  $1 $2