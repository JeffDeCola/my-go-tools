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

# OVERVIEW ---------------------------------------------------------
echo "OVERVIEW"
echo "The test_coverage.txt file will be used by the concourse pipeline to send to slack"
echo " "
# cd into cryptography-tools and create a test_coverage.txt file
echo "cd cryptography-tools"
cd cryptography-tools
echo " "
echo "touch \"$GOPATH/coverage-results/test_coverage.txt\""
touch "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# encryptfile test -------------------------------------------------
echo "encryptfile Run go test -cover --------------------------------"
echo "cd encryptfile"
cd encryptfile
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# decryptfile test --------------------------------------------------
echo "decryptfile Run go test -cover --------------------------------"
echo "cd ../decryptfile"
cd ../decryptfile
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# md5-hash-file test -------------------------------------------------
echo "md5-hash-file Run go test -cover -------------------------------"
echo "cd ../md5-hash-file"
cd ../md5-hash-file
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# sha256-hash-file test ----------------------------------------------
echo "sha256-hash-file Run go test -cover ----------------------------"
echo "cd ../sha256-hash-file"
cd ../sha256-hash-file
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# cd into markdowntools ----------------------------------------------
echo "cd ../../markdown-tools"
cd ../../markdown-tools
echo " "

# markdown-create-table-of-contents test -----------------------------
echo "markdown-create-table-of-contents Run go test -cover -----------"
echo "cd markdown-create-table-of-contents"
cd markdown-create-table-of-contents
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# markdown-delimiter-doer test ---------------------------------------
echo "markdown-delimiter-doer Run go test -cover ---------------------"
echo "cd ../markdown-delimiter-doer"
cd ../markdown-delimiter-doer
echo "go test -cover ./... | tee test/test_coverage.txt"
go test -cover ./... | tee test/test_coverage.txt
echo "Clean test_coverage.txt file - add some whitespace to the begining of each line"
echo "sed -i -e 's/^/     /' test/test_coverage.txt"
sed -i -e 's/^/     /' test/test_coverage.txt
echo "cat \"test/test_coverage.txt\" >> \"$GOPATH/coverage-results/test_coverage.txt\""
cat "test/test_coverage.txt" >> "$GOPATH/coverage-results/test_coverage.txt"
echo " "

# ---------------------------------------------------------------------
echo "Print out test covergage file -----------------------------------"
cat "$GOPATH/coverage-results/test_coverage.txt"
echo " "

echo "unit-tests.sh (END)"
echo " "
