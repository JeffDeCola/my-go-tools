#!/bin/bash
# my-go-tools destroy-pipeline.sh

fly -t ci destroy-pipeline --pipeline my-go-tools
