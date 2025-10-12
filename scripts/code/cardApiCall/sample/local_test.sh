#!/bin/bash
set -Ceu

curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name": "Some_one"}'