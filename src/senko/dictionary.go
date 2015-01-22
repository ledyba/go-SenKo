package senko

import (
	"encoding/csv"
	"github.com/tchap/go-patricia/patricia"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Dictionary struct {
	tree *patricia.Trie
	eos  *Word
}

type Word struct {
	lCat  int
	rCat  int
	score int
	word  string
	repr  string
}

func NewDictionary() *Dictionary {
	return &Dictionary{patricia.NewTrie(), &Word{0, 0, 0, "$", "EOS"}}
}

func (word *Word) Repr() string {
	return word.repr
}
func (word *Word) Word() string {
	return word.word
}

func (dic *Dictionary) LoadDictionaries(dname string) {
	log.Printf("Loading: \"%s\"\n", dname)
	abs, _ := filepath.Abs(dname)
	infos, _ := ioutil.ReadDir(abs)
	for _, info := range infos {
		fpath := dname + "/" + info.Name()
		if info.Mode().IsRegular() {
			if strings.HasSuffix(info.Name(), ".csv.utf8") {
				dic.loadDictionary(fpath)
			}
		} else if info.Mode().IsDir() {
			dic.LoadDictionaries(fpath)
		}
	}
	dic.tree.Insert(patricia.Prefix("$"), dic.eos)
	log.Printf("Loaded: \"%s\"\n", dname)
}

func (dic *Dictionary) loadDictionary(fname string) {
	var err error = nil
	log.Printf("Loading: %s\n", fname)
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		panic(err)
	}
	f := csv.NewReader(fp)
	cnt := 0
	var line []string
	for line, err = f.Read(); err == nil; line, err = f.Read() {
		if len(line) == 1 && line[0] == "#" {
			continue
		}
		if len(line) < 4 {
			log.Fatalf("Invalid format: \"%v\", len: %d", line, len(line))
		}
		w := line[0]
		l, _ := strconv.ParseInt(line[1], 10, 32)
		r, _ := strconv.ParseInt(line[2], 10, 32)
		s, _ := strconv.ParseInt(line[3], 10, 32)
		//log.Printf("%s: %d %d %d\n",w,l,r,s);
		dic.tree.Insert(patricia.Prefix(Reverse(w)), &Word{int(l), int(r), int(s), w, strings.Join(line, ",")})
		cnt++
	}
	if err != nil && err != io.EOF {
		panic(err)
	}
	log.Printf("Loaded: %s(%d)\n", fname, cnt)
}
