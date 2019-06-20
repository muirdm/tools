package index

func wantsInt(int) {}

func _() {
	var (
		aa = "123" //@item(indexAA, "aa", "string", "var")
		ab = 123   //@item(indexAB, "ab", "int", "var")
	)

	var foo [1]int //@item(indexFoo, "foo", "[1]int", "var")
	foo[a]         //@complete("]", indexAB, indexAA)
	foo[:a]        //@complete("]", indexAB, indexAA)
	a[:a]          //@complete("[", indexAA, indexAB)
	a[a]           //@complete("[", indexAA, indexAB)
	wantsInt()     //@snippet(")", indexFoo, "foo[${1}]", "foo[${1:}]")

	var bar map[string]int //@item(indexBar, "bar", "map[string]int", "var")
	bar[a]                 //@complete("]", indexAA, indexAB)
	wantsInt()             //@snippet(")", indexBar, "bar[${1}]", "bar[${1:string}]")

	type myMap map[string]int
	var baz myMap
	baz[a] //@complete("]", indexAA, indexAB)

	var qux []int //@item(indexQux, "qux", "[]int", "var")
	wantsInt()    //@snippet(")", indexQux, "qux[${1}]", "qux[${1:}]")
}
