/*
csprecon - Discover new target domains using Content Security Policy

This repository is under MIT License https://github.com/tr3ss/csprecon/blob/main/LICENSE
*/

package csprecon

import (
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/tr3ss/csprecon/pkg/input"
	"github.com/projectdiscovery/mapcidr"
)

func CompileRegex(regex string) *regexp.Regexp {
	r, _ := regexp.Compile(regex)

	return r
}

func domainOk(input string, domains []string) bool {
	for _, domain := range domains {
		if len(input) > len(domain)+1 && input[len(input)-len(domain)-1:] == "."+domain {
			return true
		}
	}

	return false
}

// PrepareURL takes as input a string and prepares
// the input URL in order to get the favicon icon.
func PrepareURL(inputURL string) (string, error) {
	if len(inputURL) < MinURLLength {
		return "", input.ErrMalformedURL
	}

	if !strings.Contains(inputURL, "://") {
		inputURL = "http://" + inputURL
	}

	u, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	return u.Scheme + "://" + u.Host + u.Path, nil
}

func handleCIDRInput(inputCidr string) ([]string, error) {
	if !isCIDR(inputCidr) {
		return nil, input.ErrCidrBadFormat
	}

	ips, err := mapcidr.IPAddresses(inputCidr)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

// isCIDR determines if the given ip is a cidr range.
func isCIDR(inputCidr string) bool {
	_, _, err := net.ParseCIDR(inputCidr)
	return err == nil
}
