package main

import (
	"os/exec"
	"strings"

	"github.com/go-check/check"
)

func isNetworkPresent(c *check.C, name string) bool {
	runCmd := exec.Command(dockerBinary, "network", "ls")
	out, _, _, err := runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}
	lines := strings.Split(out, "\n")
	for i := 1; i < len(lines)-1; i++ {
		if strings.Contains(lines[i], name) {
			return true
		}
	}
	return false
}

func (s *DockerSuite) TestDockerNetworkLsDefault(c *check.C) {
	defaults := []string{"bridge", "host", "none"}
	for _, nn := range defaults {
		if !isNetworkPresent(c, nn) {
			c.Fatalf("Missing Default network : %s", nn)
		}
	}
}

func (s *DockerSuite) TestDockerNetworkCreateDelete(c *check.C) {
	runCmd := exec.Command(dockerBinary, "network", "create", "test")
	out, _, _, err := runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}
	if !isNetworkPresent(c, "test") {
		c.Fatalf("Network test not found")
	}

	runCmd = exec.Command(dockerBinary, "network", "rm", "test")
	out, _, _, err = runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}
	if isNetworkPresent(c, "test") {
		c.Fatalf("Network test is not removed")
	}
}

func (s *DockerSuite) TestDockerNetworkValidRunContainer(c *check.C) {
	runCmd := exec.Command(dockerBinary, "network", "create", "udn")
	out, _, _, err := runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}

	out, _, err = runCommandWithOutput(exec.Command(dockerBinary, "run", "-d", "--net", "udn", "--name", "host_container", "busybox", "top"))
	if err != nil {
		c.Fatal(err, out)
	}

	out, _, err = runCommandWithOutput(exec.Command(dockerBinary, "rm", "-f", "host_container"))
	if err != nil {
		c.Fatal(err, out)
	}

	runCmd = exec.Command(dockerBinary, "network", "rm", "udn")
	out, _, _, err = runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}
}

func (s *DockerSuite) TestDockerNetworkInvalidRunContainer(c *check.C) {
	_, _, err := runCommandWithOutput(exec.Command(dockerBinary, "run", "-d", "--net", "invalid", "--name", "host_container", "busybox", "top"))
	if err == nil {
		c.Fatal("Invalid user-defined network should fail container creation")
	}
}

func (s *DockerSuite) TestDockerNetworkLinksUserNetwork(c *check.C) {
	runCmd := exec.Command(dockerBinary, "network", "create", "userdefined")
	out, _, _, err := runCommandWithStdoutStderr(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}

	out, _, err = runCommandWithOutput(exec.Command(dockerBinary, "run", "-d", "--net", "userdefined", "--name", "host_container", "busybox", "top"))
	if err != nil {
		c.Fatal(err, out)
	}

	out, _, err = runCommandWithOutput(exec.Command(dockerBinary, "run", "--name", "should_fail", "--link", "host_container:tester", "busybox", "true"))
	if err == nil || !strings.Contains(out, "User-Defined --netoption can't be used with links") {
		c.Fatalf("Running container linking to a container with an user-defined --net should have failed: %s", out)
	}
}
