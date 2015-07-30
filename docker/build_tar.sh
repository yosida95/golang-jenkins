#!/bin/sh
set -x
set -e
set -u

tar -cf Dockerfile.tar Dockerfile test_data
