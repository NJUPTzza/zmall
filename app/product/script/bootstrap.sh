#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=product
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}