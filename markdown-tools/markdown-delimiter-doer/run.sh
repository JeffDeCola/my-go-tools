#!/bin/sh -e
# my-go-tools markdown-delimiter-doer run.sh

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

echo "go run markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md"
echo " "
go run markdown-delimiter-doer.go -htmltable -delimiter \$\$ -i input.md -o output.md

echo " "

echo "************************************************************************"
echo "* run.sh (END) **************************************************"
echo "************************************************************************"
echo " "
