#!/bin/bash
set -Ceu

startId=$1
delta=$2
curl "http://localhost:9002/2015-03-31/functions/function/invocations" -d '{"startId": '$startId', "delta": '$delta'}'