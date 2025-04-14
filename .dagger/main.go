package main

import (
	"dagger/tutorial/internal/dagger"
)

const (
	helmVersion     = "3.16.1"
	helmDocsVersion = "v1.14.2"
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
func (m *Tutorial) Build() *dagger.Container {
	return dag.Container().
		From("nginx:1.16.0").
		WithFile("/usr/share/nginx/html/index.html", m.Source.File("index.html"))
}

// Run the application (for demo purposes).
func (m *Tutorial) Serve() *dagger.Service {
	return m.Build().WithExposedPort(80).AsService()
}
