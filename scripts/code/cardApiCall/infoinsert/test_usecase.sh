#!/bin/bash
set -Ceu

cd code/cardApiCall/infoInsert
go test -p 1 ./usecase/...