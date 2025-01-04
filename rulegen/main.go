package main

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
)

func main() {
	list := []string{"Surge", "Loon", "sing-box"}
	// 遍历目录下的文件
	dir := "../ConfigFile/DataFile/"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && !strings.Contains(info.Name(), ".DS_Store") {
			name := strings.ToLower(info.Name())
			var ipFile, domainFile, ipInboxFile, domainInboxFile []string
			switch {
			case strings.Contains(name, "ip"):
				ipFile = readFile(path)
				// logger.Info("Read File: " + name)
			case strings.Contains(name, "domain"):
				domainFile = readFile(path)
				// logger.Info("Read File: " + name)
			case strings.Contains(name, "inbox-ip"):
				ipInboxFile = readFile(path)
				// logger.Info("Read File: " + name)
			case strings.Contains(name, "inbox-domain"):
				domainInboxFile = readFile(path)
			// logger.Info("Read File: " + name)
			default:
				logger.Warn("出现不规范文件" + path)
			}
			// 生成各个软件对应的规则
			rules := make([][]Rule, 4)
			var im, dm map[string]string
			rules[0], im = ipGen(ipFile)
			rules[1], dm = domainGen(domainFile)
			rules[2] = ipInboxGen(ipInboxFile, im)
			rules[3] = domainInboxGen(domainInboxFile, dm)
			// 写入文件
			// 拼接写入文件路径
			p := strings.Split(path, "/")
			baseDir := "../ConfigFile"
			for i := 0; i < 4; i++ {
				if len(rules[i]) <= 0 {
					continue
				}
				for _, l := range list {
					policy := p[len(p)-2]
					config, err := gen(rules[i], l, policy)
					if err != nil {
						logger.Error(err.Error())
					}
					p := filepath.Join(baseDir, l, policy, strings.Split(name, ".")[0]+".list")
					// logger.Info(p)

					err = writeFile(p, config)
					if err != nil {
						logger.Error(err.Error())
					}
				}

			}
		}

		return nil
	})
	if err != nil {
		logger.Error(err.Error())
	}
}

// ipGen 读取，生成并排序现有的 IP 规则
func ipGen(lines []string) ([]Rule, map[string]string) {
	var rules []Rule
	ans := make(map[string]string)
	for _, line := range lines {
		line = formatCorrection(line)
		switch {
		case isASN(line):
			ans[line] = "IP-ASN"
		case isIPv4(line):
			if !strings.Contains(line, "/") {
				line += "/32"
			}
			ans[line] = "IP-CIDR"
		case isIPv6(line):
			if !strings.Contains(line, "/") {
				line += "/128"
			}
			ans[line] = "IP-CIDR6"
		case isDomain(line):
			logger.Warn("IP 规则中出现了 Domain 规则：" + line)
		default:
			logger.Warn("IP 匹配 Case 出现漏网之🐟:" + line)
		}
	}

	for k, v := range ans {
		rules = append(rules, Rule{Type: v, Value: k})
	}

	ruleSort(rules)

	return rules, ans
}

func domainGen(lines []string) ([]Rule, map[string]string) {
	var rules []Rule
	ans := make(map[string]string)
	for _, line := range lines {
		line = formatCorrection(line)
		switch {
		case strings.HasPrefix(line, "USER-AGENT"):
			ans[strings.Split(line, ",")[1]] = "USER-AGENT"
		case isASN(line) || isIPv4(line) || isIPv6(line):
			logger.Warn("Domain 规则中出现了 IP 规则：" + line)
		case isDomain(line) || line == ".cn":
			switch {
			case strings.HasPrefix(line, ".") || strings.HasPrefix(line, "*."):
				line = strings.TrimPrefix(line, ".")
				line = strings.TrimPrefix(line, "*.")
				ans[line] = "DOMAIN-SUFFIX"
			case line[0] == '*' && strings.Count(line, string('*')) == 1:
				ans[line[1:]] = "DOMAIN-KEYWORD"
			case strings.Contains(line, "*"):
				line = strings.TrimPrefix(line, "*.")
				ans[line] = "HOST-WILDCARD"
			default:
				ans[line] = "DOMAIN"
			}

		default:
			logger.Warn("Domain 匹配 Case 出现漏网之🐟:" + line)
		}
	}

	for k, v := range ans {
		rules = append(rules, Rule{Type: v, Value: k})
	}

	ruleSort(rules)
	return rules, ans
}

func ipInboxGen(lines []string, base map[string]string) []Rule {
	var rules []Rule
	for _, line := range lines {
		if _, ok := base[line]; ok {
			continue
		} else {
			base[line] = "IP"
		}
	}

	return rules
}

func domainInboxGen(lines []string, base map[string]string) []Rule {
	var rules []Rule

	// for _, line := range lines {
	// 	switch {
	// 	case isDomain(line):
	//
	// 	default:
	// 		logger.Warn("Inbox 匹配 Case 出现漏网之🐟:" + line)
	// 	}
	// }

	return rules
}

func gen(rules []Rule, platform string, policy string) ([]string, error) {
	var g PlatformGen

	// 根据平台选择不同的处理器
	switch platform {
	// case "AdGuardHome":
	// 	g = &AdGuardHome{}
	// case "DomainSet":
	// 	g = &DomainSet{}
	case "Loon":
		g = &Loon{}
	case "sing-box":
		g = &SingBox{}
	case "Surge":
		g = &SurgeGen{}
	default:
		return nil, errors.New("Gen 匹配 Case 出现漏网之🐟: " + platform)
	}

	config, err := g.gen(rules, policy)
	if err != nil {
		return nil, err
	}
	return config, nil
}
