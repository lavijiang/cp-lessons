#!/usr/bin/env bash
succtime=0
while true; do
    bash "$1" >> history.log
    if [[ $? -ne 0 ]]; then
        echo "script run $succtime times"
        break
    fi
    let succtime=succtime+1
done


