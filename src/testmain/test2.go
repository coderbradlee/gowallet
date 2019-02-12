package main

import (
	
)
// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string
	phylum string
	class  string
	order  string
	family string
	genus  string
	species string
	
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}
func test(){
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)
	fmt.Println(category)
}
