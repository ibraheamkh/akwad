package moby

import (
	"context"
	"fmt"

	"github.com/moby/moby/api/types"
	"github.com/moby/moby/client"
)

//This package wraps moby/client api with helper methods

//moby is opes source tools to build container systems
//we will wrap the needed method to make it easeier to run containers in code

func PrintContainers() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
