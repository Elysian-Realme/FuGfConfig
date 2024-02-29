package util

import "strings"

// FormatCorrection 规则格式统一
func FormatCorrection(s string) string {
	// s = strings.TrimPrefix(s, ".")
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	//s = strings.Replace(s, " ", "", -1)
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

// IsNote 忽略注释和空行
func IsNote(s string) bool {
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

// CleanAll 只返回域名或者 IP
func CleanAll(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	if strings.HasPrefix(s, "*.") {
		s = strings.TrimPrefix(s, "*")
	}
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "||", "", -1)
	s = strings.Replace(s, "^", "", -1)
	s = strings.Replace(s, "127.0.0.1", "", -1)
	s = strings.Replace(s, "0.0.0.0", "", -1)
	s = strings.Replace(s, "::", "", -1)
	s = strings.Replace(s, ",no-resolve", "", -1)

	ss := strings.Split(s, ",")
	if len(ss) >= 2 {
		return ss[1]
	}

	return s
}
