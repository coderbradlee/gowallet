package main

import (
	"fmt"
	// "unsafe"
)

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	fmt.Println("call setname ")
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}
func NewCat(name string) Cat {
	return Cat{name: name}
}

func test() {
	num := 1
	sign := make(chan struct{}, num)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < 10; j++ {
		<-sign
	}

	// NewCat("little pig").SetName("monster")
	// dog := Cat{name: "little pig"}
	// dogP := &dog
	// dogPtr := uintptr(unsafe.Pointer(dogP))
	// namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	// nameP := (*string)(unsafe.Pointer(namePtr))
	// fmt.Println(*nameP)
	// cat := New("little pig", "American Shorthair", "cat")
	// cat.SetName("monster") // (&cat).SetName("monster")
	// fmt.Printf("The cat: %s\n", cat)

	// // cat.SetNameOfCopy("little pig")
	// // fmt.Printf("The cat: %s\n", cat)

	// type Pet interface {
	// 	// SetName(name string)
	// 	Name() string
	// 	Category() string
	// 	ScientificName() string
	// }

	// // _, ok := interface{}(cat).(Pet)
	// // fmt.Printf("Cat implements interface Pet: %v\n", ok)
	// // cat.SetName("afafdagaga")
	// // _, ok = interface{}(&cat).(Pet)
	// // fmt.Printf("*Cat implements interface Pet: %v\n", ok)

	// var pet = cat
	// cat.SetName("aafadgagdafdadfafdadfafdafdafdafafa") // (&cat).SetName("monster")
	// fmt.Printf("The cat: %s\n", cat)
	// fmt.Printf("The cat: %s\n", pet)
}
