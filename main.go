package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	var help bool
	var h bool

	flag.BoolVar(&h, "h", false, "Show help")
	flag.BoolVar(&help, "help", false, "Show help")

	if h || help {
		showHelp()
		os.Exit(0)
	}

	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		fmt.Println("GOPATH environment variable not set")
		os.Exit(1)
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Wrong usage")
		fmt.Println("  fork yourID user/repo")
		fmt.Println("")
		fmt.Println("  yourID: your username")
		fmt.Println("  user/repo: an identifier like user/repo")
		os.Exit(1)
	}

	r := args[1]
	fqr := filepath.Join("github.com", r)
	fqr = strings.Replace(fqr, "\\", "/", -1)

	//todo: add more detection mechanism
	p := filepath.Join(gopath, "src", "github.com", r)
	os.MkdirAll(p, os.ModePerm)
	fmt.Println("forking in " + p)

	execCmd("go", "get", fqr).
		Or("go get failed....")

	os.Chdir(p)

	execCmd("git", "remote", "rename", "origin", "upstream").
		Or("git remote rename failed....")

	//  git remote add origin https://github.com/ScottMansfield/netns

	execCmd("git", "remote", "add", "origin", "git@github.com:"+r+".git").
		Or("git remote add failed....")

	fmt.Println("All done!")
	fmt.Println("")
	execCmd("git", "remote", "-v")
	fmt.Println(`
You can now create a new branch,
  git checkout -t -b awesome-feature

Then put your changes on your origin,
  git push origin awesome-feature

Do not forget to pull the changes from upstream to avoid merge conflicts.,
  git fetch upstream
  git merge upstream/master

Read more: http://blog.sgmansfield.com/2016/06/working-with-forks-in-go/
`)
}

func showHelp() {
	fmt.Println(`fork
Fork a remote repository to propose a change.

Example
  fork yourID user/repo

Options
    yourID     your username
    user/repo  an identifier like user/rep
`)
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
