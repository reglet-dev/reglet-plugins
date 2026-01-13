package main

import (
	"context"
	"time"

	regletsdk "github.com/reglet-dev/reglet-sdk/go"
)

// MXRecord represents an MX DNS record.
type MXRecord struct {
	Host string
	Pref uint16
}

// DNSLookupResult represents the result of a DNS lookup.
// This is a local type for testability.
type DNSLookupResult struct {
	Records   []string
	MXRecords []MXRecord
	Error     *DNSError
}

// DNSError represents an error from DNS lookup.
type DNSError struct {
	Message    string
	Type       string
	IsTimeout  bool
	IsNotFound bool
}

func (e *DNSError) Error() string {
	return e.Message
}

// DNSLookupFunc is the function signature for DNS lookups.
type DNSLookupFunc func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error)

// dnsPlugin implements the sdk.Plugin interface.
type dnsPlugin struct {
	// LookupFunc allows dependency injection for testing
	LookupFunc DNSLookupFunc
}

// Describe returns plugin metadata.
func (p *dnsPlugin) Describe(ctx context.Context) (regletsdk.Metadata, error) {
	return regletsdk.Metadata{
		Name:        "dns",
		Version:     "1.0.0",
		Description: "DNS resolution and record validation",
		Capabilities: []regletsdk.Capability{
			{
				Kind:    "network",
				Pattern: "outbound:53", // Required for DNS lookups
			},
		},
	}, nil
}

type DNSConfig struct {
	Hostname   string `json:"hostname" validate:"required" description:"Hostname to resolve"`
	RecordType string `json:"record_type" validate:"oneof=A AAAA CNAME MX TXT NS" default:"A" description:"DNS record type to query"`
	Nameserver string `json:"nameserver,omitempty" description:"Custom nameserver (optional, e.g., 8.8.8.8:53)"`
}

// Schema returns the JSON schema for the plugin's configuration.
func (p *dnsPlugin) Schema(ctx context.Context) ([]byte, error) {
	return regletsdk.GenerateSchema(DNSConfig{})
}

// Check executes the DNS observation.
func (p *dnsPlugin) Check(ctx context.Context, config regletsdk.Config) (regletsdk.Evidence, error) {
	// Set default record type if not provided
	if _, ok := config["record_type"]; !ok {
		config["record_type"] = "A"
	}

	var cfg DNSConfig
	if err := regletsdk.ValidateConfig(config, &cfg); err != nil {
		return regletsdk.Evidence{
			Status: false,
			Error:  regletsdk.ToErrorDetail(&regletsdk.ConfigError{Err: err}),
		}, nil
	}

	if p.LookupFunc == nil {
		return regletsdk.Failure("internal", "LookupFunc not initialized"), nil
	}

	start := time.Now()
	result, err := p.LookupFunc(ctx, cfg.Hostname, cfg.RecordType, cfg.Nameserver)
	queryTime := time.Since(start).Milliseconds()

	// Prepare data for evidence.
	data := map[string]interface{}{
		"hostname":      cfg.Hostname,
		"record_type":   cfg.RecordType,
		"query_time_ms": queryTime,
	}

	var evidence regletsdk.Evidence
	var dnsErr *DNSError

	if err != nil {
		// Go error from lookup
		dnsErr = &DNSError{
			Message: err.Error(),
			Type:    "internal",
		}
		evidence = regletsdk.Failure("dns_lookup_error", err.Error())
	} else if result.Error != nil {
		// DNS-specific error
		dnsErr = result.Error
		if dnsErr.Type == "config" {
			evidence = regletsdk.Evidence{
				Status: false,
				Error:  regletsdk.ToErrorDetail(&regletsdk.ConfigError{Err: dnsErr}),
			}
		} else {
			evidence = regletsdk.Evidence{
				Status: false,
				Error: regletsdk.ToErrorDetail(&regletsdk.NetworkError{
					Operation: "dns_lookup",
					Target:    cfg.Hostname,
					Err:       dnsErr,
				}),
			}
		}
	} else {
		// Success path
		recordCount := 0
		if result.Records != nil {
			data["records"] = result.Records
			recordCount = len(result.Records)
		}
		if result.MXRecords != nil {
			var mxRecords []map[string]interface{}
			for _, mx := range result.MXRecords {
				mxRecords = append(mxRecords, map[string]interface{}{"host": mx.Host, "pref": mx.Pref})
			}
			data["mx_records"] = mxRecords
			recordCount = len(mxRecords)
		}
		data["record_count"] = recordCount
		evidence = regletsdk.Success(data)
	}

	// Always populate error flags
	if dnsErr != nil {
		data["error_message"] = dnsErr.Message
		data["is_timeout"] = dnsErr.IsTimeout
		data["is_not_found"] = dnsErr.IsNotFound
	} else {
		data["is_timeout"] = false
		data["is_not_found"] = false
	}
	evidence.Data = data

	return evidence, nil
}
