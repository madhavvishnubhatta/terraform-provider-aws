---
subcategory: "CodeBuild"
layout: "aws"
page_title: "AWS: aws_codebuild_source_credential"
description: |-
  Provides a CodeBuild Source Credential resource.
---

# Resource: aws_codebuild_source_credential

Provides a CodeBuild Source Credentials Resource.

~> **NOTE:
** [Codebuild only allows a single credential per given server type in a given region](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.GitHubSourceCredentials.html).
Therefore, when you define `aws_codebuild_source_credential`, [
`aws_codebuild_project` resource](/docs/providers/aws/r/codebuild_project.html) defined in the same module will use it.

## Example Usage

```terraform
resource "aws_codebuild_source_credential" "example" {
  auth_type   = "PERSONAL_ACCESS_TOKEN"
  server_type = "GITHUB"
  token       = "example"
}
```

### Bitbucket Server Usage

```terraform
resource "aws_codebuild_source_credential" "example" {
  auth_type   = "BASIC_AUTH"
  server_type = "BITBUCKET"
  token       = "example"
  user_name   = "test-user"
}
```

### AWS CodeStar Connection Usage

```terraform
resource "aws_codebuild_source_credential" "example" {
  auth_type   = "CODECONNECTIONS"
  server_type = "GITHUB"
  token       = "arn:aws:codestar-connections:us-east-1:123456789012:connection/guid-string"
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `auth_type` - (Required) The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket
  repository. Valid values are `BASIC_AUTH`,
  `PERSONAL_ACCESS_TOKEN`, `CODECONNECTIONS`, and `SECRETS_MANAGER`. An OAUTH connection is not supported by the API.
* `server_type` - (Required) The source provider used for this project.
* `token` - (Required) For a GitHub and GitHub Enterprise, this is the personal access token. For Bitbucket, this is the
  app password. When using an AWS CodeStar connection (`auth_type = "CODECONNECTIONS")`, this is an AWS CodeStar
  Connection ARN.
* `user_name` - (Optional) The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for
  other types of source providers or connections.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ARN of Source Credential.
* `arn` - The ARN of Source Credential.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to
import CodeBuild Source Credential using the CodeBuild Source Credential arn. For example:

```terraform
import {
  to = aws_codebuild_source_credential.example
  id = "arn:aws:codebuild:us-west-2:123456789:token:github"
}
```

Using `terraform import`, import CodeBuild Source Credential using the CodeBuild Source Credential arn. For example:

```console
% terraform import aws_codebuild_source_credential.example arn:aws:codebuild:us-west-2:123456789:token:github
```
