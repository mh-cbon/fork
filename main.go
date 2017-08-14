// Package fork helps to fork a repository in your GOPATH.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var name = "fork"
var version = "HEAD"

func main() {

	var force bool
	var help bool
	var h bool

	flag.BoolVar(&h, "h", false, "Show help")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&force, "force", false, "force")

	flag.Parse()

	if h || help {
		showHelp()
		os.Exit(0)
	}

	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		fmt.Println("GOPATH environment variable not set")
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Wrong usage")
		fmt.Println("")
		showHelp()
		os.Exit(1)
	}

	r := args[1]
	fqr := filepath.Join("github.com", r)
	fqr = strings.Replace(fqr, "\\", "/", -1)

	//todo: add more detection mechanism
	p := filepath.Join(gopath, "src", "github.com", r)
	if _, err := os.Stat(p); !os.IsNotExist(err) && !force {
		fmt.Println("the repo was already forked!")
		fmt.Println("")
		fmt.Println("use -force to ignore this error.")
		os.Exit(1)
	} else if !os.IsNotExist(err) {
		os.RemoveAll(p)
	}

	execCmd("go", "get", fqr).
		Or("go get failed....")

	os.Chdir(p)

	execCmd("git", "remote", "rename", "origin", "upstream").
		Or("git remote rename failed....")

	v := strings.Split(r, "/")
	v[0] = args[0]
	execCmd("git", "remote", "add", "origin", "git@github.com:"+strings.Join(v, "/")+".git").
		Or("git remote add failed....")

	fmt.Println("All done!")
	fmt.Println("")
	execCmd("git", "remote", "-v")
	fmt.Printf(`
The fork is available at
	%v

To push your fork on github, create a bare repository in github then run
  git push -u origin master

To create a new branch,
  git checkout -t -b awesome-feature

To upload your fork changes
  git push origin awesome-feature

To avoid merge conflicts, pull the changes from upstream
  git fetch upstream
  git merge upstream/master

Read more: http://blog.sgmansfield.com/2016/06/working-with-forks-in-go/
`, p)
}

func showHelp() {
	fmt.Printf(`%v - %v
	Fork a remote repository to propose a change.

Example
  fork yourID user/repo
  fork mh-cbon kaneshin/pigeon
  fork -force yourID user/repo
  fork -help

Options
    yourID     your username
    user/repo  an identifier like user/repo

Flags
    -force  Force continue if the directory exists.
    -help   Show this help
    -h      Show this help

Good to know

	To push your fork on github, create a bare repository in github then run
	  git push -u origin master

	To create a new branch,
	  git checkout -t -b awesome-feature

	To upload your fork changes
	  git push origin awesome-feature

	To avoid merge conflicts, pull the changes from upstream
	  git fetch upstream
	  git merge upstream/master

`, name, version)
}

type thenWhat struct {
	err error
}

func (t *thenWhat) Or(msg string) {
	if t.err != nil {
		fmt.Println()
		fmt.Println(msg)
		os.Exit(1)
	}
}

func execCmd(b string, args ...string) *thenWhat {
	fmt.Printf("$ %v %v\n", b, args)
	cmd := exec.Command(b, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return &thenWhat{err: cmd.Run()}
}
