package main

import (
	"io"
	"time"
)

// a program that organizes or sells digitized cultural artifacts like music, films and books.
// Entity: Album, Book, Movie, Magazine, Podcast, Music, TVEpisode, Track

// Artifact some attributes are common to all artifacts
type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

// Text for Book, Magazine
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

// Audio for Podcast, Music
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

// Video for Movie, TVEpisode
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

// Streamer for video and audio
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

func main() {

}
