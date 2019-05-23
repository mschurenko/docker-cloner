package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <container_id>", os.Args[0])
		os.Exit(1)
	}

	cID := os.Args[1]

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	cli.NegotiateAPIVersion(ctx)

	cJSON, err := cli.ContainerInspect(ctx, cID)
	if err != nil {
		log.Fatal(err)
	}

	cConf := cJSON.Config
	hConf := cJSON.HostConfig
	nConf := &network.NetworkingConfig{}

	c, err := cli.ContainerCreate(ctx, cConf, hConf, nConf, "new-test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting container", c.ID)

	if err := cli.ContainerStart(ctx, c.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}

}
