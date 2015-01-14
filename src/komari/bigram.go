package komari

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	scores [][]int
}

func (mat *Matrix) Get(from int, to int) int {
	return mat.scores[from][to]
}
func LoadMatrix(fname string) *Matrix {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(fp))
	if !scanner.Scan() {
		panic("Header not found")
	}
	head := strings.Split(scanner.Text(), " ")
	xsize, _ := strconv.ParseInt(head[0], 10, 32)
	ysize, _ := strconv.ParseInt(head[1], 10, 32)
	mat := make([][]int, xsize)
	for i := range mat {
		mat[i] = make([]int, ysize)
	}

	for scanner.Scan() {
		line := scanner.Text()
		scores := strings.Split(line, " ")
		if len(scores) != 3 {
			log.Fatalf("Invalid format: \"%s\", len: %d", line, len(scores))
		}
		x, _ := strconv.ParseInt(scores[0], 10, 32)
		y, _ := strconv.ParseInt(scores[1], 10, 32)
		s, _ := strconv.ParseInt(scores[2], 10, 32)
		mat[x][y] = int(s)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return &Matrix{mat}
}
