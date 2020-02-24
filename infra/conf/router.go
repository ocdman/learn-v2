package conf

import (
	"encoding/json"
)

type BalancingRule struct {
	Tag string `json:"tag"`
	Selectors StringList `json:"selector`
}

type RouterConfig struct {
	RuleList []json.RawMessage `json:"rules"`
	DomainStrategy *string `json:"domainStrategy"`
	Balancers []*BalancingRule `json:"balancers"`
}