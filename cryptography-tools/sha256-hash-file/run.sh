#!/bin/sh -e
# my-go-tools sha256-hash-file run.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "************************************************************************"
    echo "* run.sh -debug (START) ************************************************"
    echo "************************************************************************"
    # set -x enables a mode of the shell where all executed commands are printed to the terminal.
    set -x
    echo " "
else
    echo "************************************************************************"
    echo "* run.sh (START) *******************************************************"
    echo "************************************************************************"
    echo " "
fi

echo "go run sha256-hash-file.go"
echo " "
go run sha256-hash-file.go testfile.txt
echo " "

echo "************************************************************************"
echo "* run.sh (END) **************************************************"
echo "************************************************************************"
echo " "
