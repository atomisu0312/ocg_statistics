#!/bin/bash
set -Ceu

# 環境変数を読み込む
set -a
source .env
set +a

AWS_ECR_REPO_NAME=$ECR_REPOSITORY_NAME_IDCHECK

# ディレクトリを移動
cd ./code/cardApiCall/cardidcheck

# ログイン
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# ビルド
docker build -t $AWS_ECR_REPO_NAME .

# タグ付け
docker tag $AWS_ECR_REPO_NAME:latest $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$AWS_ECR_REPO_NAME:latest

# プッシュ
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$AWS_ECR_REPO_NAME:latest