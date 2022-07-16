# go-pdf2svg
[![Go](https://github.com/fajrulaulia/go-pdf2svg/actions/workflows/go.yml/badge.svg)](https://github.com/fajrulaulia/go-pdf2svg/actions/workflows/go.yml)

go-pdf2svg is a Go package for convert PDF to SVG and use for Web and Easy to use it.
SVG is best for high quality for web.


## Installation

This package required `inkscape`
```bash
# apt install inkscape # debian, Ubuntu
# dnf install inkscape # Fedora, Redhat

```

## Usage

```go
package main

import (
	"log"
	c "github.com/fajrulaulia/go-pdf2svg"
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


