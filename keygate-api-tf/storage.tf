resource "aws_s3_bucket" "nft_storage" {
  bucket = var.s3_nft_image_bucket_name

  tags = {
    Name = "${var.environment_name}-${var.app_name}-nft-storage"
    Environment = var.environment_name
  }
}

resource "aws_s3_bucket_ownership_controls" "nft_storage_owner_pref" {
  bucket = aws_s3_bucket.nft_storage.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

# Create an S3 bucket policy to allow public read access
resource "aws_s3_bucket_public_access_block" "nft_storage_public_access" {
  bucket = aws_s3_bucket.nft_storage.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_acl" "nft_storage_acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.nft_storage_owner_pref,
    aws_s3_bucket_public_access_block.nft_storage_public_access,
  ]

  bucket = aws_s3_bucket.nft_storage.id
  acl    = "public-read"
}

resource "aws_s3_bucket_cors_configuration" "nft_storage_cors" {
  bucket = aws_s3_bucket.nft_storage.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }

  cors_rule {
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}