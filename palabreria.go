package main
import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDict(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(dat),"\n")
}

func getWord(words []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return words[rand.Intn(len(words))]
}

func fetchDicts(root string) [][]string {
	var dicts [][]string
	dictNames, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(dictNames); i++ {
		dictPath := fmt.Sprintf("%s/%s", root,  dictNames[i].Name())
		dicts = append(dicts, readDict(dictPath))
	}
	return dicts
}

func getPhrase(words int, separator string) string {
	var output []string
	dicts := fetchDicts("dict")
	for i := 0; i < words; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		dict := dicts[rand.Intn(len(dicts))]
		output = append(output, getWord(dict))
	}
	return strings.Join(output, separator)
}


func main() {
	wordCountPtr := flag.Int("words", 3, "Number of words")
	separatorPtr := flag.String("separator", "-", "Separator used to join words")
	flag.Parse()
	phrase := getPhrase(*wordCountPtr, *separatorPtr)
	fmt.Print(phrase)
}
