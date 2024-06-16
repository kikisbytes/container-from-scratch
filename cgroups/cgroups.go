package cgroups

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strconv"
)

func ConfigureCgroup() {
  cgroupRoot := "/sys/fs/cgroup"
  unifiedCgroup := filepath.Join(cgroupRoot, "mydocker")

  fmt.Printf("Creating unified cgroup at %s\n", unifiedCgroup)
  os.MkdirAll(unifiedCgroup, 0755);

  // Set memory limit
  memoryMax := filepath.Join(unifiedCgroup, "memory.max")
  fmt.Printf("Setting memory max at %s\n", memoryMax)
  // set memory limit to 100MB
  must(ioutil.WriteFile(memoryMax, []byte("100000000"), 0700))

  // add process to cgroup
  cgroupProcs := filepath.Join(unifiedCgroup, "cgroup.procs")
  fmt.Printf("Adding process to unified cgroup at %s\n", cgroupProcs)
  must(ioutil.WriteFile(cgroupProcs, []byte(strconv.Itoa(os.Getpid())), 0644))
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
