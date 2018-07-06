package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	log.Printf("Printed Containers: %d", len(containers))
}

//TODO: add more parameters to customize
func RunContainer() {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//steps
	//select image name
	//assign ports
	//assign container name

	//contianer config

	//host config

	//network config

	cli.ContainerCreate(context.Background(), "mongoFromGolangCode")

}

func GetMongoURL() string {

	mongoContainerID := ""
	mongoContainerURL := ""

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == "mongo" {

			mongoContainerID = container.ID
			fmt.Printf("found mongo %s %s\n", mongoContainerID, container.Image)

			containerJson, _ := cli.ContainerInspect(context.Background(), mongoContainerID)

			//log.Printf("Inspect: %v", containerJson)
			mongoContainerURL = containerJson.NetworkSettings.IPAddress
			log.Printf("Mongo URL: %v", mongoContainerURL)
		}

		fmt.Printf("Container Info %T\n", container)
	}
	log.Printf("Number of Containers: %d", len(containers))

	return mongoContainerURL
}
