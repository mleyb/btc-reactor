data "aws_region" "current" {
}

resource "aws_lambda_function" "btc-reactor" {
  function_name    = "btc-reactor"
  filename         = "bin/btc-reactor.zip"
  handler          = "main"
  source_code_hash = filebase64sha256("bin/btc-reactor.zip")
  role             = "${aws_iam_role.btc-reactor.arn}"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 30
}

resource "aws_iam_role" "btc-reactor" {
  name               = "btc-reactor"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}