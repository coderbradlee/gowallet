package main

import (
	"fmt"
	// "unsafe"
	// "sync/atomic"
	// "time"
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
	var chanArray [10]chan struct{}
	for i := 0; i < 10; i++ {
		chanArray[i] = make(chan struct{}, 0)
	}
	for i := 0; i < 10; i++ {
		go func(j int) {
			chanArray[j] <- struct{}{}
			fmt.Println(j)
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-chanArray[i]
	}
	// var count uint32
	// trigger := func(i uint32, fn func()) {
	// 	for {
	// 		if n := atomic.LoadUint32(&count); n == i {
	// 			fn()
	// 			atomic.AddUint32(&count, 1)
	// 			break
	// 		}
	// 		// time.Sleep(time.Nanosecond)
	// 	}
	// }
	// for i := uint32(0); i < 10; i++ {
	// 	go func(i uint32) {
	// 		fn := func() {
	// 			fmt.Println(i)
	// 		}
	// 		trigger(i, fn)
	// 	}(i)
	// }
	// trigger(10, func() {})
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
