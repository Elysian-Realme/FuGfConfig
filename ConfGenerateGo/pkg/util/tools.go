package util

import (
	"fmt"
	"os"
	"regexp"
)

// IsIPV4 判断是否为 IPV4
func IsIPV4(s string) bool {
	ipv4Pattern := regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}(\/\d{1,2})?$`)
	// ipv4Pattern := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	// `^(\d{1,3}\.){3}\d{1,3}(\/\d{1,2})?$` 匹配 192.168.1.1 和 192.168.1.1/24
	return ipv4Pattern.MatchString(s)
}

// func IsIPV4(s string) bool {
// 	ip := net.ParseIP(s)
// 	if ip == nil {
// 		fmt.Println("IP address is not valid")
// 		return false
// 	}

// 	if ip.To4() != nil {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// IsIPV6 判断是否为 IPV6
func IsIPV6(s string) bool {
	ipv6Pattern := regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`)
	return ipv6Pattern.MatchString(s)
}

// func IsIPV6(s string) bool {
// 	ip := net.ParseIP(s)
// 	if ip == nil {
// 		fmt.Println("IP address is not valid")
// 		return false
// 	}

// 	if ip.To16() != nil {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// IsDomainRule 是否为 域名
func IsDomainRule(s string) bool {
	// todo 会匹配 128.1.102.20, 203.205.179.215, 39.107.142.158
	//domainPattern := regexp.MustCompile(`^.?[^\s/$.?#].\S*$`)
	// todo 无法匹配有下划线或者*的 domain；会匹配 128.1.102.20, 203.205.179.215, 39.107.142.158
	domainPattern := regexp.MustCompile(`^[a-zA-Z0-9_*\-.]+\.[a-zA-Z0-9_*\-.]+$`)
	return domainPattern.MatchString(s)
}

// IsASN 是否为 ASN 规则
func IsASN(s string) bool {
	domainPattern := regexp.MustCompile(`^[0-9]+$`)
	return domainPattern.MatchString(s)
}

// IsFileExist 判断文件是否存在
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(filePath, " 文件不存在")
		} else {
			fmt.Println("发生错误:", err)
		}
		return false
	} else {
		return true
	}
}

// SliceReverse 翻转
func SliceReverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
