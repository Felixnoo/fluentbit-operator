package parser

import (
	"fmt"

	"fluent.io/fluent-operator/apis/fluentbit/v1alpha2/plugins"
	"fluent.io/fluent-operator/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// The LSTV parser plugin
type LSTV struct {
	// Time_Key
	TimeKey string `json:"timeKey,omitempty"`
	// Time_Format, eg. %Y-%m-%dT%H:%M:%S %z
	TimeFormat string `json:"timeFormat,omitempty"`
	// Time_Keep
	TimeKeep *bool  `json:"timeKeep,omitempty"`
	Types    string `json:"types,omitempty"`
}

func (_ *LSTV) Name() string {
	return "ltsv"
}

func (l *LSTV) Params(_ plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if l.TimeKey != "" {
		kvs.Insert("Time_Key", l.TimeKey)
	}
	if l.TimeFormat != "" {
		kvs.Insert("Time_Format", l.TimeFormat)
	}
	if l.TimeKeep != nil {
		kvs.Insert("Time_Format", fmt.Sprint(*l.TimeKeep))
	}
	if l.Types != "" {
		kvs.Insert("Types", l.Types)
	}
	return kvs, nil
}
