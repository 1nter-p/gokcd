# xkcd - XKCD library for Go

## Example usage

```go
package main

import "github.com/1nter-p/xkcd"

func main() {
	// Get the comic with number 327 (Exploits of a Mom/"Bobby Tables")
	c, err := xkcd.ComicFromNum(327)
	if err != nil {
		log.Fatal(err)
	}

	// Print the title and alt text
	fmt.Println("Title:", c.Title)
	fmt.Println(c.Alt)
}
```
