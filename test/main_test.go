package test

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func testDijkstra() {
	var name = make(map[string]int)
	name["s"] = 0
	name["a"] = 1
	name["b"] = 2
	name["c"] = 3
	name["d"] = 4
	name["e"] = 5
	name["f"] = 6
	name["g"] = 7
	name["h"] = 8
	var index = make(map[int]string)
	index[0] = "s"
	index[1] = "a"
	index[2] = "b"
	index[3] = "c"
	index[4] = "d"
	index[5] = "e"
	index[6] = "f"
	index[7] = "g"
	index[8] = "h"

	var weight [9][9]float64

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i == j {
				weight[i][j] = 0
			} else {
				weight[i][j] = 1000
			}
		}
	}
	weight[name["s"]][name["a"]] = 0.5
	weight[name["s"]][name["b"]] = 0.3
	weight[name["s"]][name["c"]] = 0.2
	weight[name["s"]][name["d"]] = 0.4

	weight[name["a"]][name["e"]] = 0.3

	weight[name["b"]][name["a"]] = 0.2
	weight[name["b"]][name["f"]] = 0.1

	weight[name["c"]][name["f"]] = 0.4
	weight[name["c"]][name["h"]] = 0.8

	weight[name["d"]][name["c"]] = 0.1
	weight[name["d"]][name["h"]] = 0.6

	weight[name["e"]][name["g"]] = 0.1

	weight[name["f"]][name["e"]] = 0.1
	weight[name["f"]][name["h"]] = 0.2

	weight[name["h"]][name["g"]] = 0.4
	//settled := make(map[int]int)
	settledWei := make(map[int]float64)
	mw := make(map[int]float64)
	for i := 0; i < 9; i++ {
		mw[i] = weight[name["s"]][i]
	}
	dijkstra(weight, mw, settledWei)
	//for k, v := range settled {
	//	fmt.Println(k, ":", v)
	//}
	for k, v := range settledWei {
		fmt.Printf("%s:%0.2f ", index[k], v)
	}
	fmt.Println()
}
func dijkstra(weight [9][9]float64, mw map[int]float64, settledWei map[int]float64) {
	//if len(settledWei) >= 9 {
	//	return
	//}
	if len(mw) == 0 {
		return
	}
	minDis := 1000.0
	node := 0
	// 找出距离最小的节点
	for k, v := range mw {
		if v < minDis {
			minDis = v
			node = k
		}
	}
	// update min weight
	for k, v := range mw {
		if v > mw[node]+weight[node][k] {
			mw[k] = mw[node] + weight[node][k]
		}
	}
	settledWei[node] = minDis
	delete(mw, node)

	dijkstra(weight, mw, settledWei)

}
func testopenfile() {
	xx, err := os.OpenFile("./xx", os.O_RDWR|os.O_CREATE, 0666)
	fmt.Println(xx, ":", err)

	flag := syscall.LOCK_SH
	//flag = syscall.LOCK_EX

	// Otherwise attempt to obtain an exclusive lock.
	err = syscall.Flock(int(xx.Fd()), flag|syscall.LOCK_NB)
	fmt.Println(err)

	xx2, err := os.OpenFile("./xx", os.O_RDWR, 0666)
	fmt.Println(xx2, ":", err)
}
func TestXx(t *testing.T) {
	//testDijkstra()
	testopenfile()
}
