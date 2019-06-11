# imports
Command line utility for compiling a list of imports for all Go source files
in a project directory. The utility does not traverse subdirectories and also
expects all source files to be valid Go source.

The output is a JSON object where each key is an import path and each key's
value is an array of source files which import that import path.

## Build
### Clone
```shell
git clone git@github.com:joberly/imports.git
cd imports/imports
go build
```
### Go Get
```shell
go get github.com/joberly/imports/imports
```

## Syntax
```
imports [directory]
```
**Directory** must be either a full directory path or a path relative to the
current directory.

## Example Output

Using the ```testdata``` directory from the imports project itself:

```
$ ./imports ../testdata
{
  "fmt": [
    "a.go",
    "b.go",
    "read_test.go"
  ],
  "github.com/user/pkg": [
    "read_test.go"
  ],
  "github.com/user2/pkg2": [
    "read_test.go"
  ],
  "github.com/user3/pkg3": [
    "read_test.go"
  ],
  "github.com/user4/pkg4": [
    "read_test.go"
  ],
  "github.com/user6/pkg6": [
    "a.go"
  ],
  "github.com/user7/pkg7": [
    "b.go"
  ],
  "io": [
    "read_test.go"
  ],
  "os": [
    "read_test.go"
  ],
  "path": [
    "a.go"
  ],
  "runtime": [
    "b.go",
    "read_test.go"
  ]
}
```
