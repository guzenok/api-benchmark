#!/bin/bash

BINARY=api-benchmark
FORMAT=gif

for PPROF in CPU Mem; do
    for PARSER in Standart Buger; do
        echo Make ${PARSER}${PPROF}.${FORMAT}
        ./$BINARY -httpserv Gin -parser $PARSER -pprof $PPROF &
        sleep 40 # wait minute
        pkill $BINARY
        sleep 2
        go tool pprof -${FORMAT} ./$BINARY ./*.pprof > img/${PARSER}${PPROF}.${FORMAT}
        rm -f ./*.pprof
    done
done

