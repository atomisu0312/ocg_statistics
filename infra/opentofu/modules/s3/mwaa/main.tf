resource "aws_s3_bucket" "mwaa" {
  bucket = var.mwaa_bucket_name

  tags = {
    Name = var.tag_base
  }
}

resource "aws_s3_bucket_versioning" "mwaa" {
  bucket = aws_s3_bucket.mwaa.bucket
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_public_access_block" "mwaa" {
  # required: https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html
  bucket                  = aws_s3_bucket.mwaa.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}