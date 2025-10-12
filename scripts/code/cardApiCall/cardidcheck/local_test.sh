#!/bin/bash
set -Ceu

curl "http://localhost:9001/2015-03-31/functions/function/invocations" -d '{"id": "10000"}'