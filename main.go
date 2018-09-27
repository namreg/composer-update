package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const sedPattern = `s,%s": *".\+",%[1]s": "%s",`

func main() {
	projects := flag.String("projects", "", "Projects to update separated by a comma.")
	dep := flag.String("dep", "", `Dependency to update in format "name@master".`)
	basepath := flag.String("basepath", "/var/www", "Base path for projects.")
	branch := flag.String("branch", "master", `Branch to create.`)

	flag.Parse()

	if *projects == "" {
		fmt.Println("projects can not be empty")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *dep == "" {
		fmt.Println("dependency can not be empty")
		flag.PrintDefaults()
		os.Exit(1)
	}
	depConf := strings.SplitN(*dep, "@", 2)
	if len(depConf) != 2 {
		fmt.Println("dependency has invalid format")
		flag.PrintDefaults()
		os.Exit(1)
	}

	wg := new(sync.WaitGroup)
	for _, p := range strings.Split(*projects, ",") {
		p = strings.TrimSpace(p)
		wg.Add(1)
		go func(p string) {
			path := filepath.Join(*basepath, p, "current")
			depName := depConf[0]
			depVer := depConf[1]
			updateProject(path, depName, depVer, *branch)
			wg.Done()
		}(p)
	}
	wg.Wait()
}

func updateProject(path, depName, depVersion, branch string) {
	cmds := make([]*exec.Cmd, 0)
	cmds = append(cmds, exec.Command("git", "fetch", "origin", "+refs/heads/master"))
	cmds = append(cmds, exec.Command("git", "reset", "--hard", "origin/master"))
	if branch != "master" {
		cmds = append(cmds, exec.Command("git", "checkout", "-b", branch))
	}
	sp := fmt.Sprintf(sedPattern, depName, depVersion)
	cmds = append(cmds, exec.Command("sed", "-i", sp, filepath.Join(path, "composer.json")))
	cmds = append(cmds, exec.Command("composer", "update", depName))
	cmds = append(cmds, exec.Command("git", "push", "origin", branch))

	for _, cmd := range cmds {
		cmd.Dir = path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("+ %s [%s]\n", strings.Join(cmd.Args, " "), cmd.Dir)
		if err := cmd.Run(); err != nil {
			fmt.Printf("could not execute command: %v\n", err)
			break
		}
	}
}
