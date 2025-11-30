package main

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
)

func main() {
	CreatePCollection()
}

// Tutorial 1 . Create a pipeline
func CreatePipeline() *beam.Pipeline {
	// setup default logging ,schema ...
	beam.RegisterInit(func() {
		log.SetupLogging("debug", "dev")
	})
	beam.Init()
	return beam.NewPipeline()
}

func CreatePipelineWithRoot() (beam.Scope, *beam.Pipeline) {
	beam.Init()
	p, s := beam.NewPipelineWithRoot()
	return s, p
}

// Tutorial 2 . Create a PCollection
func CreatePCollection() {
	beam.Init()

	// First create pipeline
	_, s := beam.NewPipelineWithRoot()

	//Now create the PCollection using list of strings
	beam.Create(s, "To", "be", "or", "not", "to", "be", "that", "is", "the", "question")

	// create List 
	beam.CreateList(s,[]string{"hello"})

}
