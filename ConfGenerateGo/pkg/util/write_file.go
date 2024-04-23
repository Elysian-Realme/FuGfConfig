package util

import (
	"ConfGenerateGo/pkg/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteFile(matchType string, data model.Pairs, policyName string, filePath string, writeType bool) error {
	var flag int
	if writeType {
		//	如果 为 true 就覆盖写入
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	} else {
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	}
	file, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		fmt.Println("open file error !")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	switch {
	case strings.Contains(matchType, "LoonRule"):
		fmt.Println("LoonRule")
		for _, v := range data {
			if strings.Contains(v.Value, "HOST-WILDCARD") {
				continue
			}
			switch {
			case strings.Contains(v.Value, "IP-ASN"):
				fmt.Fprintln(write, "IP-ASN,"+v.Key+",no-resolve")
				continue
			case strings.Contains(v.Value, "IP-CIDR6"):
				fmt.Fprint(write, "IP-CIDR6,")
			case strings.Contains(v.Value, "IP-CIDR"):
				fmt.Fprint(write, "IP-CIDR,")
				fmt.Fprint(write, v.Key)
				if !strings.Contains(v.Key, "/") {
					fmt.Fprint(write, "/32")
				}
				fmt.Fprintln(write, ",no-resolve")
				continue
			case strings.Contains(v.Value, "DOMAIN-SUFFIX"):
				fmt.Fprint(write, "DOMAIN-SUFFIX,")
			case strings.Contains(v.Value, "DOMAIN-KEYWORD"):
				fmt.Fprint(write, "DOMAIN-KEYWORD,")
			case strings.Contains(v.Value, "DOMAIN"):
				fmt.Fprint(write, "DOMAIN,")
			default:
				fmt.Fprint(write, v.Value+",")
			}
			fmt.Fprintln(write, v.Key)
		}
	case strings.Contains(matchType, "LoonHost"):
		fmt.Println("LoonHost")
		fmt.Fprintln(write, "[Host]")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") {
				v.Key = strings.Replace(v.Key, "\r", "", -1)
				v.Key = strings.Replace(v.Key, "\n", "", -1)
				if strings.Contains(v.Value, "DOMAIN-SUFFIX") {
					fmt.Fprint(write, "*")
				}
				fmt.Fprint(write, v.Key)
				fmt.Fprintln(write, " = 0.0.0.0")
			}
		}
	case strings.Contains(matchType, "QuantumultXHost"):
		fmt.Println("QuantumultXHost")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") {
				v.Key = strings.Replace(v.Key, "\r", "", -1)
				v.Key = strings.Replace(v.Key, "\n", "", -1)
				fmt.Fprint(write, "server=/")
				if strings.Contains(v.Value, "DOMAIN-SUFFIX") {
					fmt.Fprint(write, "*")
				}
				fmt.Fprint(write, v.Key)
				fmt.Fprintln(write, "/0.0.0.0")
			}
		}
	case strings.Contains(matchType, "QuantumultXRules"):
		fmt.Println("QuantumultXRules")
		for _, v := range data {
			switch {
			case strings.Contains(v.Value, "IP-ASN"):
				fmt.Fprint(write, "IP-ASN,"+v.Key)
				fmt.Fprintln(write, ","+policyName+",no-resolve")
				continue
			case strings.Contains(v.Value, "IP-CIDR6"):
				fmt.Fprint(write, "IP6-CIDR,")
			case strings.Contains(v.Value, "IP-CIDR"):
				fmt.Fprint(write, "IP-CIDR,"+v.Key)
				if !strings.Contains(v.Key, "/") {
					fmt.Fprint(write, "/32")
				}
				fmt.Fprintln(write, ","+policyName+",no-resolve")
				continue
			case strings.Contains(v.Value, "DOMAIN-SUFFIX"):
				fmt.Fprint(write, "HOST-SUFFIX,")
			case strings.Contains(v.Value, "DOMAIN"):
				fmt.Fprint(write, "HOST,")
			case strings.Contains(v.Value, "DOMAIN-KEYWORD"):
				fmt.Fprint(write, "HOST-KEYWORD,")
			case strings.Contains(v.Value, "HOST-WILDCARD"):
				fmt.Fprint(write, "HOST-WILDCARD,")
			default:
				fmt.Fprint(write, v.Value+",")
			}

			fmt.Fprintln(write, v.Key+","+policyName)
		}
	case strings.Contains(matchType, "SurgeRule"):
		fmt.Println("SurgeRule")
		for _, v := range data {
			if strings.Contains(v.Value, "HOST-WILDCARD") {
				continue
			}
			switch {
			case strings.Contains(v.Value, "IP-ASN"):
				fmt.Fprintln(write, "IP-ASN,"+v.Key+",no-resolve")
				continue
			case strings.Contains(v.Value, "IP-CIDR6"):
				fmt.Fprint(write, "IP-CIDR6,")
			case strings.Contains(v.Value, "IP-CIDR"):
				fmt.Fprint(write, "IP-CIDR,")
				fmt.Fprint(write, v.Key)
				if !strings.Contains(v.Key, "/") {
					fmt.Fprint(write, "/32")
				}
				fmt.Fprintln(write, ",no-resolve")
				continue
			case strings.Contains(v.Value, "DOMAIN-SUFFIX"):
				fmt.Fprint(write, "DOMAIN-SUFFIX,")
			case strings.Contains(v.Value, "DOMAIN-KEYWORD"):
				fmt.Fprint(write, "DOMAIN-KEYWORD,")
			case strings.Contains(v.Value, "DOMAIN"):
				fmt.Fprint(write, "DOMAIN,")
			default:
				fmt.Fprint(write, v.Value+",")
			}
			fmt.Fprintln(write, v.Key)
		}
	case strings.Contains(matchType, "SurgeHost"):
		fmt.Println("SurgeHost")
		fmt.Fprintln(write, "#!name="+policyName)
		fmt.Fprintln(write, "#!desc="+policyName)
		fmt.Fprintln(write, "#!category="+"FuckGFW")
		fmt.Fprintln(write, "[Host]")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") {
				v.Key = strings.Replace(v.Key, "\r", "", -1)
				v.Key = strings.Replace(v.Key, "\n", "", -1)
				if strings.Contains(v.Value, "DOMAIN-SUFFIX") {
					fmt.Fprint(write, "*.")
				}
				fmt.Fprint(write, v.Key)
				fmt.Fprintln(write, " = 0.0.0.0")
			}
		}
	case strings.Contains(matchType, "Host"):
		fmt.Println("Host")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") || strings.Contains(v.Value, "IP-CIDR") || strings.Contains(v.Value, "IP-CIDR6") {
				if strings.Contains(v.Key, "/") {
					continue
				}
				fmt.Fprint(write, "0.0.0.0 ")
				fmt.Fprintln(write, v.Key)
			}
		}
	case strings.Contains(matchType, "DomainSetRule"):
		fmt.Println("DomainSetRule")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") {
				v.Key = strings.Replace(v.Key, "\r", "", -1)
				v.Key = strings.Replace(v.Key, "\n", "", -1)
				if strings.Contains(v.Value, "DOMAIN-SUFFIX") {
					fmt.Fprint(write, ".")
				}
				fmt.Fprintln(write, v.Key)
			}
		}
	case strings.Contains(matchType, "AdGuardHome"):
		fmt.Println("AdGuardHome")
		for _, v := range data {
			if strings.Contains(v.Value, "DOMAIN-SUFFIX") || strings.Contains(v.Value, "DOMAIN") || strings.Contains(v.Value, "HOST-WILDCARD") {
				v.Key = strings.Replace(v.Key, "\r", "", -1)
				v.Key = strings.Replace(v.Key, "\n", "", -1)
				fmt.Fprint(write, "||")
				fmt.Fprint(write, v.Key)
				fmt.Fprintln(write, "^")
			}
			if strings.Contains(v.Value, "IP-CIDR") || strings.Contains(v.Value, "IP-CIDR6") {
				if strings.Contains(v.Key, "/") {
					continue
				}
				fmt.Fprint(write, "0.0.0.0 ")
				fmt.Fprintln(write, v.Key)
			}
		}
	case strings.Contains(matchType, "ShadowrocketRule"):
		fmt.Println("ShadowrocketRule")

		fmt.Fprintln(write, "#!name="+policyName)
		fmt.Fprintln(write, "#!desc="+policyName)
		fmt.Fprintln(write, "[Rule]")
		for _, v := range data {
			if strings.Contains(v.Value, "HOST-WILDCARD") {
				continue
			}
			switch {
			case strings.Contains(v.Value, "IP-CIDR6"):
				fmt.Fprint(write, "IP-CIDR6,")
			case strings.Contains(v.Value, "IP-CIDR"):
				fmt.Fprint(write, "IP-CIDR,")
				fmt.Fprint(write, v.Key)
				if !strings.Contains(v.Key, "/") {
					fmt.Fprint(write, "/32")
				}
				fmt.Fprintln(write, ",no-resolve,"+policyName)
				continue
			case strings.Contains(v.Value, "DOMAIN-SUFFIX"):
				fmt.Fprint(write, "DOMAIN-SUFFIX,")
			case strings.Contains(v.Value, "DOMAIN-KEYWORD"):
				fmt.Fprint(write, "DOMAIN-KEYWORD,")
			case strings.Contains(v.Value, "DOMAIN"):
				fmt.Fprint(write, "DOMAIN,")
			}
			fmt.Fprintln(write, v.Key+","+policyName)
		}
	case strings.Contains(matchType, "Clash"):
		// todo clash 是否支持 USER-AGENT
		fmt.Fprintln(write, "payload:")
		for _, v := range data {
			if strings.Contains(v.Value, "HOST-WILDCARD") || strings.Contains(v.Value, "USER-AGENT") {
				continue
			}
			v.Key = strings.Replace(v.Key, "\r", "", -1)
			v.Key = strings.Replace(v.Key, "\n", "", -1)

			fmt.Fprint(write, "  - ")

			switch {
			case strings.Contains(v.Value, "IP-ASN"):
				fmt.Fprintln(write, "IP-ASN,"+v.Key+",no-resolve")
				continue
			case strings.Contains(v.Value, "IP-CIDR6"):
				fmt.Fprint(write, "IP-CIDR6,")
			case strings.Contains(v.Value, "IP-CIDR"):
				fmt.Fprint(write, "IP-CIDR,")
				fmt.Fprint(write, v.Key)
				if !strings.Contains(v.Key, "/") {
					fmt.Fprint(write, "/32")
				}
				fmt.Fprintln(write, ",no-resolve")
				continue
			case strings.Contains(v.Value, "DOMAIN-SUFFIX"):
				fmt.Fprint(write, "DOMAIN-SUFFIX,")
			case strings.Contains(v.Value, "DOMAIN-KEYWORD"):
				fmt.Fprint(write, "DOMAIN-KEYWORD,")
			case strings.Contains(v.Value, "DOMAIN"):
				fmt.Fprint(write, "DOMAIN,")
			}
			fmt.Fprintln(write, v.Key)
		}
	default:
		fmt.Println("匹配 error")
	}

	return write.Flush()
}

func NormalWriteFile(data []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file error !")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for _, v := range data {
		fmt.Fprintln(write, v)
	}

	return write.Flush()
}

func QuantumultXMITMWriteFile(data []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file error !")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	fmt.Fprint(write, "hostname = %APPEND% ")
	for i, v := range data {
		fmt.Fprint(write, v)
		if i < len(data)-1 {
			fmt.Fprint(write, ", ")
		}
	}
	fmt.Fprintln(write, "")
	fmt.Fprintln(write, "")
	fmt.Fprintln(write, `^http:\/\/182\.256\.116\.116\/d - reject`)

	return write.Flush()
}

func LoonMITMWriteFile(data []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file error !")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	fmt.Fprintln(write, "[MITM]")
	fmt.Fprint(write, "hostname = %APPEND% ")
	for i, v := range data {
		fmt.Fprint(write, v)
		if i < len(data)-1 {
			fmt.Fprint(write, ", ")
		}
	}
	fmt.Fprintln(write, "")
	// fmt.Fprintln(write, "")
	// fmt.Fprintln(write, `^http:\/\/182\.256\.116\.116\/d - reject`)

	return write.Flush()
}
