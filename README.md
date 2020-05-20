# spamcheck

spamcheck uses Postmark's spam API. The API is free to use. No API key is
required. For more information visit https://spamcheck.postmarkapp.com/.

```go
package main

import (
	"fmt"

	"github.com/golevi/spamcheck"
)

func main() {
    scr := spamcheck.NewRequest("I am a nigerian prince and will give you $1 million")
    scr = spamcheck.Short

	resp, err := scr.CheckScore()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(resp)

    // {true 7.9 [] }
}
```
