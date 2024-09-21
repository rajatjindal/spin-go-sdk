// A generated module for SpinGoSdk functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/spin-go-sdk/internal/dagger"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Builder struct {
	base             string
	tinygoVersion    string
	goVersion        string
	spinVersion      string
	wasmToolsVersion string
}

func (m *Builder) WithTinyGoVersion(version string) *Builder {
	m.tinygoVersion = version
	return m
}

func (m *Builder) WithGoVersion(version string) *Builder {
	m.goVersion = version
	return m
}

func (m *Builder) WithSpinVersion(version string) *Builder {
	m.spinVersion = version
	return m
}

func (m *Builder) WithWasmToolsVersion(version string) *Builder {
	m.wasmToolsVersion = version
	return m
}

func (m *Builder) applyDefaults() *Builder {
	if m.spinVersion == "" {
		m.spinVersion = "canary"
	}

	if m.tinygoVersion == "" {
		m.tinygoVersion = "0.33.0"
	}

	if m.goVersion == "" {
		m.goVersion = "1.22.6"
	}

	if m.wasmToolsVersion == "" {
		m.wasmToolsVersion = "1.216.0"
	}

	return m
}

func (m *Builder) BuildEnv(ctx context.Context) *dagger.Container {
	// TODO(rajatjindal): maybe there is a better way to do this
	m.applyDefaults()

	// the base has to be ubuntu. I tried alpine, that didn't work
	return dag.Container().From("ubuntu:latest").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "wget"}).
		With(withSpin(m.spinVersion)).
		With(withTinyGo(m.tinygoVersion)).
		With(withGo(m.goVersion)).
		With(withWasmTools(m.wasmToolsVersion))
}

func (m *Builder) BuildExample(ctx context.Context, source *dagger.Directory, example string) *dagger.Container {
	return m.BuildEnv(ctx).
		WithDirectory("/app", source).
		WithWorkdir(fmt.Sprintf("/app/v2/examples/%s", example)).
		WithExec([]string{"spin", "build"}).
		WithExec([]string{"spin", "up", "-f", fmt.Sprintf("/app/v2/examples/%s", "http"), "--listen", "0.0.0.0:3000"}).
		WithExposedPort(3000)
}

func (m *Builder) TestIncomingHttp(ctx context.Context, source *dagger.Directory) error {
	svc := m.BuildExample(ctx, source, "http").AsService()

	output, err := dag.Container().
		WithServiceBinding("test-service", svc).
		From("curlimages/curl:latest").
		WithExec([]string{"sh", "-c", fmt.Sprintf("curl -vXGET 'http://test-service:3000/hello'")}).
		Stdout(ctx)
	if err != nil {
		return err
	}

	if !strings.Contains(output, "Hello Fermyon!") {
		return fmt.Errorf("expected output to contain `Hello Fermyon!`. output: ", output)
	}

	return nil
}

func withTinyGo(version string) dagger.WithContainerFunc {
	return func(c *dagger.Container) *dagger.Container {
		releaseArtifactName := fmt.Sprintf("tinygo%s.%s-%s", version, runtime.GOOS, runtime.GOARCH)
		releaseArtifactTarFile := fmt.Sprintf("%s.tar.gz", releaseArtifactName)
		releaseArtifactDownloadLink := fmt.Sprintf("https://github.com/tinygo-org/tinygo/releases/download/v%s/%s", version, releaseArtifactTarFile)
		return c.
			WithExec([]string{"wget", releaseArtifactDownloadLink, "-O", releaseArtifactTarFile}).
			WithExec([]string{"tar", "-xvf", releaseArtifactTarFile}).
			WithExec([]string{"mkdir", "-p", "/opt"}).
			WithExec([]string{"mv", "tinygo", "/opt/tinygo"}).
			WithEnvVariable("TINYGOROOT", "/opt/tinygo").
			WithEnvVariable("PATH", "/opt/tinygo/bin:$PATH", dagger.ContainerWithEnvVariableOpts{
				Expand: true,
			})
	}
}

func withSpin(version string) dagger.WithContainerFunc {
	return func(c *dagger.Container) *dagger.Container {
		// TODO(rajatjindal): allow pulling specific version
		// right now defaults to canary
		return c.WithFile("/usr/local/bin/spin", dag.Container().
			From("ghcr.io/fermyon/spin:canary-distroless").File("/usr/local/bin/spin"))
	}
}

func withGo(version string) dagger.WithContainerFunc {
	return func(c *dagger.Container) *dagger.Container {
		releaseArtifactName := fmt.Sprintf("go%s.%s-%s", version, runtime.GOOS, runtime.GOARCH)
		releaseArtifactTarFile := fmt.Sprintf("%s.tar.gz", releaseArtifactName)
		releaseArtifactDownloadLink := fmt.Sprintf("https://go.dev/dl/%s", releaseArtifactTarFile)
		return c.
			WithExec([]string{"wget", releaseArtifactDownloadLink, "-O", releaseArtifactTarFile}).
			WithExec([]string{"rm", "-rf", "/usr/local/go"}).
			WithExec([]string{"tar", "-C", "/usr/local", "-xvf", releaseArtifactTarFile}).
			WithEnvVariable("PATH", "/usr/local/go/bin:$PATH", dagger.ContainerWithEnvVariableOpts{
				Expand: true,
			})
	}
}

func withWasmTools(version string) dagger.WithContainerFunc {
	return func(c *dagger.Container) *dagger.Container {
		arch := "x86_64"
		if runtime.GOARCH == "arm64" {
			arch = "aarch64"
		}

		releaseArtifactName := fmt.Sprintf("wasm-tools-%s-%s-%s", version, arch, runtime.GOOS)
		releaseArtifactTarFile := fmt.Sprintf("%s.tar.gz", releaseArtifactName)
		releaseArtifactDownloadLink := fmt.Sprintf("https://github.com/bytecodealliance/wasm-tools/releases/download/v%s/%s", version, releaseArtifactTarFile)
		return c.
			WithExec([]string{"wget", releaseArtifactDownloadLink, "-O", releaseArtifactTarFile}).
			WithExec([]string{"tar", "-xvf", releaseArtifactTarFile}).
			WithExec([]string{"mv", filepath.Join(releaseArtifactName, "wasm-tools"), "/usr/local/bin/wasm-tools"})
	}
}
