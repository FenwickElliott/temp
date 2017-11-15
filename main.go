package main

import rest "github.com/FenwickElliott/GoRest"

func main() {
	gh := rest.Rest{"http://localhost:7410/"}
	gh.Get("foo")
}
