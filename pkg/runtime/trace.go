// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package runtime

import (
	"go.opencensus.io/trace"
)

// exporterStore allows us to capture the trace exporter store registrations.
//
// This is needed because the OpenCensus library only expose global methods for
// exporter registration.
type exporterStore interface {
	// RegisterExporter registers a trace.Exporter.
	RegisterExporter(exporter trace.Exporter)
	// ApplyConfig applies a trace.Config to the underlying trace exporter store.
	ApplyConfig(config trace.Config)
}

// openCensusExporterStore is an implementation of exporterStore
// that makes use of OpenCensus's library's global exporer stores (`trace`).
type openCensusExporterStore struct{}

// RegisterExporter implements exporterStore using OpenCensus's global registration.
func (s openCensusExporterStore) RegisterExporter(exporter trace.Exporter) {
	trace.RegisterExporter(exporter)
}

// ApplyConfig implements exporterStore using OpenCensus's global registration.
func (s openCensusExporterStore) ApplyConfig(config trace.Config) {
	trace.ApplyConfig(config)
}

// fakeExporterStore implements exporterStore by merely recrod the exporters
// and config that were registered/applied.
//
// This is only for use in unit tests.
type fakeExporterStore struct {
	exporters []trace.Exporter
	config    *trace.Config
}

func (r *fakeExporterStore) RegisterExporter(exporter trace.Exporter) {
	r.exporters = append(r.exporters, exporter)
}

func (r *fakeExporterStore) ApplyConfig(config trace.Config) {
	r.config = &config
}
