#!/usr/bin/env bash
find $1 -name '*.html' | xargs -d '\n' zip res.zip 
