package main

import (
	"fmt"
	_ "time"
	"time"
	"math/rand"
)

const (
	VEDIO_IDX    = iota //0
	SPEAKER_IDX         //1
	ARLPIC_IDX          //2
	NEWSP_IDX           //39
	BBSTOPIC_IDX        //4
	PRDTLIB_IDX         //5
)

func main() {
	//maxUseTime := 3600 * 24
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Float64()
	fmt.Println(randNum)
}

func test(m map[int]int) {
	m[3] += 1
}
