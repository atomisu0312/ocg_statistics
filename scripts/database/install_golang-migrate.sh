#!/bin/bash
set -Ceu

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3