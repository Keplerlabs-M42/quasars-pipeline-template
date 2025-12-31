package main

import (
	"context"
	"dagger/quasars-pipeline-template/internal/dagger"
)

type QuasarsPipelineTemplate struct{}

func (m *QuasarsPipelineTemplate) Build(ctx context.Context, src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.21-alpine").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "build/"})
}
