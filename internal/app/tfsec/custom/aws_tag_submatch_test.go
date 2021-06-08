package custom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	givenCheck(`{
  "checks": [
    {
      "code": "CUSTOM_AWSTAG",
      "description": "AWS Environment tag is required",
      "requiredTypes": [
        "resource"
      ],
      "requiredLabels": [
        "aws_instance"
      ],
      "severity": "ERROR",
      "matchSpec": {
        "name": "tags",
        "action": "contains",
        "value": "Environment",
        "subMatch": {
          "action": "contains",
          "name": "Environment",
          "value": "production"
        }
      }
    }
  ]
}
`)
}

func TestAWSTagSubMatchFailure(t *testing.T) {
	scanResults := scanTerraform(t, `
resource "aws_instance" "bastion" {
  metadata_options {
    http_endpoint               = "enabled"
    http_put_response_hop_limit = 1
    http_tokens                 = "required"
  }

  tags = {
    Environment = "test"
  }
}
`)
	assert.Len(t, scanResults, 1)
}

func TestAWSTagSubMatchSuccess(t *testing.T) {
	scanResults := scanTerraform(t, `
resource "aws_instance" "bastion" {
  metadata_options {
    http_endpoint               = "enabled"
    http_put_response_hop_limit = 1
    http_tokens                 = "required"
  }

  tags = {
    Environment = "production"
  }
}
`)
	assert.Len(t, scanResults, 0)
}
