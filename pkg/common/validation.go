package common

import (
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/hashicorp/terraform/helper/validation"
)

const dnsRegex = `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`

var DnsRegexMatcher = regexp.MustCompile(dnsRegex)

var ValidateDNSName = validation.StringMatch(DnsRegexMatcher,
	"the DNS name does not follow  RFC 952 and RFC 1123 requirements")

var ValidateDNSNameOrIP = validation.Any(validation.SingleIP(), ValidateDNSName)

func ValidateAbsPath(v interface{}, k string) (ws []string, errors []error) {
	if !filepath.IsAbs(v.(string)) {
		errors = append(errors, fmt.Errorf("%q is not an absolute path", k))
	}
	return
}

func ValidateURL(v interface{}, k string) (ws []string, errors []error) {
	if _, err := url.ParseRequestURI(v.(string)); err != nil {
		errors = append(errors, fmt.Errorf("%q does not seem a valid URL: %s", k, err))
	}
	return
}
