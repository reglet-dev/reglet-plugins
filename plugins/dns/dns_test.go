//go:build !wasip1

package main

import (
	"context"
	"errors"
	"testing"

	regletsdk "github.com/reglet-dev/reglet-sdk/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDNSPlugin_Check_ConfigValidation(t *testing.T) {
	// Create a mock lookup that doesn't actually make network calls
	mockLookup := func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
		return &DNSLookupResult{
			Records: []string{"93.184.216.34"},
		}, nil
	}

	plugin := &dnsPlugin{LookupFunc: mockLookup}

	tests := []struct {
		name      string
		config    regletsdk.Config
		wantError bool
		errMsg    string
	}{
		{
			name: "Valid A record config",
			config: regletsdk.Config{
				"hostname":    "example.com",
				"record_type": "A",
			},
			wantError: false,
		},
		{
			name: "Valid MX record config",
			config: regletsdk.Config{
				"hostname":    "gmail.com",
				"record_type": "MX",
			},
			wantError: false,
		},
		{
			name: "Valid config with nameserver",
			config: regletsdk.Config{
				"hostname":    "example.com",
				"record_type": "A",
				"nameserver":  "8.8.8.8:53",
			},
			wantError: false,
		},
		{
			name: "Missing hostname",
			config: regletsdk.Config{
				"record_type": "A",
			},
			wantError: true,
			errMsg:    "Hostname' failed on the 'required' tag",
		},
		{
			name: "Invalid record type",
			config: regletsdk.Config{
				"hostname":    "example.com",
				"record_type": "INVALID",
			},
			wantError: true,
			errMsg:    "RecordType' failed on the 'oneof' tag",
		},
		{
			name:      "Empty config (missing hostname)",
			config:    regletsdk.Config{},
			wantError: true,
			errMsg:    "Hostname' failed on the 'required' tag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evidence, err := plugin.Check(context.Background(), tt.config)
			require.NoError(t, err, "Check should not return a Go error directly")

			if tt.wantError {
				assert.False(t, evidence.Status, "Expected status to be false for config error")
				require.NotNil(t, evidence.Error, "Expected evidence to contain an error")
				assert.Contains(t, evidence.Error.Message, tt.errMsg)
				assert.Equal(t, "config", evidence.Error.Type)
			} else {
				assert.True(t, evidence.Status, "Expected status to be true")
			}
		})
	}
}

func TestDNSPlugin_Check_Success(t *testing.T) {
	mockLookup := func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
		return &DNSLookupResult{
			Records: []string{"93.184.216.34", "93.184.216.35"},
		}, nil
	}

	plugin := &dnsPlugin{LookupFunc: mockLookup}
	config := regletsdk.Config{
		"hostname":    "example.com",
		"record_type": "A",
	}

	evidence, err := plugin.Check(context.Background(), config)
	require.NoError(t, err)

	assert.True(t, evidence.Status)
	assert.Equal(t, 2, evidence.Data["record_count"])
	assert.Equal(t, false, evidence.Data["is_timeout"])
	assert.Equal(t, false, evidence.Data["is_not_found"])
}

func TestDNSPlugin_Check_MXRecords(t *testing.T) {
	mockLookup := func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
		return &DNSLookupResult{
			MXRecords: []MXRecord{
				{Host: "mail1.example.com", Pref: 10},
				{Host: "mail2.example.com", Pref: 20},
			},
		}, nil
	}

	plugin := &dnsPlugin{LookupFunc: mockLookup}
	config := regletsdk.Config{
		"hostname":    "example.com",
		"record_type": "MX",
	}

	evidence, err := plugin.Check(context.Background(), config)
	require.NoError(t, err)

	assert.True(t, evidence.Status)
	assert.Equal(t, 2, evidence.Data["record_count"])
}

func TestDNSPlugin_Check_NotFound(t *testing.T) {
	mockLookup := func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
		return &DNSLookupResult{
			Error: &DNSError{
				Message:    "no such host",
				Type:       "network",
				IsNotFound: true,
			},
		}, nil
	}

	plugin := &dnsPlugin{LookupFunc: mockLookup}
	config := regletsdk.Config{
		"hostname": "nonexistent.example.com",
	}

	evidence, err := plugin.Check(context.Background(), config)
	require.NoError(t, err)

	assert.False(t, evidence.Status)
	assert.Equal(t, true, evidence.Data["is_not_found"])
}

func TestDNSPlugin_Check_LookupError(t *testing.T) {
	mockLookup := func(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
		return nil, errors.New("network unreachable")
	}

	plugin := &dnsPlugin{LookupFunc: mockLookup}
	config := regletsdk.Config{
		"hostname": "example.com",
	}

	evidence, err := plugin.Check(context.Background(), config)
	require.NoError(t, err)

	assert.False(t, evidence.Status)
}
