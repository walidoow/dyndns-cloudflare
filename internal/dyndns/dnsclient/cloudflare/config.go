package cloudflare

import (
	"dyndns-cloudflare/pkg/config"
	"encoding/json"
)

type DNSRecordConfig struct {
	Token          string `yaml:"token"`
	ZoneIdentifier string `yaml:"zone-identifier"`
	DNSRecord      struct {
		Identifier string   `yaml:"identifier"`
		Name       string   `yaml:"name"`
		Proxied    bool     `yaml:"proxied"`
		Type       string   `yaml:"type"`
		Comment    string   `yaml:"comment"`
		Tag        []string `yaml:"tag"`
		Ttl        int      `yaml:"ttl"`
	} `yaml:"cloudflare-dns-record"`
}

func (receiver DNSRecordConfig) Payload(content string) ([]byte, error) {
	record := DNSRecordPayload{Content: content, Name: receiver.DNSRecord.Name, Proxied: receiver.DNSRecord.Proxied,
		Type: receiver.DNSRecord.Type, Comment: receiver.DNSRecord.Comment, Tags: receiver.DNSRecord.Tag,
		Ttl: receiver.DNSRecord.Ttl}
	return json.Marshal(record)
}

func LoadYAMLConfig(configPath string) DNSRecordConfig {
	recordConfig := DNSRecordConfig{}
	config.DecodeYAMLConfig(config.ReadConfigFile(configPath), &recordConfig)
	return recordConfig
}
