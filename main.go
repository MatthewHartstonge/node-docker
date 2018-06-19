package main

import (
	"fmt"
	"strings"
	"os/exec"
	"os"
	"text/template"
	"syscall"
	"log"
)

var nodeVersions = []string{"6", "8", "9", "10"}
var dockerVersions = []string{
	// Versions obtained from: https://apt.dockerproject.org/repo/dists/debian-jessie/main/filelist
	"1.5.0-0~jessie",
	"1.6.0-0~jessie",
	"1.6.1-0~jessie",
	"1.6.2-0~jessie",
	"1.7.0-0~jessie",
	"1.7.1-0~jessie",
	"1.8.0-0~jessie",
	"1.8.1-0~jessie",
	"1.8.2-0~jessie",
	"1.8.3-0~jessie",
	"1.9.0-0~jessie",
	"1.9.1-0~jessie",
	"1.10.0-0~jessie",
	"1.10.1-0~jessie",
	"1.10.2-0~jessie",
	"1.10.3-0~jessie",
	"1.11.0-0~jessie",
	"1.11.1-0~jessie",
	"1.11.2-0~jessie",
	"1.12.0-0~jessie",
	"1.12.1-0~jessie",
	"1.12.2-0~jessie",
	"1.12.3-0~jessie",
	"1.12.4-0~debian-jessie",
	"1.12.5-0~debian-jessie",
	"1.12.6-0~debian-jessie",
	"1.13.0-0~debian-jessie",
	"1.13.1-0~debian-jessie",
	"17.03.0~ce-0~debian-jessie",
	"17.03.1~ce-0~debian-jessie",
	"17.04.0~ce-0~debian-jessie",
	"17.05.0~ce-0~debian-jessie",
}

const dockerFileTemplate = `FROM node:{{.NodeVersion}}-slim
LABEL maintainer="Matthew Hartstonge <matt@mykro.co.nz>" \
      repo="https://github.com/matthewhartstonge/node-docker"

ENV DOCKER_VERSION="{{.DockerVersion}}"
RUN apt-get update \
    && apt-get install -y \
        apt-transport-https \
        ca-certificates \
        g++ \
        make \
        python \
    && apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D \
    && echo "deb https://apt.dockerproject.org/repo debian-jessie main" >> /etc/apt/sources.list.d/docker.list \
    && apt-get update && apt-get install -y \
        docker-engine="${DOCKER_VERSION}" \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
`

type Model struct {
	NodeVersion   string
	DockerVersion string
}

func prettyDockerVersion(dockerVersion string) string {
	docker := strings.Split(dockerVersion, "-0~")[0]
	docker = strings.Replace(docker, "~ce", "ce", 1)
	return docker
}

func gitBranchRemoteExists(nodeVersion, dockerVersion string) bool {
	var (
		cmdOut []byte
		err    error
	)
	log.Println("checking if remote git branch exists...")

	cmd := "git"
	branch := fmt.Sprintf("node%s-docker%s", nodeVersion, prettyDockerVersion(dockerVersion))
	args := []string{"ls-remote", "--heads", "git@github.com:matthewhartstonge/node-docker.git", branch}

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	found := strings.TrimSpace(string(cmdOut))
	if len(found) > 1 {
		// Branch exists, don't need to process version
		log.Println("remote git branch exists!")
		return true
	}
	log.Println("remote git branch does not exist!")
	return false
}

func gitBranchLocalExists(nodeVersion, dockerVersion string) bool {
	var (
		err error
	)

	log.Println("checking if local git branch exists...")

	cmd := "git"
	branch := fmt.Sprintf("node%s-docker%s", nodeVersion, prettyDockerVersion(dockerVersion))
	args := []string{"rev-parse", "--quiet", "--verify", branch}

	command := exec.Command(cmd, args...)
	_, err = command.Output()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 1 {
					// we know it couldn't be found
					log.Println("local git branch does not exist!")
					return false
				}
			}
		}

		// Some other error
		log.Println("errr... Something broke. Halp.")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	log.Println("local git branch does not exist!")
	return true
}

func gitCheckoutBranch(branch string, isNew bool) {
	var (
		cmdOut []byte
		err    error
	)

	log.Printf("checking out git branch %s...\n", branch)

	cmd := "git"
	args := []string{"checkout", branch}
	if isNew {
		args = []string{"checkout", "-b", branch}
	}

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Printf("failed to check out git branch %s!\n", branch)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, string(cmdOut))
		os.Exit(1)
	}

	log.Printf("checked out git branch %s!\n", branch)
}

func gitAdd(fpath string) {
	var (
		cmdOut []byte
		err    error
	)

	log.Printf("adding %s to git commit...\n", fpath)

	cmd := "git"
	args := []string{"add", fpath}

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Printf("failed to add %s to git commit!\n", fpath)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, string(cmdOut))
		os.Exit(1)
	}

	log.Printf("added %s to git commit!\n", fpath)
}

func gitCommit(nodeVersion, dockerVersion string) {
	var (
		cmdOut []byte
		err    error
	)

	cmd := "git"
	commitMsg := fmt.Sprintf(":new: Adds support for Node %s with Docker %s", nodeVersion, prettyDockerVersion(dockerVersion))
	args := []string{"commit", "-m", commitMsg}

	log.Printf("commiting %s to git...\n", commitMsg)

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Printf("failed commiting %s to git!\n", commitMsg)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, string(cmdOut))
		os.Exit(1)
	}

	log.Printf("commited %s to git!\n", commitMsg)
}

func generateDockerFile(nodeVersion, dockerVersion string) {
	version := fmt.Sprintf("node%s-docker%s", nodeVersion, prettyDockerVersion(dockerVersion))
	log.Printf("generating Dockerfile for %s...\n", version)

	t, err := template.New("").Parse(dockerFileTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("Dockerfile")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	model := Model{
		NodeVersion:   nodeVersion,
		DockerVersion: dockerVersion,
	}
	err = t.Execute(f, model)
	if err != nil {
		panic(err)
	}

	log.Printf("generated Dockerfile for %s!\n", version)
}

func main() {
	for _, nodeVersion := range nodeVersions {
		for _, dockerVersion := range dockerVersions {
			if gitBranchRemoteExists(nodeVersion, dockerVersion) {
				continue
			}

			if gitBranchLocalExists(nodeVersion, dockerVersion) {
				continue
			}

			newVersion := fmt.Sprintf("node%s-docker%s", nodeVersion, prettyDockerVersion(dockerVersion))

			gitCheckoutBranch("development", false)
			gitCheckoutBranch(newVersion, true)
			generateDockerFile(nodeVersion, dockerVersion)
			gitAdd("Dockerfile")
			gitCommit(nodeVersion, dockerVersion)
		}
	}
}
