# xkcd - XKCD library for Go

## Example usage

```go
package main

import "github.com/1nter-p/xkcd"

func main() {
	// Get the comic with number 327 (Exploits of a Mom/"Bobby Tables")
	c := xkcd.ComicFromNum(327)

	// Print the title and alt text
	fmt.Println("Title:", c.Title)
	fmt.Println(c.Alt)
}
```
