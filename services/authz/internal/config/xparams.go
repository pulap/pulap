package config

import "github.com/pulap/pulap/pkg/lib/core"

type XParams struct {
	log      core.Logger
	cfg      *Config
	tracer   core.Tracer
	metrics  core.Metrics
	reporter core.ErrorReporter
}

func NewXParams(log core.Logger, cfg *Config) XParams {
	if log == nil {
		log = core.NewNoopLogger()
	}
	return XParams{
		log:      log,
		cfg:      cfg,
		tracer:   core.NoopTracer(),
		metrics:  core.NewNoopMetrics(),
		reporter: core.NewNoopErrorReporter(),
	}
}

func (xp XParams) Log() core.Logger {
	return xp.log
}

func (xp XParams) Cfg() *Config {
	return xp.cfg
}

func (xp XParams) Tracer() core.Tracer {
	return xp.tracer
}

func (xp *XParams) SetTracer(tracer core.Tracer) {
	if tracer == nil {
		tracer = core.NoopTracer()
	}
	xp.tracer = tracer
}

func (xp XParams) Metrics() core.Metrics {
	return xp.metrics
}

func (xp *XParams) SetMetrics(metrics core.Metrics) {
	if metrics == nil {
		metrics = core.NewNoopMetrics()
	}
	xp.metrics = metrics
}

func (xp XParams) ErrorReporter() core.ErrorReporter {
	return xp.reporter
}

func (xp *XParams) SetErrorReporter(reporter core.ErrorReporter) {
	if reporter == nil {
		reporter = core.NewNoopErrorReporter()
	}
	xp.reporter = reporter
}
