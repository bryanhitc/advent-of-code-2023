#!/bin/bash

set -e

INPUTS_DIR="../inputs"
NUM_DAYS=25
INVALID_SESSION_ERR=$'Bash variable `AOC_SESSION` must be set to a valid 128 hexadecimal hash.\n\nRetrieve your session token by logging into https://adventofcode.com => right click on the page => inspect => `Application` tab => `Cookies` => see `session` value.'

if [[ "${#AOC_SESSION}" -ne 128 ]]
then
    echo "$EMPTY_SESSION_ERR"
    exit 1
fi

for ((i=1;i<=NUM_DAYS;++i))
do
    if test -f "$INPUTS_DIR/day$i/problem.txt"
    then
        continue
    fi

    echo $"Downloading day$i..."
    problem=$(curl -s "https://adventofcode.com/2023/day/$i/input" -H \
        "Cookie: session=$AOC_SESSION")
    if [[ "$problem" =~ ^"Please don't repeatedly request" ]]
    then
        echo $"Day$i not ready yet."
        exit 0
    fi

    mkdir -p "$INPUTS_DIR/day$i"
    echo "$problem" > "$INPUTS_DIR/day$i/problem.txt"
done
