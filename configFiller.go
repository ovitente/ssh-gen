package main

import (
	"fmt"
)

func configFiller() {

	type sshConfigOpts struct {
		optName  string
		optValue string
	}

	type sshConfigBlock struct {
		hostAlias    string
		hostName     string
		userName     string
		portNumber   string
		identityFile string
		additionOpts sshConfigOpts
	}

	// Filling structure with values
	confOpt := sshConfigOpts{
		"IdentitiesOnly",
		"yes",
	}

	confElement := sshConfigBlock{
		"Radas",
		"radas.gitlab.com",
		"bob",
		"223",
		"~/.ssh/work",
		confOpt,
	}

	// Sample printing structure values to the console
	fmt.Printf(
		"Host %s\n Hostname %s\n User %s\n Port %s\n IdentityFile %s\n %s %s\n",
		confElement.hostAlias,
		confElement.hostName,
		confElement.userName,
		confElement.portNumber,
		confElement.identityFile,
		confOpt.optName,
		confOpt.optValue,
	)


}
