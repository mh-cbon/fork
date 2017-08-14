# fork

[![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Package fork helps to fork a repository in your GOPATH.


# TOC
- [Install](#install)
  - [go](#go)
- [Usage](#usage)
  - [$ fork -help](#-fork--help)
- [Cli examples](#cli-examples)
- [Recipes](#recipes)
  - [Update the README](#update-the-readme)
- [Credits](#credits)

# Install

#### go
```sh
go get github.com/mh-cbon/fork
```

# Usage

#### $ fork -help
```sh
fork - HEAD
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
```

# Cli examples

```sh
$ fork mh-cbon kaneshin/pigeon
go [get github.com/kaneshin/pigeon]
git [remote rename origin upstream]
git [remote add origin git@github.com:mh-cbon/pigeon.git]
All done!

git [remote -v]
origin	git@github.com:mh-cbon/pigeon.git (fetch)
origin	git@github.com:mh-cbon/pigeon.git (push)
upstream	https://github.com/kaneshin/pigeon (fetch)
upstream	https://github.com/kaneshin/pigeon (push)

The fork is available at
  .../src/github.com/kaneshin/pigeon

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

$ ls -al $GOPATH/src/github.com/kaneshin/pigeon/
total 88
drwxrwxr-x 6 mh-cbon mh-cbon  4096  3 mai   11:51 .
drwxrwxr-x 3 mh-cbon mh-cbon  4096  3 mai   11:51 ..
drwxrwxr-x 8 mh-cbon mh-cbon  4096  3 mai   11:52 .git
...

$ cd $GOPATH/src/github.com/kaneshin/pigeon/ && git remote -v
origin	git@github.com:mh-cbon/nanolog.git (fetch)
origin	git@github.com:mh-cbon/nanolog.git (push)
upstream	https://github.com/kaneshin/pigeon (fetch)
upstream	https://github.com/kaneshin/pigeon (push)

```

# Recipes

#### Update the README

```sh
emd gen -out README.md
```

# Credits

http://blog.sgmansfield.com/2016/06/working-with-forks-in-go/
