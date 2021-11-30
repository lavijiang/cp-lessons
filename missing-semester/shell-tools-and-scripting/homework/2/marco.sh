#!/bin/bash
marco(){
    export last_dir=$(pwd)
    echo $last_dir
}

polo(){
    cd $last_dir
    echo $last_dir
}
