# go-incr

go-incr increments semantic version of file.

## Installation

```bash
go build -o $(go env GOPATH)/bin/incr
```

You can export path like this:
```bash:~/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

```bash
$ cat VERSION
1.2.3
$ incr VERSION
before: 1.2.3
after : 1.2.4
$ cat VERSION
1.2.4
$ incr VERSION
before: 1.2.4
after : 1.2.5
$ cat VERSION
1.2.5
```

## Feature work

* Use [suggested regular expression](https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string) including test cases

## Related packages

[This package](https://github.com/blang/semver/blob/master/semver.go) can increment semantic version string