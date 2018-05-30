package main

import (
	"fmt"
	"strconv"
	"time"
)

type info struct {
	node string
}

var (
	gInfo  []info
	gCalc  [][]int
	gRelay [][][]info

	l2Cnt = 500
	l1Cnt = 2000
	OsCnt = 50 * 1e4

	defaultTimeout = 5000
)

func preReal() {
	gInfo = make([]info, 0)
	gCalc = make([][]int, 0)
	gRelay = make([][][]info, 0)

	for i := 0; i < l2Cnt; i++ {
		gInfo = append(gInfo, info{node: "L2" + strconv.Itoa(i)})
	}

	for i := 0; i < l2Cnt; i++ {
		initGroup := make([]int, 0)
		for j := 0; j < l2Cnt; j++ {
			init := 0
			initGroup = append(initGroup, init)
		}
		gCalc = append(gCalc, initGroup)
	}

	for i := 0; i < l2Cnt; i++ {
		initGroup := make([][]info, 0)
		for j := 0; j < l2Cnt; j++ {
			init := make([]info, 0)
			initGroup = append(initGroup, init)
		}
		gRelay = append(gRelay, initGroup)
	}
}

func pre() {
	gInfo = make([]info, 0)
	gCalc = make([][]int, 0)
	gRelay = make([][][]info, 0)

	gInfo = append(gInfo, info{node: "L21"})
	gInfo = append(gInfo, info{node: "L22"})
	gInfo = append(gInfo, info{node: "L23"})
	gInfo = append(gInfo, info{node: "L24"})

	gCalc = [][]int{
		/*L21*/ {0, 10, 100, 1},
		/*L22*/ {1, 0, 10, 100},
		/*L23*/ {1, 1, 0, 10},
		/*L24*/ {1, 1, 1, 0},
	}

	for i := 0; i < l2Cnt; i++ {
		initGroup := make([][]info, 0)
		for j := 0; j < l2Cnt; j++ {
			init := make([]info, 0)
			initGroup = append(initGroup, init)
		}
		gRelay = append(gRelay, initGroup)
	}
}

func calcL2() {
	for k := 0; k < l2Cnt; k++ {
		for i := 0; i < l2Cnt; i++ {
			for j := 0; j < l2Cnt; j++ {
				if i == j || j == k || i == k {
					continue
				}
				if 2 <= len(gRelay[i][j]) {
					continue
				}
				if gCalc[i][j] < gCalc[i][k]+gCalc[k][j] {
					continue
				}

				// fmt.Println(i, j, k, gInfo[i], gInfo[k], gInfo[j], gCalc[i][j], gCalc[i][k], gCalc[k][j], gRelay[i][k], gRelay[k][j])
				r1 := gRelay[i][k]
				r2 := gInfo[k]
				r3 := gRelay[k][j]
				rNew := make([]info, 0)
				for _, one := range r1 {
					rNew = append(rNew, one)
				}
				rNew = append(rNew, r2)
				for _, one := range r3 {
					rNew = append(rNew, one)
				}

				gCalc[i][j] = gCalc[i][k] + gCalc[k][j]
				gRelay[i][j] = rNew
			}
		}
	}
	// for i := 0; i < l2Cnt; i++ {
	// 	for j := 0; j < l2Cnt; j++ {
	// 		fmt.Println(gInfo[i], "->", gRelay[i][j], "->", gInfo[j], ":", gCalc[i][j])
	// 	}
	// }
}

func calcL1() {

}

func calcOs() {

}

func calc() {
	tStart := time.Now()
	calcL2()
	fmt.Println("l2", time.Since(tStart))

	tStart = time.Now()
	calcL1()
	fmt.Println("l1", time.Since(tStart))

	tStart = time.Now()
	calcOs()
	fmt.Println("os", time.Since(tStart))
}

func main() {
	// pre()
	preReal()
	calc()
}
