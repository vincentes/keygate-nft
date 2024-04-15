// IAM Policy for keygate-api
data "aws_iam_policy_document" "keygate-api-policy-document" {
  statement {
    sid       = "AllowNFTImageUpload"
    effect    = "Allow"
    actions   = ["s3:PutObject", "s3:GetObject", "s3:DeleteObject"]
    resources = [format("%s%s", aws_s3_bucket.nft_storage.arn, "/*")]
  }
}

// IAM Policy for Session Manager access
data "aws_iam_policy_document" "session_manager_policy_document" {
  statement {
    effect = "Allow"
    actions = [
      "ssm:UpdateInstanceInformation",
      "ssmmessages:CreateControlChannel",
      "ssmmessages:CreateDataChannel",
      "ssmmessages:OpenControlChannel",
      "ssmmessages:OpenDataChannel"
    ]
    resources = ["*"]
  }
}

// IAM Role for keygate-api
data "aws_iam_policy_document" "keygate_api_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "keygate-api" {
  name               = "KeygateAPIS3Access"
  assume_role_policy = data.aws_iam_policy_document.keygate_api_assume_role_policy.json
}

// keygate-api (IAM Role) -> keygate-api-policy (IAM Policy) attachment -> S3 access
resource "aws_iam_policy" "keygate-api-policy" {
  name        = "keygate-api-policy"
  description = "Policy for Keygate API to access S3"
  policy      = data.aws_iam_policy_document.keygate-api-policy-document.json
}

resource "aws_iam_role_policy_attachment" "keygate-api" {
  role       = aws_iam_role.keygate-api.name
  policy_arn = aws_iam_policy.keygate-api-policy.arn
}

// Session Manager policy
resource "aws_iam_policy" "session_manager_policy" {
  name        = "session-manager-policy"
  description = "Policy for Session Manager access"
  policy      = data.aws_iam_policy_document.session_manager_policy_document.json
}

resource "aws_iam_role_policy_attachment" "session_manager" {
  role       = aws_iam_role.keygate-api.name
  policy_arn = aws_iam_policy.session_manager_policy.arn
}