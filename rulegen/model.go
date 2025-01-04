package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

var (
	NoResolve        = ",no-resolve"
	PreMatching      = ",pre-matching"
	ExtendedMatching = ",extended-matching"
)

type Rule struct {
	Type  string
	Value string
}

func ruleSort(rules []Rule) {
	sort.Slice(rules, func(i, j int) bool {
		// 先比较 Type 字段
		if rules[i].Type == rules[j].Type {
			// 如果 Type 相同，再比较 Value 字段
			return rules[i].Value < rules[j].Value
		}
		// 否则按照 Type 字段排序
		return rules[i].Type < rules[j].Type
	})
}

type PlatformGen interface {
	gen(rules []Rule, policy string) ([]string, error)
}

type AdGuardHome struct{}

func (g *AdGuardHome) gen(rules []Rule, policy string) ([]string, error) {
	return []string{"AdGuardHome config generated"}, nil
}

type Clash struct{}

func (g *Clash) gen(rules []Rule, policy string) ([]string, error) {
	return []string{"Clash config generated"}, nil
}

type DomainSet struct{}

func (g *DomainSet) gen(rules []Rule, policy string) ([]string, error) {
	var config []string
	for _, r := range rules {
		switch r.Type {
		case "HOST-WILDCARD", "IP-ASN", "IP-CIDR", "IP-CIDR6", "DOMAIN-KEYWORD":
			continue
		case "DOMAIN":
			config = append(config, r.Value)
		case "DOMAIN-SUFFIX":
			var s strings.Builder
			s.WriteString(".")
			s.WriteString(r.Value)
			config = append(config, s.String())
		default:
			return nil, fmt.Errorf("domain unknown type: %s", r.Type)
		}
	}
	return config, nil
}

type Host struct{}

func (g *Host) gen(rules []Rule, policy string) ([]string, error) {
	return []string{"Host config generated"}, nil
}

type Loon struct{}

func (g *Loon) gen(rules []Rule, policy string) ([]string, error) {
	var config []string
	for _, r := range rules {
		var s strings.Builder
		s.WriteString(r.Type)
		s.WriteString(",")
		s.WriteString(r.Value)
		switch r.Type {
		case "HOST-WILDCARD":
			continue
		case "IP-ASN", "IP-CIDR", "IP-CIDR6":
			if policy != "Direct" && policy != "ChinaASN" {
				s.WriteString(NoResolve)
			}
			config = append(config, s.String())
		case "DOMAIN-SUFFIX", "DOMAIN-KEYWORD", "DOMAIN", "USER-AGENT":
			config = append(config, s.String())
		default:
			return nil, fmt.Errorf("surge unknown type: %s", r.Type)
		}
	}
	return config, nil
}

type QuantumultX struct{}

func (g *QuantumultX) gen(rules []Rule, policy string) ([]string, error) {
	return []string{"QuantumultX config generated"}, nil
}

type Shadowrocket struct{}

func (g *Shadowrocket) gen(rules []Rule, policy string) ([]string, error) {
	return []string{"Shadowrocket config generated"}, nil
}

type SingBoxConfig struct {
	Version int           `json:"version"`
	Rules   []SingBoxRule `json:"rules"`
}
type SingBoxRule struct {
	QueryType        []interface{} `json:"query_type,omitempty"`
	Network          []string      `json:"network,omitempty"`
	Domain           []string      `json:"domain,omitempty"`
	DomainSuffix     []string      `json:"domain_suffix,omitempty"`
	DomainKeyword    []string      `json:"domain_keyword,omitempty"`
	DomainRegex      []string      `json:"domain_regex,omitempty"`
	SourceIPCIDR     []string      `json:"source_ip_cidr,omitempty"`
	IPCIDR           []string      `json:"ip_cidr,omitempty"`
	SourcePort       []int         `json:"source_port,omitempty"`
	SourcePortRange  []string      `json:"source_port_range,omitempty"`
	Port             []int         `json:"port,omitempty"`
	PortRange        []string      `json:"port_range,omitempty"`
	ProcessName      []string      `json:"process_name,omitempty"`
	ProcessPath      []string      `json:"process_path,omitempty"`
	ProcessPathRegex []string      `json:"process_path_regex,omitempty"`
	PackageName      []string      `json:"package_name,omitempty"`
	// NetworkType          []string      `json:"network_type,omitempty"`
	// NetworkIsExpensive   bool     `json:"network_is_expensive,omitempty"`
	// NetworkIsConstrained bool     `json:"network_is_constrained,omitempty"`
	WifiSSID  []string `json:"wifi_ssid,omitempty"`
	WifiBSSID []string `json:"wifi_bssid,omitempty"`
	Invert    bool     `json:"invert,omitempty"`
}

type SingBox struct{}

func (g *SingBox) gen(rules []Rule, policy string) ([]string, error) {
	var config SingBoxConfig
	var rr SingBoxRule
	config.Version = 1

	for _, r := range rules {
		switch r.Type {
		case "HOST-WILDCARD":
			continue
		case "IP-ASN", "IP-CIDR6":
			continue
		case "IP-CIDR":
			rr.IPCIDR = append(rr.IPCIDR, r.Value)
		case "DOMAIN-SUFFIX":
			rr.DomainSuffix = append(rr.DomainSuffix, r.Value)
		case "DOMAIN-KEYWORD":
			rr.DomainKeyword = append(rr.DomainKeyword, r.Value)
		case "DOMAIN":
			rr.Domain = append(rr.Domain, r.Value)
		case "USER-AGENT":
			continue
		default:
			return nil, fmt.Errorf("surge unknown type: %s", r.Type)
		}
	}
	config.Rules = append(config.Rules, rr)
	// 序列化为 JSON
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return []string{string(data)}, nil
}

type SurgeGen struct{}

func (g *SurgeGen) gen(rules []Rule, policy string) ([]string, error) {
	var config []string
	for _, r := range rules {
		var s strings.Builder
		s.WriteString(r.Type)
		s.WriteString(",")
		s.WriteString(r.Value)
		switch r.Type {
		case "HOST-WILDCARD":
			continue
		case "IP-ASN", "IP-CIDR", "IP-CIDR6":
			if policy != "Direct" && policy != "ChinaASN" {
				s.WriteString(NoResolve)
			}
			config = append(config, s.String())
		case "DOMAIN-SUFFIX", "DOMAIN-KEYWORD", "DOMAIN":
			s.WriteString(ExtendedMatching)
			config = append(config, s.String())
		case "USER-AGENT":
			config = append(config, s.String())
		default:
			return nil, fmt.Errorf("surge unknown type: %s", r.Type)
		}
	}
	return config, nil
}
