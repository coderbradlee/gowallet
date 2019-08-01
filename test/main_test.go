package test

import (
	//"runtime"

	"os"
	//"os"
	"testing"
)

//func readDir(dir, model string) (uint64, error) {
//	files, err := ioutil.ReadDir(dir)
//	if err != nil {
//		return 0, err
//	}
//	maxN := uint64(0)
//	for _, file := range files {
//		name := file.Name()
//		lens := len(name)
//		if lens < 11 || !strings.Contains(name, model) {
//			continue
//		}
//		num := name[lens-11 : lens-3]
//		n, err := strconv.Atoi(num)
//		if err != nil {
//			continue
//		}
//		if uint64(n) > maxN {
//			maxN = uint64(n)
//		}
//	}
//	return maxN + 1, nil
//}
//func Test(t *testing.T) {
//	x, err := readDir(".", "chain")
//	fmt.Println(x, ":", err)
//	name := "chain" + fmt.Sprintf("-%08d", x) + ".db"
//	fmt.Println(name)
//}
func BenchmarkOsStat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Stat("chain.db")
		f.Size()
	}
}
