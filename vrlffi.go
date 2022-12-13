package main

import (
	"fmt"
	"log"

	govrl "github.com/gh123man/go-vrl/v5"
)

type VRLFFIProcessor struct {
	runtime         *govrl.Runtime
	program         *govrl.Program
	originalPattern string
}

func NewVRLFFIProcessor(pattern string) *VRLFFIProcessor {
	program, err := govrl.CompileWithExternal(fmt.Sprintf(`replace(., "%s", "rust")`, pattern), govrl.GetExternalEnv(govrl.Bytes, govrl.Bytes))

	if err != nil {
		log.Panicln(err)
	}

	runtime := govrl.NewRuntime()

	return &VRLFFIProcessor{
		originalPattern: pattern,
		runtime:         runtime,
		program:         program,
	}
}

func (r *VRLFFIProcessor) Process(in string) string {
	res, err := r.runtime.Resolve(r.program, in)
	if err != nil {
		log.Panicln(err)
	}
	return res
}

func (r *VRLFFIProcessor) Describe() string {
	return fmt.Sprintf("VRL via FFI, pattern=%q", r.originalPattern)
}
