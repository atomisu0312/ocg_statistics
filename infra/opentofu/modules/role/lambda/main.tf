# IAM Role for the Lambda function
resource "aws_iam_role" "lambda_exec" {
  name = var.lambda_exec_role_name
  tags = {
    project = var.tag_base
  }

  assume_role_policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

# IAM Policy for logging, using local variables for ARNs.
resource "aws_iam_role_policy" "lambda_policy" {
  name = var.lambda_exec_role_policy_name
  role = aws_iam_role.lambda_exec.id

  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = ["logs:CreateLogGroup"],
        Resource = var.log_group_arn_wildcard
      },
      {
        Effect   = "Allow",
        Action   = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = [var.log_stream_arn]
      }
    ]
  })
}