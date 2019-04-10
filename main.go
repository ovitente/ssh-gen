package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)


// Modify to the structure https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go

func main()  {

	mount_dir := "~/.mount"
	remote_path := "/"
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	server := os.Args[1]

	InitSshConnection(mount_dir, server, remote_path)
}

func PrintUsage() {
	fmt.Println("Usage : sshmount <servername>")
//	Put panic with msg

}

func InitSshConnection(mount_dir, server, remote_path string) {
	if len(server) == 0 {
		fmt.Println("Server string is empty.")
		os.Exit(1)
	}

	// Creating new dir
	os.Mkdir(os.Getenv("HOME") + "/.mount/" + server, os.ModePerm)

	// Unmounting volume if exist.
	cmd := exec.Command("fusermount", "-u", os.Getenv("HOME") + "/.mount/" + server)
	//cmd := exec.Command("umount", os.Getenv("HOME") + "/.mount/" + server)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}


	//Mounting sshfs resource
	cmd = exec.Command("sshfs", server + ":" + remote_path, os.Getenv("HOME") + "/.mount/" + server)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// Show mount path
	fmt.Printf("Resource [ %s ] was mounted to %s/.mount/%s\n", server, os.Getenv("HOME"), server)
}