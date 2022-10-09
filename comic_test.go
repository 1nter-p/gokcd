package xkcd_test

import (
	"testing"

	"github.com/1nter-p/xkcd"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	c := xkcd.ComicFromNum(69)

	assert.Equal(t, c.Number,     69)
	assert.Equal(t, c.Title,      "Pillow Talk")
	assert.Equal(t, c.SafeTitle,  "Pillow Talk")
	assert.Equal(t, c.Alt,        "Maybe I should've tried Wexler?")
	assert.Equal(t, c.ImageURL,   "https://imgs.xkcd.com/comics/pillow_talk.jpg")
	assert.Equal(t, c.Link,       "")
	assert.Equal(t, c.News,       "")
	assert.Equal(t, c.Transcript, "Guy: Staring at the ceiling, she asked me what I was thinking about.\nGuy: I should have made something up.\nGuy: The Bellman-Ford algorithm makes terrible pillow talk.\n{{Title Text: Maybe I should've tried Wexler?}}")
	assert.Equal(t, c.Year,       "2006")
	assert.Equal(t, c.Month,      "2")
	assert.Equal(t, c.Day,        "27")
}
