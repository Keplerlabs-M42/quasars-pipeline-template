package main

import (
	"context"
	"dagger/quasars-pipeline-template/internal/dagger"
	"fmt"
)

type QuasarsPipelineTemplate struct{}

func (m *QuasarsPipelineTemplate) Build(
	src *dagger.Directory,

	// +optional
	// +default="1.24"
	version string,

	// +optional
	// +default="."
	mainFile string,
) *dagger.Container {
	image := fmt.Sprintf("golang:%s", version)

	return dag.Container().
		From(image).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "build/", mainFile})
}

func (m *QuasarsPipelineTemplate) Lint(ctx context.Context, src *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golangci/golangci-lint:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run"})
}
