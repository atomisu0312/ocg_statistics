#!/bin/bash
set -Ceu

# This script retrieves specific parameters from AWS SSM Parameter Store.

# Load environment variables from .env file
if [ -f .env ]; then
  set -a
  source .env
  set +a
else
  echo "Error: .env file not found." >&2
  exit 1
fi

echo "Retrieving SSM parameters..."

# Construct full parameter names with a leading slash
FULL_PARAM_NAME_DELTA="/${PARAMETER_NAME_DELTA_ID}"
FULL_PARAM_NAME_CURRENT="/${PARAMETER_NAME_CURRENT_ID}"
FULL_PARAM_NAME_MAX="/${PARAMETER_NAME_MAX_ID}"

# --- Get deltaid ---
DELTA_ID=$(aws ssm get-parameter --name "${FULL_PARAM_NAME_DELTA}" --query 'Parameter.Value' --output text)
if [ $? -ne 0 ]; then
  echo "Error: Failed to retrieve deltaid parameter." >&2
  exit 1
fi

# --- Get currentid ---
CURRENT_ID=$(aws ssm get-parameter --name "${FULL_PARAM_NAME_CURRENT}" --query 'Parameter.Value' --output text)
if [ $? -ne 0 ]; then
  echo "Error: Failed to retrieve currentid parameter." >&2
  exit 1
fi

# --- Get maxid ---
MAX_ID=$(aws ssm get-parameter --name "${FULL_PARAM_NAME_MAX}" --query 'Parameter.Value' --output text)
if [ $? -ne 0 ]; then
  echo "Error: Failed to retrieve maxid parameter." >&2
  exit 1
fi

echo "Successfully retrieved parameters:"
echo "deltaid: $DELTA_ID"
echo "currentid: $CURRENT_ID"
echo "maxid: $MAX_ID"