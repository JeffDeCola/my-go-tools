#!/bin/sh
# my-go-tools unit-test.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "unit-tests.sh -debug (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status. Needed for concourse.
    # set -x enables a mode of the shell where all executed commands are printed to the terminal.
    set -e -x
    echo " "
else
    echo "unit-tests.sh (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status.  Needed for concourse.
    set -e
    echo " "
fi

echo "The goal is to set up a go src/github.com/JeffDeCola/my-go-tools directory"
echo "Then tests will be run in that directory"
echo "Test coverage results, text_coverage.txt, will be moved to /coverage-results directory"
echo " "

echo "At start, you should be in a /tmp/build/xxxxx directory with two folders:"
echo "   /my-go-tools"
echo "   /coverage-results (created in task-unit-test.yml task file)"
echo " "

echo "pwd is: $PWD"
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "Setup the GOPATH based on current directory"
export GOPATH=$PWD
echo " "

echo "Now we must move our code from the current directory ./my-go-tools to" 
echo "$GOPATH/src/github.com/JeffDeCola/my-go-tools"
mkdir -p src/github.com/JeffDeCola/
cp -R ./my-go-tools src/github.com/JeffDeCola/.
echo " "

echo "cd src/github.com/JeffDeCola/my-go-tools"
cd src/github.com/JeffDeCola/my-go-tools
echo " "

echo "Check that you are set and everything is in the right place for go:"
echo "gopath is: $GOPATH"
echo "pwd is: $PWD"
ls -la
echo " "

echo "Run go test -cover"
echo "   -cover shows the percentage coverage"
echo "   Put results in /testcode/test_coverage.txt file"
mkdir testcode
# go test -cover ./... | tee testcode/test_coverage.txt
echo "Placeholder to run go tests for my-go-examples" | tee testcode/test_coverage.txt
echo " "

echo "Clean testcode/test_coverage.txt file - add some whitespace to the begining of each line"
sed -i -e 's/^/     /' testcode/test_coverage.txt
echo " "

echo "The test_coverage.txt file will be used by the concourse pipeline to send to slack"
echo " "

echo "Move testcode/text_coverage.txt to /coverage-results directory"
mv "testcode/test_coverage.txt" "$GOPATH/coverage-results/"
echo " "

echo "unit-tests.sh (END)"
echo " "
