package main

import (
	"ConfGenerateGo/pkg/model"
	"ConfGenerateGo/pkg/util"
	"fmt"
	"os"
	"sort"
	"strings"
)

// loon data file path
// var loonInboxRulesUrl = [...]string{
// 	"https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list",
// 	"https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"}

func main() {
	println("开始")

	// fmt.Println("是否要更新 or 下载远程数据 (y or n)")
	// var input string
	// // fmt.Scanln(&input)
	// input = "y"
	// // input = "n"
	// if input == "y" || input == "Y" {
	// 	downloadFiles()
	// }

	var base, inbox, inboxResult []string
	var ans []model.Pair

	// names := []string{"FuckRogueSoftware"}
	names := []string{"AI", "Bank", "Direct", "Cryptocurrency", "Telegram", "Proxy", "CodeTools", "Microsoft", "Tracker", "FuckGarbageFeature", "FuckRogueSoftware", "Games"}
	for _, name := range names {
		//  清空残留的数据
		base, inbox, inboxResult = []string{}, []string{}, []string{}
		ans = []model.Pair{}
		fmt.Println("----开始处理 ", name, " ----")
		// 拼接文件路径
		buildString := func(ss ...string) string {
			var builder strings.Builder
			for _, s := range ss {
				builder.WriteString(s)
			}
			return builder.String()
		}
		base, inbox = readRule(buildString("../ConfigFile/DataFile/", name, "/", name, ".txt"), buildString("../ConfigFile/DataFile/", name, "/inbox.txt"))
		ans, inboxResult = policyProcessing(base, inbox)
		util.WriteFile("QuantumultXRules", ans, name, buildString("../ConfigFile/QuantumultX/", name, "Rules.conf"), true)
		util.WriteFile("LoonRule", ans, name, buildString("../ConfigFile/Loon/LoonRemoteRule/", name, "Rules.conf"), true)
		util.WriteFile("SurgeRule", ans, name, buildString("../ConfigFile/Surge/", name, "Rules.conf"), true)
		util.WriteFile("DomainSetRule", ans, name, buildString("../ConfigFile/DomainSet/", name, "Rules.conf"), true)
		util.WriteFile("Clash", ans, name, buildString("../ConfigFile/Clash/", name, "Rules.conf"), true)

		if name == "FuckRogueSoftware" || strings.EqualFold(name, "FuckGarbageFeature") {
			util.WriteFile("Host", ans, name, buildString("../ConfigFile/Host/", name, "Rules.conf"), true)
			util.WriteFile("AdGuardHome", ans, name, buildString("../ConfigFile/AdGuardHome/", name, "Rules.conf"), true)
			util.WriteFile("SurgeHost", ans, name, buildString("../ConfigFile/Surge/SurgeModule/", name, "Module.sgmodule"), true)
		}

		if name == "FuckRogueSoftware" {
			util.WriteFile("ShadowrocketRule", ans, "Reject", buildString("../ConfigFile/Shadowrocket/", name, "Rules.conf"), true)
		}

		if name == "CnAppRule" {
			data := util.ReadFile("../ConfigFile/DataFile/CnAppRule/CnAppRule.txt")
			util.QuantumultXMITMWriteFile(data, "../ConfigFile/QuantumultX/RewriteRemote/MITM.conf")
			util.LoonMITMWriteFile(data, "../ConfigFile/Loon/LoonPlugin/MITM.plugin")
		}

		if len(inboxResult) != 0 {
			util.NormalWriteFile(inboxResult, buildString("../ConfigFile/DataFile/", name, "/inbox.txt"))
		}
	}

	println("处理完成")
	println("结束")
}

func readRule(baseFilePath string, inboxFilePath string) ([]string, []string) {
	//
	var base, inbox []string
	// 读取 base
	// 判断文件是否存在
	_, err := os.Stat(baseFilePath)
	if err == nil {
		base = util.ReadFile(baseFilePath)
	} else {
		fmt.Println("发生错误:", err)
	}

	// 读取 inbox
	_, err = os.Stat(inboxFilePath)
	if err == nil {
		inbox = util.ReadFile(inboxFilePath)
	} else {
		fmt.Println(err)
	}

	return base, inbox
}

func policyProcessing(base []string, inbox []string) ([]model.Pair, []string) {
	// map 来存取数据 key 是唯一的 放置域名或者 ip，value 放置规则

	// 构建 base map
	var ansMap = make(map[string]string)
	for _, v := range base {
		v = util.FormatCorrection(v)
		a := ""
		if strings.Count(v, ",") >= 1 {
			a, v = splitRule(v)
			ansMap[v] = a
		} else if util.IsASN(v) {
			ansMap[v] = "IP-ASN"
		} else if util.IsIPV4(v) {
			ansMap[v] = "IP-CIDR"
		} else if util.IsIPV6(v) {
			ansMap[v] = "IP-CIDR6"
		} else if util.IsDomainRule(v) {
			if strings.HasPrefix(v, ".") {
				v = strings.TrimPrefix(v, ".")
				ansMap[v] = "DOMAIN-SUFFIX"
			} else if strings.HasPrefix(v, "*.") {
				v = strings.TrimPrefix(v, "*.")
				ansMap[v] = "DOMAIN-SUFFIX"
			} else if v[0] == '*' && strings.Count(v, string('*')) == 1 { // 如果第一个字符为 * 且 * 只出现了一次
				ansMap[v[1:]] = "DOMAIN-KEYWORD"
			} else if strings.Contains(v, "*") {
				ansMap[v] = "HOST-WILDCARD"
			} else {
				ansMap[v] = "DOMAIN"
			}
		} else {
			fmt.Println("发现未匹配到的规则，规则为：" + v)
			if strings.HasPrefix(v, ".") {
				v = strings.TrimPrefix(v, ".")
				ansMap[v] = "DOMAIN-SUFFIX"
			} else {
				ansMap[v] = "DOMAIN"
			}
		}
	}

	fmt.Println("规则基础库构建完成，共:", len(ansMap), "条规则")

	// 遍历 inbox
	var inboxResult []string
	type void struct{}
	var member void

	inboxResultSet := make(map[string]void)
	if len(inbox) > 0 {
		for _, v := range inbox {
			// 清除一下格式，只保留 IP 或者域名
			v = util.CleanAll(v)
			// 跳过 待处理 规则中 类似于 xxx.com 的规则 或者 xxx 的规则
			if strings.Count(v, ".") == 1 || strings.Count(v, ".") == 0 {
				continue
			}
			// 检查是否在 base map 里面
			if _, ok := ansMap[v]; !ok {
				// 如果不存在
				if util.IsDomainRule(v) {
					// 如果是 domain 规则，则检查后缀域名是否存在 base map 里面
					count := strings.Count(v, ".") - 1
					flag1, flag2 := false, false
					s := v
					for i := 0; i < count; i++ {
						s = domainRuleIntercept(s)
						if _, ok := ansMap[s]; ok {
							flag1 = true
							break
						}
						// 检查一下 v 的后缀域名是否在 inboxResultSet 中存在
						if _, ok := inboxResultSet[s]; ok {
							flag2 = true
							break
						}
					}
					// 如果所有后缀域名均不在 base map 且不在 inboxResultSet 里面 ，则加入到 inbox map
					if !flag1 && !flag2 {
						inboxResultSet[v] = member
					}
				} else if util.IsIPV4(v) {
					// todo 如果是 IPV4 规则，匹配一下 去掉 CIDR格式 前后的字符串是否在 base map 和 inbox map 中
					inboxResultSet[v] = member
				} else {
					inboxResultSet[v] = member
				}
			}
		}
	}

	// 将 inbox map 输出到 [] string 中
	for key := range inboxResultSet {
		inboxResult = append(inboxResult, key)
	}

	sort.Slice(inboxResult, func(i, j int) bool {
		a, b := inboxResult[i], inboxResult[j]
		for k := 1; k <= len(a) && k <= len(b); k++ {
			if a[len(a)-k] != b[len(b)-k] {
				return a[len(a)-k] < b[len(b)-k]
			}
		}
		return len(a) > len(b)
	})
	inboxResult = util.SliceReverse(inboxResult)

	fmt.Println("查重后未处理的规则还剩 ", len(inboxResult), " 条")

	var data model.Pairs
	for k, v := range ansMap {
		data = append(data, model.Pair{Key: k, Value: v})
	}
	sort.Sort(sort.Reverse(data))

	fmt.Println("排序处理完后的规则共: ", len(data), " 条")

	return data, inboxResult
}

// 对 规则进行切片，返回中间
func splitRule(s string) (string, string) {
	ss := strings.Split(s, ",")
	if len(ss) >= 2 {
		return ss[0], ss[1]
	}
	return "", ""
}

// 对域名规则按 "." 切片
func domainRuleIntercept(s string) string {
	firstInd := strings.Index(s, ".")
	return s[firstInd+1:]
}

// 错误处理函数
func handleError(fn func() error) {
	if err := fn(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
}
