# spamcheck

spamcheck uses Postmark's spam API. The API is free to use. No API key is
required. For more information visit https://spamcheck.postmarkapp.com/.

```go
package main

import (
	"fmt"
	"log"

	"github.com/golevi/spamcheck"
)

const msg string = `hello i give you money`

func main() {
	req := &spamcheck.Request{
		Email:   msg,
		Options: spamcheck.Long,
		// Options: spamcheck.Short,
	}

	resp, err := req.Process()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}
```
