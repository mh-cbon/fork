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
  fork -force yourID user/repo
  fork -help

Options
    yourID     your username
    user/repo  an identifier like user/rep

Flags
    -force  Force continue if the directory exists.
    -help   Show this help
    -h      Show this help
```

# Cli examples

```sh
$ fork mh-cbon ScottMansfield/nanolo
forking in /home/mh-cbon/gow/src/github.com/ScottMansfield/nanolog
$ ls -al ~/gow/src/github.com/ScottMansfield/nanolog/
total 88
drwxrwxr-x 6 mh-cbon mh-cbon  4096  3 mai   11:51 .
drwxrwxr-x 3 mh-cbon mh-cbon  4096  3 mai   11:51 ..
drwxrwxr-x 3 mh-cbon mh-cbon  4096  3 mai   11:51 cmd
drwxrwxr-x 8 mh-cbon mh-cbon  4096  3 mai   11:52 .git
-rw-rw-r-- 1 mh-cbon mh-cbon 11357  3 mai   11:51 LICENSE
-rw-rw-r-- 1 mh-cbon mh-cbon 15789  3 mai   11:51 nanolog.go
-rw-rw-r-- 1 mh-cbon mh-cbon 22358  3 mai   11:51 nanolog_test.go
drwxrwxr-x 2 mh-cbon mh-cbon  4096  3 mai   11:51 reader
-rw-rw-r-- 1 mh-cbon mh-cbon  4603  3 mai   11:51 README.md
drwxrwxr-x 3 mh-cbon mh-cbon  4096  3 mai   11:51 test
-rw-rw-r-- 1 mh-cbon mh-cbon   182  3 mai   11:51 .travis.yml
$ cd ~/gow/src/github.com/ScottMansfield/nanolog/ && git remote -v
origin	git@github.com:ScottMansfield/nanolog.git (fetch)
origin	git@github.com:ScottMansfield/nanolog.git (push)
upstream	https://github.com/ScottMansfield/nanolog (fetch)
upstream	https://github.com/ScottMansfield/nanolog (push)

```

# Recipes

#### Update the README

```sh
emd gen -out README.md
```

# Credits

http://blog.sgmansfield.com/2016/06/working-with-forks-in-go/
