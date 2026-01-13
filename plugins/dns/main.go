// Package main provides a DNS plugin for Reglet.
// This is compiled to WASM and loaded by the Reglet runtime.
//go:build wasip1

package main

import (
	"context"
	"log/slog"

	regletsdk "github.com/reglet-dev/reglet-sdk/go"
	regletnet "github.com/reglet-dev/reglet-sdk/go/net"
)

// wrapLookup wraps the SDK's resolver to return local types
func wrapLookup(ctx context.Context, hostname, recordType, nameserver string) (*DNSLookupResult, error) {
	resolver := &regletnet.WasmResolver{Nameserver: nameserver}
	wireResult, err := resolver.Lookup(ctx, hostname, recordType)
	if err != nil {
		return nil, err
	}

	result := &DNSLookupResult{}

	if wireResult.Error != nil {
		result.Error = &DNSError{
			Message:    wireResult.Error.Message,
			Type:       wireResult.Error.Type,
			IsTimeout:  wireResult.Error.IsTimeout,
			IsNotFound: wireResult.Error.IsNotFound,
		}
		return result, nil
	}

	result.Records = wireResult.Records
	if wireResult.MXRecords != nil {
		for _, mx := range wireResult.MXRecords {
			result.MXRecords = append(result.MXRecords, MXRecord{
				Host: mx.Host,
				Pref: mx.Pref,
			})
		}
	}

	return result, nil
}

func init() {
	slog.Info("DNS plugin init() started")
	regletsdk.Register(&dnsPlugin{
		LookupFunc: wrapLookup,
	})
	slog.Info("DNS plugin init() registered")
}

// main function for the WASM plugin.
func main() {}
