package slo

import (
	"fmt"

	oslov1 "github.com/alexandreLamarre/oslo/pkg/manifest/v1"
	"github.com/kralicky/yaml/v3"
	prommodel "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	apis "github.com/rancher/opni/plugins/slo/pkg/apis/slo"
)

var (
	RecordingRuleSuffix = "-recording"
	MetadataRuleSuffix  = "-metadata"
	AlertRuleSuffix     = "-alerts"
)

// these types are defined to support yaml v2 (instead of the new Prometheus
// YAML v3 that has some problems with marshaling).
type ruleGroupsYAMLv2 struct {
	Groups []ruleGroupYAMLv2 `yaml:"groups"`
}

type ruleGroupYAMLv2 struct {
	Name     string             `yaml:"name"`
	Interval prommodel.Duration `yaml:"interval,omitempty"`
	Rules    []rulefmt.Rule     `yaml:"rules"`
}

// Guess this could be generic
type zipperHolder struct {
	Spec    *oslov1.SLO
	Service *apis.Service
}

type CortexRuleWrapper struct {
	recording string
	metadata  string
	alerts    string
}

func zipOpenSLOWithServices(ps []oslov1.SLO, as []*apis.Service) ([]*zipperHolder, error) {
	if len(as) != len(ps) {
		return nil, fmt.Errorf("Expected Generated SLOGroups to match the number of Services provided in the request")
	}
	res := make([]*zipperHolder, 0)
	for idx, p := range ps {
		res = append(res, &zipperHolder{
			Spec:    &p,
			Service: as[idx],
		})
	}
	return res, nil
}

// Marshal result SLORuleFmtWrapper to cortex-approved yaml
func toCortexRequest(rw SLORuleFmtWrapper, sloId string) (*CortexRuleWrapper, error) {

	recording, metadata, alerts := rw.SLIrules, rw.MetaRules, rw.AlertRules
	// Check length is 0?
	rrecording, err := yaml.Marshal(ruleGroupYAMLv2{
		Name:  fmt.Sprintf("%s%s", sloId, RecordingRuleSuffix),
		Rules: recording,
	})
	if err != nil {
		return nil, err
	}

	rmetadata, err := yaml.Marshal(ruleGroupYAMLv2{
		Name:  fmt.Sprintf("%s%s", sloId, MetadataRuleSuffix),
		Rules: metadata,
	})
	if err != nil {
		return nil, err
	}

	ralerts, err := yaml.Marshal(ruleGroupYAMLv2{
		Name:  fmt.Sprintf("%s%s", sloId, AlertRuleSuffix),
		Rules: alerts,
	})
	if err != nil {
		return nil, err
	}
	return &CortexRuleWrapper{
		recording: string(rrecording),
		metadata:  string(rmetadata),
		alerts:    string(ralerts),
	}, nil
}
