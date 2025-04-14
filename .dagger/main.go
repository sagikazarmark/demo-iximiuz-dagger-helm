package main

import (
	"context"
	"dagger/tutorial/internal/dagger"
	"fmt"
	"strings"
	"time"
)

// This is a hack to ensure that a few necessary packages are imported.
var (
	_ = time.UTC
	_ = fmt.Sprintf
	_ = strings.Join
)

type Tutorial struct {
	// Project source directory
	//
	// +private
	Source *dagger.Directory
}

func New(
	// Project source directory.
	//
	// +defaultPath="/"
	// +ignore=[".devenv", ".direnv", ".github"]
	source *dagger.Directory,
) *Tutorial {
	return &Tutorial{
		Source: source,
	}
}

// Build the application container.
func (m *Tutorial) Build(_ context.Context) *dagger.Container {
	return dag.Container().
		From("nginx:1.16.0").
		WithFile("/usr/share/nginx/html/index.html", m.Source.File("index.html"))
}

// Run the application (for demo purposes).
func (m *Tutorial) Serve(ctx context.Context) *dagger.Service {
	return m.Build(ctx).WithExposedPort(80).AsService()
}
