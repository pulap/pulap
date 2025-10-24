package config

import "github.com/pulap/pulap/pkg/lib/core"

type XParams struct {
	log    core.Logger
	cfg    *Config
	tracer core.Tracer
}

func NewXParams(log core.Logger, cfg *Config) XParams {
	if log == nil {
		log = core.NewNoopLogger()
	}
	return XParams{
		log:    log,
		cfg:    cfg,
		tracer: core.NoopTracer(),
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
