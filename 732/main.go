package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tal-tech/go-zero/core/hash"
	"github.com/tal-tech/go-zero/core/mathx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	rounds = flag.Int("r", 10000, "rounds")
	subset = flag.Int("s", 30, "subset")
	elements = flag.Int("e", 1000, "elements")
)

func simpleSubset(set []string, sub int) []string {
	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
	if len(set) <= sub {
		return set
	}

	return set[:sub]
}

func complexSubset(set []string, sub int) []string {
	if len(set) <= sub {
		rand.Shuffle(len(set), func(i, j int) {
			set[i], set[j] = set[j], set[i]
		})
		return set
	}

	// group clients into rounds, each round uses the same shuffled list
	count := uint64(len(set) / sub)
	clientID := hash.Hash([]byte(stringx.RandId()))
	round := clientID / count

	r := rand.New(rand.NewSource(int64(round)))
	r.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})

	start := (clientID % count) * uint64(sub)
	return set[start : start+uint64(sub)]
}

func calcEntropy(vals []string, fn func([]string, int) []string) {
	r := make(map[interface{}]int)
	for i := 0; i < *rounds; i++ {
		subs := fn(vals, *subset)
		for _, sub := range subs {
			r[sub]++
		}
	}

	fmt.Println(mathx.CalcEntropy(r))
}

func main() {
	var vals []string
	for i := 0; i < *elements; i++ {
		vals = append(vals, strconv.Itoa(i))
	}

	fmt.Println("original:")
	calcEntropy(vals, simpleSubset)
	fmt.Println("new:")
	calcEntropy(vals, complexSubset)
}
