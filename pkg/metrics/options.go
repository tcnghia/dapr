// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package metrics

import (
	"strconv"
)

const (
	defaultMetricsPort            = "9090"
	defaultMetricsExporterEnabled = true
)

// Options defines the sets of options for Dapr metrics
type Options struct {
	// Enable the metrics exporter
	MetricsExporterEnabled bool

	metricsPort string
}

func defaultMetricOptions() *Options {
	return &Options{
		metricsPort:            defaultMetricsPort,
		MetricsExporterEnabled: defaultMetricsExporterEnabled,
	}
}

// MetricsPort gets metrics port.
func (o *Options) MetricsPort() uint64 {
	port, err := strconv.ParseUint(o.metricsPort, 10, 64)
	if err != nil {
		// Use default metrics port as a fallback
		port, _ = strconv.ParseUint(defaultMetricsPort, 10, 64)
	}

	return port
}

// AttachCmdFlags attaches metrics options to command flags
func (o *Options) AttachCmdFlags(
	stringVar func(p *string, name string, value string, usage string),
	boolVar func(p *bool, name string, value bool, usage string)) {
	stringVar(
		&o.metricsPort,
		"metrics-port",
		defaultMetricsPort,
		"The port for the metrics server")
	boolVar(
		&o.MetricsExporterEnabled,
		"enable-metrics-exporter",
		defaultMetricsExporterEnabled,
		"Enable prometheus metric exporter")
}
