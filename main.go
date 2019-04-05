package main

import (
	"fmt"
	"os"
	"os/exec"
)


// Modify to the structure https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go

func main()  {

	mount_dir := "~/.mount"
	remote_path := "/"
	server := os.Args[1]

	fmt.Println(mount_dir, server, remote_path)

//	Check to zero vars: mount_dir server and conpath
// Give server argument from commandline cli
	InitSshConnection(mount_dir, server)
}

func PrintUsage() {
	fmt.Println("Usage : sshmount <servername>")
//	Put panic with msg

}

func InitSshConnection(mount_dir, server string) {
	if len(server) == 0 {
		fmt.Println("Server string is empty.")
		os.Exit(1)
	}
	fmt.Printf("You will be connected to the [ %s ]\n", server)
	fmt.Printf("Directory has been mounted to %s/%s\n", mount_dir, server)

	cmd := exec.Command("df", "-h")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
}