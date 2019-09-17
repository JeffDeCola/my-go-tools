#!/bin/bash
# my-go-tools set-pipeline.sh

fly -t ci set-pipeline -p my-go-tools -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
