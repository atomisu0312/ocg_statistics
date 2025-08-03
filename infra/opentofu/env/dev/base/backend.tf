terraform {
  backend "s3" {
    bucket = "${var.project}-terraform-state"
    key    = "dev/base/terraform.tfstate"
    region = var.region
    
    # バージョニングが有効になっているため、状態ファイルの履歴を保持
    # 暗号化も有効になっているため、状態ファイルは暗号化される
    
    # ロック機能（DynamoDB）を使用する場合は以下を追加
    # dynamodb_table = "terraform-locks"
    # encrypt        = true
  }
} 