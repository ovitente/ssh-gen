package main

import (
	"fmt"
)

func main() {

	type sshConfigBlock struct {
		hostAlias    string
		hostName     string
		userName     string
		portNumber   string
		identityFile string
		additionOpts string
	}

	type sshConfigOpts struct {
		optName string
		optValue string
	}

	element := sshConfigBlock{"Radas", "radas.gitlab.com", "bob", "223", "~/.ssh/work", "yes"}

	fmt.Printf("Host %s\n Hostname %s\n User %s\n Port %s\n IdentityFile %s\n IdentitiesOnly %s\n", element.hostAlias, element.hostName, element.userName, element.portNumber, element.identityFile, element.additionOpts)
// Todo: add additionOpts as integrated structure

}
