package main

import "fmt"

/*
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
         result, err := db.Exec(
             "SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
             artist, minYear, maxYear)
         // ...
}
*/
func sqlQuote(x interface{}) string {
	switch x.(type) {
	case bool:
		fmt.Println(x)
	case int, int32, int64:
		//fmt.Println(x + 1) // compile error
		fmt.Println(x)
	}
	return ""
}

func sqlQuote2(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return string(x) // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {

}
