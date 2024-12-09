package main

import (
	"regexp"
	"strings"
)

// formatCorrection 规则格式统一
func formatCorrection(s string) string {
	// s = strings.TrimPrefix(s, ".")
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	// s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "/32", "", -1)
	s = strings.Replace(s, "HOST", "DOMAIN", 1)
	s = strings.Replace(s, "host", "DOMAIN", 1)
	s = strings.Replace(s, "domain", "DOMAIN", 1)
	s = strings.Replace(s, "DOMAIN-suffix", "DOMAIN-SUFFIX", 1)
	s = strings.Replace(s, "IP6-CIDR", "IP-CIDR6", 1)
	s = strings.Replace(s, "ip6-cidr", "IP-CIDR6", 1)
	s = strings.Replace(s, "ip-cidr6", "IP-CIDR6", 1)
	s = strings.Replace(s, "ip-cidr,", "IP-CIDR", 1)
	s = strings.Replace(s, "USER-agent,", "USER-AGENT", 1)
	s = strings.Replace(s, "user-agent,", "USER-AGENT", 1)
	s = strings.Replace(s, "user-AGENT,", "USER-AGENT", 1)
	s = strings.Replace(s, "ip-asn,", "IP-ASN", 1)

	return s
}

// 对域名规则按 "." 切片
func domainRuleIntercept(s string) string {
	firstInd := strings.Index(s, ".")
	return s[firstInd+1:]
}

// isNote 忽略注释和空行
func isNote(s string) bool {
	if strings.HasPrefix(s, "#") ||
		strings.HasPrefix(s, ";") ||
		strings.HasPrefix(s, "@") ||
		strings.HasPrefix(s, "\n") ||
		strings.HasPrefix(s, "//") ||
		strings.HasPrefix(s, "!") {
		return true
	}
	return false
}

// isIPv4 是否为 IPv4 规则
func isIPv4(s string) bool {
	regex := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/([1-9]|[1-2][0-9]|3[0-2]))?$`
	// `^(\d{1,3}\.){3}\d{1,3}(/\d{1,2})?$`
	ipv4Pattern := regexp.MustCompile(regex)
	return ipv4Pattern.MatchString(s)
}

// isIPv6 是否为 IPv6 规则
func isIPv6(s string) bool {
	regex := `^([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4}|:)$`
	// `^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`
	pattern := regexp.MustCompile(regex)
	return pattern.MatchString(s)
}

// isASN 是否为 ASN 规则
func isASN(s string) bool {
	pattern := regexp.MustCompile(`^[0-9]+$`)
	return pattern.MatchString(s)
}

// isDomain 是否为 域名
func isDomain(s string) bool {
	regex := `^[a-zA-Z0-9_*\-.]+\.[a-zA-Z0-9_*\-.]+$`
	pattern := regexp.MustCompile(regex)
	return pattern.MatchString(s)
}
