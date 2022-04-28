package main

import (
	"fmt"
	"os"
	"sort"
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

// It will run faster if each element is a pointer because the sort function may swap many pairs of elements.
var tracks = []*Track{
	&Track{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	&Track{"Go", "Moby", "Moby", 1992, length("3m37s")},
	&Track{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	&Track{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	// *tabwriter.Writer satisfies io.Writer
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// To sort the play list by the Artist field, we define a new slice type with the necessary Len, Less, and Swap methods
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// To sort the play list by the Year field, we define a new slice type with the necessary Len, Less, and Swap methods
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// ...

// sort.Interface
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	// print original tracks
	printTracks(tracks)
	// first click on Artist column, so sort tracks by Artist ascending
	sort.Sort(byArtist(tracks))
	fmt.Println("==========================================================")
	printTracks(tracks)
	fmt.Println("==========================================================")
	// click on Artist column again, so sort tracks by Artist descending
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println("==========================================================")
	// click on Year column, so sort tracks by Year ascending
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	fmt.Println("==========================================================")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
	fmt.Println("==========================================================")
	values := []int{3, 1, 4, 1}
	// []int is wrapped as IntSlice in sort package
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"

	values2 := []float32{6.8, 5.4, 3.2}
	values3 := FuncVar(values2)
	sort.Sort(sort.Float64Slice(values3))
	fmt.Println(values3)
}

func FuncVar(f32 []float32) []float64 {
	f64 := make([]float64, len(f32))
	var f float32
	var i int
	for i, f = range f32 {
		f64[i] = float64(f)
	}
	return f64
}
