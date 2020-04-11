## How to make module

```
$ go mod init github.com/cloudrain21/learn_go_with_test/mytestapp
$ go test ./...
$ ls go.mod go.sum
$ cat go.mod
$ go list -m all
$ go mod tidy
```

Maintain go.mod and go.sum in github.com

```
$ git add go.mod go.sum
$ git commit -m "go module"
$ git push origin master
```
