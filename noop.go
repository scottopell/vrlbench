package main

type NoopProcessor struct {
}

func NewNoopProcessor() *NoopProcessor {
	return &NoopProcessor{}
}

func (r *NoopProcessor) Process(in string) string {
	return in
}

func (r *NoopProcessor) Describe() string {
	return "No-Op"
}
