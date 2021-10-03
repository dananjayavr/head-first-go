package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alice": 0, "Alma": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	} else {
		fmt.Printf("%s is passing!\n", name)
	}
}

func main() {
	var ranks map[string]float64
	ranks = make(map[string]float64)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks)
	fmt.Println(ranks["bronze"])

	myMap := map[string]float64{"a": 1.2, "b": 5.6}

	fmt.Println(myMap)

	emptyMap := map[string]float64{}
	fmt.Println(emptyMap)

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)

	status("Alice")
	status("Alma")
	status("Bob")

	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}

	fmt.Println(counts)

	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}

	delete(ranks, "bronze")
	rank, ok := ranks["bronze"]
	fmt.Printf("rank: %.1f, ok: %v\n", rank, ok)
}
