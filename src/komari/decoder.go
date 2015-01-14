package komari

import (
	"github.com/tchap/go-patricia/patricia"
	//"log"
)

type Decoder struct {
	dict *Dictionary
	mat  *Matrix
}

func NewDecoder(dict *Dictionary, mat *Matrix) *Decoder {
	return &Decoder{dict, mat}
}

type Score struct {
	score int
	back  int
	rcat  int
	word  *Word
}

func (dec *Decoder) Search(sentence string) []*Word {
	runes := []rune(Reverse("$" + sentence + "$"))
	tbl := make([]Score, len(runes))
	for i := range tbl {
		tbl[i].score = 1000000
		tbl[i].back = -1
		tbl[i].rcat = 0
	}
	tbl[0].score = 0
	tbl[0].word = dec.dict.eos
	for i := 1; i < len(tbl); i++ {
		runeIdx := len(tbl) - i - 1
		cell := &tbl[i]
		for slen := 1; slen <= i; slen++ {
			sub := runes[runeIdx : runeIdx+slen]
			suffix := string(sub)
			item := dec.dict.tree.Get(patricia.Prefix(suffix))
			if item == nil {
				continue
			}
			word := item.(*Word)
			subInd := i - slen
			subCell := &tbl[subInd]
			score := subCell.score + word.score + dec.mat.Get(subCell.rcat, word.lCat)
			if score < cell.score {
				//log.Printf("%s\n", word.repr);
				cell.score = score
				cell.back = subInd
				cell.rcat = word.rCat
				cell.word = word
			}
		}
	}
	lst := make([]*Word, 0)
	for pt := &tbl[len(tbl)-1]; pt.back >= 0; pt = &tbl[pt.back] {
		lst = append(lst, pt.word)
	}
	r := make([]*Word, 0)
	for i := len(lst) - 1; i >= 0; i-- {
		r = append(r, lst[i])
	}
	return r
}
