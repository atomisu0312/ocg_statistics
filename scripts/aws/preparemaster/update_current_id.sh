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

# Construct full parameter names with a leading slash
FULL_PARAM_NAME_CURRENT="/${PARAMETER_NAME_CURRENT_ID}"

# --- Get currentid ---
CURRENT_ID=$(aws ssm get-parameter --name "${FULL_PARAM_NAME_CURRENT}" --query 'Parameter.Value' --output text)
if [ $? -ne 0 ]; then
  echo "Error: Failed to retrieve currentid parameter." >&2
  exit 1
fi

NEW_VALUE=$1

# --- Update currentid ---
aws ssm put-parameter --name "${FULL_PARAM_NAME_CURRENT}" --value "$NEW_VALUE" --type "String" --overwrite
if [ $? -ne 0 ]; then
  echo "Error: Failed to update currentid parameter." >&2
  exit 1
fi

echo "Successfully updated currentid parameter."