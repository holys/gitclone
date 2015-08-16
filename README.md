# gitclone
A  simple tool let you git clone repo  in the `go get` style.

So, you have to set `GOPATH` first, then `gitclone` shall clone the <repo> to `$GOPATH/src/<HOST>/<AUTHOR>/<REPO>`

## Installation

    go install -o gitclone main.go

## Features

- git clone repo in `go get` style
- authcomplete the missing HTTP Scheme and `.git` suffix.


## Usage

```
$ gitclone https://github.com/holys/gitclone.git
Cloning into '/Users/holys/code/src/github.com/holys/gitclone'...
remote: Counting objects: 5, done.
remote: Compressing objects: 100% (5/5), done.
remote: Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (5/5), done.
Checking connectivity... done.
```

**More friendly cloning**

If you clone like this:

```
$ git clone https://github.com/holys/gitclone  # missing the .git suffix
fatal: destination path 'gitclone' already exists and is not an empty directory.
```

or 

```
$ git clone github.com/holys/gitclone.git   ## missing the http scheme
fatal: repository 'github.com/holys/gitclone.git' does not exist
```

You just got error result. But `gitclone `  allows the above commands.

```
$gitclone github.com/golang/go
Cloning into '/Users/holys/code/src/github.com/golang/go'...
remote: Counting objects: 222826, done.
... ...
```

## Download Binary

https://github.com/holys/gitclone/releases

JUST YET ANOTHER STUPID TOOL.