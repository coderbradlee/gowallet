package test

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func combineCollect(collect *[][]int, result []int, all []int, m int) {
	if len(result) >= m {
		ret := make([]int, len(result))
		copy(ret, result)
		temp := *collect
		temp = append(temp, ret)
		*collect = temp
		//count++
		return
	}

	var index int
	if len(result)-1 < 0 {
		index = -1
	} else {
		lastElems := result[len(result)-1]
		for i, v := range all {
			if lastElems == v {
				index = i
			}
		}
	}

	if index >= len(all)-1 {
		return
	}
	//fmt.Println("index:", index)
	temp := all[index+1:]
	for _, v := range temp {
		//fmt.Println("v:", v)
		resultC := make([]int, len(result))
		copy(resultC, result)
		resultC = append(resultC, v)

		combineCollect(collect, resultC, all, m)
	}

}
func combine2(people []int, first, second, third int) {
	sum := first + second + third
	result := make([]int, 0)
	sumResult := make([][]int, 0)
	combineCollect(&sumResult, result, people, sum)
	fmt.Println(len(sumResult))
	for _, v := range sumResult {
		//fmt.Println(v)
		sumR1 := make([][]int, 0)
		result1 := make([]int, 0)
		combineCollect(&sumR1, result1, v, first)
		for _, v1 := range sumR1 {
			result2 := make([]int, 0)
			a := vSubv1(v, v1)
			sumR2 := make([][]int, 0)
			combineCollect(&sumR2, result2, a, second)
			for _, v2 := range sumR2 {
				if len(v2) < second {
					continue
				}
				b := vSubv1(v2, v1)
				if len(b) < second {
					continue
				}
				v1plusv2 := append(v1, v2...)
				//fmt.Println("v1plusv2:", v1plusv2)
				bb := vSubv1(v, v1plusv2)
				newRet := append(v1plusv2, bb...)
				fmt.Println(newRet)
				count++
			}
		}
	}
}
func vSubv1(all []int, except []int) []int {
	newRet := make([]int, 0)
	for _, v := range all {
		exist := false
		for _, e := range except {
			if v == e {
				exist = true
			}
		}
		if !exist {
			newRet = append(newRet, v)
		}
	}
	return newRet
}

var count int

func combineCollect2(collect *[][]int, result []int, all []int, m int) {
	if len(result) >= m {
		ret := make([]int, len(result))
		copy(ret, result)
		temp := *collect
		temp = append(temp, ret)
		*collect = temp
		//count++
		return
	}
	var index int
	if len(result)-1 < 0 {
		index = -1
	} else {
		lastElems := result[len(result)-1]
		for i, v := range all {
			if lastElems == v {
				index = i
			}
		}
	}

	if index >= len(all)-1 {
		return
	}
	//fmt.Println("index:", index)
	temp := all[index+1:]
	for _, v := range temp {
		//fmt.Println("v:", v)
		resultC := make([]int, len(result))
		copy(resultC, result)
		resultC = append(resultC, v)

		combineCollect(collect, resultC, all, m)
	}

}
func min(x, y, z int) (m int) {
	m = x
	if m > y {
		m = y
	}
	if m > z {
		m = z
	}
	return
}
func distance(a, b byte) int {
	if a == b {
		return 0
	}
	return 1
}
func constructMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}
func dynamic() {
	a := "mouse"
	b := "mousse"
	row := len(b)
	column := len(a)
	m := constructMatrix(row, column)
	m[0][0] = 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 && j > 0 {
				m[i][j] = j
			}
			if i > 0 && j == 0 {
				m[i][j] = i
			}
			if i > 0 && j > 0 {
				m[i][j] = min(m[i-1][j-1]+distance(a[j], b[i]), m[i-1][j]+1, m[i][j-1]+1)
			}
			fmt.Printf("%d ", m[i][j])
			if j >= column-1 {
				fmt.Printf("\n")
			}
		}
	}
}
func TestJson(t *testing.T) {
	data := struct {
		Xx [][]byte
	}{[][]byte{{1}}}

	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
func TestXx(t *testing.T) {
	data := "0000000000000000000000006356908ace09268130dee2b7de643314bbeb36830000000000000000000000000000000000000000000000000000000000000004"
	hb, _ := hex.DecodeString(data)
	out := crypto.Keccak256(hb)
	fmt.Println(hex.EncodeToString(out))

	data = "0000000000000000000000000ddfc506136fb7c050cc2e9511eccd81b15e74260000000000000000000000000000000000000000000000000000000000000004"
	hb, _ = hex.DecodeString(data)
	out = crypto.Keccak256(hb)
	fmt.Println(hex.EncodeToString(out))
	encodeString2 := base64.StdEncoding.EncodeToString(out)
	fmt.Println(encodeString2)

	//input := []byte("b4ecb247a6d3d67eed0ceddcfdc31b865baf1deba46a5c92b8adb649df8bba50")
	input, _ := hex.DecodeString("a4c248d840bab7ecdbde059148a289021ce154661b3107da67474e2bfa4238e8")
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	decodeBytes, err := base64.StdEncoding.DecodeString("pMJI2EC6t+zb3gWRSKKJAhzhVGYbMQfaZ0dOK/pCOOg=")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodeBytes))
	x, _ := big.NewInt(0).SetString(string(decodeBytes), 16)
	fmt.Println(x.Text(10))
	//dynamic()
	//fmt.Println(min(1, 2, 3))
	//testall := []int{1, 2, 3, 4, 5}
	//exc := []int{2, 4}
	//ret := vSubv1(testall, exc)
	//fmt.Println(ret)
	//people := make([]int, 0)
	//for i := 0; i < 10; i++ {
	//	people = append(people, i)
	//}
	//first := 1
	//second := 2
	//third := 3
	////combine2(people, first, second, third)
	//collect := make([][]int, 0)
	//result := make([]int, 0)
	//combineCollect2(&collect, result, people, first)
	//combineCollect2(&collect, result, people, second)
	//combineCollect2(&collect, result, people, third)
	//
	//fmt.Println(count)
	//for _, v := range collect {
	//	fmt.Println(v)
	//}
}
func TestSliceof(t *testing.T) {
	//var s [][]string
	//for i := 0; i < 10; i++ {
	//	sl := make([]string, 0, 5)
	//	for j := 0; j < 5; j++ {
	//		sl = append(sl, "a")
	//	}
	//	s = append(s, sl)
	//}
	//fmt.Println(s)
	//xx(s)
	//fmt.Println(s)
	//s = append(s, []string{"zzzz", "zzzz", "zzzz", "zzzz", "zzzz"})
	//fmt.Println(s)

	xx(nil)
}
func xx(input [][]string) {
	slice := make([]int, 2, 30)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	fmt.Printf("slice %v %p \n", slice, slice)

	ret := changeSlice(&slice)
	fmt.Printf("slice %v %p, ret %v %p\n", slice, slice, ret, ret)

	ret[1] = -1111

	fmt.Printf("slice %v %p, ret %v %p\n", slice, slice, ret, ret)
}
func changeSlice(s *[]int) []int {
	ss := *s
	ss = append(ss, 3)
	*s = ss
	return ss
}
