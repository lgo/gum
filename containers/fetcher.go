package containers

import (
	"github.com/fsouza/go-dockerclient"
)

var Client = initClient()

func initClient() *docker.Client {
	endpoint := "unix:///var/run/docker.sock"
	clnt, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	return clnt
}

func Pull(repository string, tag string) error {
	return Client.PullImage(docker.PullImageOptions{
		Repository: repository,
		Tag:        tag,
	}, docker.AuthConfiguration{})
}
