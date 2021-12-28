package lht

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type SortableTracks struct {
	Tracks []*Track
	By     []string
}

func (st SortableTracks) Len() int {
	return len(st.Tracks)
}

func (st SortableTracks) Less(i, j int) bool {
	for _, b := range st.By {
		if b == "Title" {
			if st.Tracks[i].Title != st.Tracks[j].Title {
				return st.Tracks[i].Title < st.Tracks[j].Title
			}
		} else if b == "Artist" {
			if st.Tracks[i].Artist != st.Tracks[j].Artist {
				return st.Tracks[i].Artist < st.Tracks[j].Artist
			}
		} else if b == "Album" {
			if st.Tracks[i].Album != st.Tracks[j].Album {
				return st.Tracks[i].Album < st.Tracks[j].Album
			}
		} else if b == "Year" {
			if st.Tracks[i].Year != st.Tracks[j].Year {
				return st.Tracks[i].Year < st.Tracks[j].Year
			}
		} else if b == "Length" {
			if st.Tracks[i].Length != st.Tracks[j].Length {
				return st.Tracks[i].Length < st.Tracks[j].Length
			}
		} else {
			return false
		}
	}
	return false
}

func (st SortableTracks) Swap(i, j int) {
	st.Tracks[i], st.Tracks[j] = st.Tracks[j], st.Tracks[i]
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

func Help() {
	fmt.Println("This package is an example to implement a custom sortable struct.")
}
