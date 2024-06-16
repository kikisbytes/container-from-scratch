package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "mydocker/cgroups"
    "mydocker/dockerfileParser"
)

func main() {
  switch os.Args[1] {
    case "run":
     run()
    case "child":
      child()
    default:
     panic("unknown command")
  }
}


func run() {
  fmt.Printf("Running in parent %v as %d\n", os.Args[2:], os.Getpid())

  // create a new command that will execute the current program in a new process.
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // NEWUTS -> unix timesharing system -> is the host name inside the container
  // NEWNS -> new mount namespace -> isolates the mount points
  // mount is recursively shared between the parent and child
  cmd.SysProcAttr = &syscall.SysProcAttr {
    Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS,
    Unshareflags: syscall.CLONE_NEWNS, // un-share the mount namespace, host won't see the child's mounts
   }

  if err := cmd.Start(); err != nil {
      fmt.Printf("Error: %v\n", err)
      os.Exit(1)
  }

  // Wait for the child process to finish
  if err := cmd.Wait(); err != nil {
      fmt.Printf("Error: %v\n", err)
      os.Exit(1)
  }
}

// should already be in a new namespace
func child() {
  fmt.Printf("Running in child %v as %d\n", os.Args[2:], os.Getpid())

  cgroups.ConfigureCgroup()

  syscall.Sethostname([]byte("container"))
  syscall.Chroot("/docker-fs")
  syscall.Chdir("/")
  syscall.Mount("proc", "proc", "proc", 0, "")

  instructions := dockerfileParser.ParseDockerfile("./express-app/Dockerfile")
  for _, instruction := range instructions {
    fmt.Println("Executing instruction: ", instruction)
  	dockerfileParser.ExecuteInstruction(instruction)
  }

  syscall.Unmount("/proc", 0)
}
