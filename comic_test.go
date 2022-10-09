package xkcd_test

import (
	"testing"

	"github.com/1nter-p/xkcd"
)

func TestBasic(t *testing.T) {
	_, err := xkcd.ComicFromNum(69)
	if err != nil {
		t.Fail()
	}
}

func TestLatest(t *testing.T) {
	_, err := xkcd.LatestComic()
	if err != nil {
		t.Fail()
	}
}

func TestRandom(t *testing.T) {
	_, err := xkcd.RandomComic()
	if err != nil {
		t.Fail()
	}
}