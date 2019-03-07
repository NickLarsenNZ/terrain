resource "aws_s3_bucket" "logs" {
  region = "${var.aws_region}"
  bucket = "logs.chirp.co.nz"
  acl    = "log-delivery-write"

  tags {
    Name        = "${lookup(var.bucket_names, var.environments[count.index], format("unknown-%d", count.index))}"
    Environment = "${var.environments[count.index]}"
    Purpose     = "logs"
    Team        = "chirp"
  }
}

resource "aws_s3_bucket" "website" {
  count = "${length(var.environments)}" # Make a bucket for each environment

  region = "${var.aws_region}"
  bucket = "${lookup(var.bucket_names, var.environments[count.index])}"
  acl    = "${var.s3_website_acl}"                                      # Todo: This needs to be a bucket-policy instead

  website {
    index_document = "${var.s3_documents["index"]}"
    error_document = "${var.s3_documents["error"]}"
  }

  versioning {
    enabled = "${var.s3_website_versioning}"
  }

  logging {
    target_bucket = "${aws_s3_bucket.logs.id}"
    target_prefix = "${lookup(var.bucket_names, var.environments[count.index])}/s3"
  }

  tags {
    Name        = "${lookup(var.bucket_names, var.environments[count.index], format("unknown-%d", count.index))}"
    Environment = "${var.environments[count.index]}"
    Purpose     = "website"
    Team        = "chirp"
  }
}
