package gokcd_test

import (
	"testing"

	"github.com/1nter-p/gokcd"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	x := gokcd.FromNum(69)

	assert.Equal(t, x.Number, 69)
	assert.Equal(t, x.Title, "Pillow Talk")
	assert.Equal(t, x.SafeTitle, "Pillow Talk")
	assert.Equal(t, x.Alt, "Maybe I should've tried Wexler?")
	assert.Equal(t, x.ImageURL, "https://imgs.xkcd.com/comics/pillow_talk.jpg")
	assert.Equal(t, x.Link, "")
	assert.Equal(t, x.News, "")
	assert.Equal(t, x.Transcript, "Guy: Staring at the ceiling, she asked me what I was thinking about.\nGuy: I should have made something up.\nGuy: The Bellman-Ford algorithm makes terrible pillow talk.\n{{Title Text: Maybe I should've tried Wexler?}}")
	assert.Equal(t, x.Year, "2006")
	assert.Equal(t, x.Month, "2")
	assert.Equal(t, x.Day, "27")
}
