#!/bin/bash

# Compiles static files into binary
# and runs go install

staticfiles -o files/statics/statics.go Web/statics
staticfiles -o files/templates/templates.go Web/templates

cd $GOPATH/src/github.com/DCNT-developer/dcnt
go install -ldflags "-X github.com/DCNT-developer/dcnt/engine.Build=`git rev-parse HEAD` -X github.com/DCNT-developer/dcnt/engine.dcntVersion=`cat VERSION`" || cerr=1