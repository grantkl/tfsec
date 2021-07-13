package rules

import (
	"fmt"

	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/severity"

	"github.com/aquasecurity/tfsec/pkg/provider"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"

	"github.com/aquasecurity/tfsec/pkg/rule"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
)

const AWSCloudfrontDistributionAccessLoggingEnabled = "AWS071"
const AWSCloudfrontDistributionAccessLoggingEnabledDescription = "Cloudfront distribution should have Access Logging configured"
const AWSCloudfrontDistributionAccessLoggingEnabledImpact = "Logging provides vital information about access and usage"
const AWSCloudfrontDistributionAccessLoggingEnabledResolution = "Enable logging for CloudFront distributions"
const AWSCloudfrontDistributionAccessLoggingEnabledExplanation = `
You should configure CloudFront Access Logging to create log files that contain detailed information about every user request that CloudFront receives
`
const AWSCloudfrontDistributionAccessLoggingEnabledBadExample = `
resource "aws_cloudfront_distribution" "bad_example" {
	// other config
	// no logging_config
}
`
const AWSCloudfrontDistributionAccessLoggingEnabledGoodExample = `
resource "aws_cloudfront_distribution" "good_example" {
	// other config
	logging_config {
		include_cookies = false
		bucket          = "mylogs.s3.amazonaws.com"
		prefix          = "myprefix"
	}
}
`

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		ID: AWSCloudfrontDistributionAccessLoggingEnabled,
		Documentation: rule.RuleDocumentation{
			Summary:     AWSCloudfrontDistributionAccessLoggingEnabledDescription,
			Impact:      AWSCloudfrontDistributionAccessLoggingEnabledImpact,
			Resolution:  AWSCloudfrontDistributionAccessLoggingEnabledResolution,
			Explanation: AWSCloudfrontDistributionAccessLoggingEnabledExplanation,
			BadExample:  AWSCloudfrontDistributionAccessLoggingEnabledBadExample,
			GoodExample: AWSCloudfrontDistributionAccessLoggingEnabledGoodExample,
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution#logging_config",
				"https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html",
			},
		},
		Provider:        provider.AWSProvider,
		RequiredTypes:   []string{"resource"},
		RequiredLabels:  []string{"aws_cloudfront_distribution"},
		DefaultSeverity: severity.Medium,
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context) {

			if resourceBlock.MissingChild("logging_config") {
				set.Add(
					result.New(resourceBlock).
						WithDescription(fmt.Sprintf("Resource '%s' does not have Access Logging configured", resourceBlock.FullName())).
						WithRange(resourceBlock.Range()),
				)
			}

		},
	})
}
