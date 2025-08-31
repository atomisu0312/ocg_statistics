output "mwaa_bucket_arn" {
  description = "The name of the S3 bucket for MWAA."
  value       = aws_s3_bucket.mwaa.arn
}