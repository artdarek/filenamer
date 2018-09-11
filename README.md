Package filenamer
===================

Package "Filenamer" provides a set of methods for filename manipulation.
It's a very simple API for adding custom prefixes
and suffixes to your base filename such as timestamps, 
random strings etc. Very useful when working with file 
uploads and you need to genereate unique filenames for your uploads.

## Installation
```shell
go get -u github.com/artdarek/filenamer
```

## Examples

```go
package main

import (
	"github.com/artdarek/filenamer"
	"fmt"
)

func main() {
	fn := filenamer.New("test file.jpg")
	fn.CleanIt()
	
	name := fn.Get()
	
	fmt.Println(name)
}
```

```go
package main

import (
	"github.com/artdarek/filenamer"
	"fmt"
)

func main() {
	fn := filenamer.New("test fil e.jpg") // original filename
	
	fn.CleanIt() // replaces whitespaces with _
	fn.WithReplacement("test", "example") // will replace 'test' with 'example'

	fn.WithRandomPrefix(10) // add random charset as a prefix to your file
	fn.WithRandomSuffix(10) // add random charset as a suffix to your file	
	
	fn.SeparateWith("-") // set active separator for all feature modifications
	
	fn.WithTimestamp() // add timestamp prefix
	fn.WithCustomPrefix("image") // add prefix to your filename	
	fn.WithCustomSuffix("renamed") // add suffix to your filename
	fn.WithExtension("png") // change extension of your filename
	fn.WithExtensionRemoved() // removes extension from your filename

	fn.HashIt() // generates md5 hash from generated filename
	
	name := fn.Get() // Get generated filename

	fmt.Println(name)
}
```

See some more examples in `cmd/examples/main.go`. You can test them in action by running following command: 

```shell
go run cmd/examples/main.go
```

## Contributing

Pull requests, bug fixes and issue reports are welcome.

Before proposing a change, please discuss your change by raising an issue.
