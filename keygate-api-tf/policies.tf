

// IAM User for keygate-api
resource "aws_iam_user" "keygate-api-instance" {
  name = "keygate-api"
  path = "/"
}

resource "aws_iam_access_key" "keygate-api-instance" {
  user = aws_iam_user.lb.name
}

// IAM Role for keygate-api
data "aws_iam_role" "example" {
  name = "KeygateAPIS3Access"
}


// IAM Policy for keygate-api
data "aws_iam_policy_document" "keygate-api-policy-document" {
  statement {
    sid = "AllowNFTImageUpload"
    effect = "Allow"
    actions   = ["s3:PutObject", "s3:GetObject", "DeleteObject"]
    resources = [aws_s3_bucket.nft_storage.arn + "/*"]
  }
}


resource "aws_iam_policy" "keygate-api-policy" {
  name       = "keygate-api-policy"
  description = "Policy for Keygate API to access S3"
  policy      = data.aws_iam_policy_document.keygate-api-policy-document.json
}

// Attach the policy to the user
resource "aws_iam_user_policy_attachment" "keygate-api-instance" {
  user       = aws_iam_user.keygate-api-instance.name
  policy_arn = aws_iam_policy.policy.arn
}

