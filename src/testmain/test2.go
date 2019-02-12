package main

import (
	"fmt"
)

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	Kingdom string
	phylum  string
	class   string
	order   string
	family  string
	genus   string
	species string
}

func (ac AnimalCategory) String() string {
	ac.Kingdom = "kindowmmmmmmmm"
	return fmt.Sprintf("sdfasfd %s%s%s%s%s%s%s",
		ac.Kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string
	AnimalCategory
}

func test() {
	category := AnimalCategory{species: "cat"}

	an := &Animal{scientificName: "statw", AnimalCategory: category}
	fmt.Printf("The animal category: %s\n", an)
	fmt.Println(an.String())
	fmt.Println(an.Kingdom)
	fmt.Println(an.species)
}
