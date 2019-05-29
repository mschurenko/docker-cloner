package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types/strslice"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

var (
	version  = "" // will be set at build time
	idSet    = false
	cmdSet   = false
	entrySet = false
)

func visit(f *flag.Flag) {
	if f.Name == "id" {
		idSet = true
	}
	if f.Name == "cmd" {
		cmdSet = true
	}
	if f.Name == "entry_point" {
		entrySet = true
	}
}

func printErr(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	id := flag.String("id", "", "container name or ID (required)")
	cmd := flag.String("cmd", "", "override command (use quotes)")
	entry := flag.String("entry_point", "", "override entry point (use quotes)")
	newName := flag.String("new_name", "my_clone", "name of new container")
	EnvVars := flag.String("env_vars", "", "added to existing env vars. Comma delmited list of key=value pairs")
	interactive := flag.Bool("interactive", true, "set container to be interactive")
	memory := flag.Int("memory", 0, "override memory limit in MB (min is 100 MB)")
	printVer := flag.Bool("version", false, "print version")

	flag.Parse()

	flag.Visit(visit)

	if *printVer {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	if !idSet {
		printErr("id must not be empty")
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		printErr(err.Error())
	}

	cli.NegotiateAPIVersion(ctx)

	cJSON, err := cli.ContainerInspect(ctx, *id)
	if err != nil {
		printErr(err.Error())
	}

	containerConfig := cJSON.Config

	if *EnvVars != "" {
		for _, e := range strings.Split(*EnvVars, ",") {
			containerConfig.Env = append(containerConfig.Env, e)
		}
	}

	if *interactive {
		containerConfig.Tty = true
		containerConfig.AttachStdin = true
		containerConfig.AttachStdin = true
		containerConfig.AttachStderr = true
		containerConfig.OpenStdin = true
		containerConfig.StdinOnce = true
	}

	if cmdSet {
		containerConfig.Cmd = strslice.StrSlice(strings.Fields(*cmd))
	}

	if entrySet {
		containerConfig.Entrypoint = strslice.StrSlice([]string{*entry})
	}

	hostConfig := cJSON.HostConfig
	hostConfig.LogConfig = container.LogConfig{}

	if *memory > 0 {
		if *memory < 100 {
			*memory = 100
		}
		hostConfig.Resources.Memory = int64(*memory * 1024 * 1024)
	}

	// remove parent cgroup
	hostConfig.CgroupParent = ""
	networkingConfig := &network.NetworkingConfig{}

	c, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, *newName)
	if err != nil {
		printErr(err.Error())
	}

	// outputs contianer id to stdout upon success
	fmt.Println(c.ID)
}
