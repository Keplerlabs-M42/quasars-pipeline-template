package main

import (
	"context"
	"dagger/quasars-pipeline-template/internal/dagger"
)

type QuasarsPipelineTemplate struct{}

func (m *QuasarsPipelineTemplate) Build(ctx context.Context, src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.24-alpine").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "build/"})
}

func (m *QuasarsPipelineTemplate) Lint(ctx context.Context, src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golangci/golangci-lint:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run"})
}
