# go-pdf2svg

go-pdf2svg is a Go package for convert PDF to SVG and use for Web and Easy to use it.
SVG is best for high quality for web.

## Installation

Use the apt for install depedencies

This package required `inkscape`

```bash
$ apt install inkscape
```

## Usage

```go
package main

import (
	"log"
	c "github.com/fajrulaulia/go-frompdf"
)

func main() {
	err := c.Exporter("DirOfFilePDF.pdf", "newSvgFile")
	if err != nil {
		log.Fatal(err)
		return
	}
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Author
[FAJRUL AULUA](https://twitter.com/fajrulgopher)


