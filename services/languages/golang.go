package datastores

import (
	"github.com/xLegoz/gum/registry"
)

const DefaultConfig = `
version: 1.7
`

func apply(options registry.LanguageOptions) {

}

const goBuildYaml = `
name: GO_BUILDER
config:
  volumes:
    "$GOPATH": /;
    "$(pwd)": /app
  workdir: /app
  cmd: sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'
  image: golang:%f
`

func preload() {
  // Compile step:
  var options = docker.CreateContainerOptions{}
  var version = 1.7
  err := yaml.Unmarshal(fmt.Sprintf(goBuildYaml, version), options)
  container, err := containers.CreateContainer(options)
  err = containers.StartContainer(container.ID, nil)
  err = containers.StopContainer(container.ID, 100)
  err = containers.RemoveContainer(docker.RemoveContainerOptions{
    ID: container.ID,
    RemoveVolumes: true,
    Force: true,
  })
}

func versions(version string) (bool, error) {
  // TODO(xLegoz): check version string to be valid string
  if version != "1.7" {
    return false, error.Error("Non-supported Go version.")
  }

  return true, nil
}

func init() {
	registry.RegisterLanguage(
		"go",
		Language{
			versions: versions,
			apply:    apply,
      preload: preload,
		},
	)
}
