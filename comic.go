package xkcd

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Comic struct {
	Number     int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	ImageURL   string `json:"img"`
	Link       string `json:"link"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
}

// ComicFromURL creates a new Comic from the given comic URL (info.0.json).
func ComicFromURL(url string) (Comic, error) {
	c := http.Client{Timeout: time.Second * 2}
	
	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Comic{}, err
	}
	
	// Perform the request and save response to res
	res, err := c.Do(req)
	if err != nil {
		return Comic{}, err
	}
	
	// Close the body if it exists
	if res.Body != nil {
		defer res.Body.Close()
	}
	
	// Read the body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Comic{}, err
	}
	
	// Unmarshal the JSON response to a Comic object
	var comic Comic
	err = json.Unmarshal(body, &comic)
	if err != nil {
		return Comic{}, err
	}

	return comic, nil
}

// ComicFromNum creates a new Comic from the given comic number/ID.
func ComicFromNum(num int) (Comic, error) {
	return ComicFromURL("https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json")
}

// LatestComic returns the latest comic as a Comic object.
func LatestComic() (Comic, error) {
	return ComicFromURL("https://xkcd.com/info.0.json")
}

// RandomComic returns a random comic as a Comic object.
func RandomComic() (Comic, error) {
	// Grab the latest comic to find our upper boundary
	l, err := LatestComic()
	if err != nil {
		return Comic{}, err
	}

	// Select a random comic number
	n := rand.Intn(l.Number - 1) + 1

	// Get the comic
	c, err := ComicFromNum(n)
	if err != nil {
		return Comic{}, err
	}

	return c, nil
}
