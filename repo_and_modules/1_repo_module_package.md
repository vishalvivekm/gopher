/*
Module &
Packages
Go code is grouped into packages, and packages are grouped into
modules.
A module specifies the dependencies needed to run our code, including
the Go version and the set of other modules it requires in the go.mod
file

A collection of Go source code becomes a module when there’s a
valid go.mod file in its root directory.
It consists of the module declaration, the minimum compatible version
for Go, and the dependencies for the imported third-party packages.
*/
// go mod init <module_path> // module_path - unique

$ go mod init vishalvivekm/learn
> go: creating new go.mod: module vishalvivekm/learn

$ go mod tidy ( when importing any third party package, running this lists in go.mod file and downloads the package)
// alternatively 
$ go get github.com/

$ go get github.com/<module->
