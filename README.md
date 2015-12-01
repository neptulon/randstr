# Random String

[![Build Status](https://travis-ci.org/neptulon/randstr.svg?branch=master)](https://travis-ci.org/neptulon/randstr)
[![GoDoc](https://godoc.org/github.com/neptulon/randstr?status.svg)](https://godoc.org/github.com/neptulon/randstr)

Random string generator of arbitrary size.

## Example

```go
import (
	"github.com/neptulon/randstr"
)

func main() {
	str := randstr.Get(96)
	// str: VkNT!pQXdtHgyffWMIqNZcnOECWhVYYafBGTDjJvE  PlyaWs!UKiKxGQkquNafewfcU ECXgQfYtyZkFIXEJmIYVPRYaIzh
}
```

## License

[MIT](LICENSE)
