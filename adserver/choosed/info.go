package main 

import (
	"fmt"
	"math/rand"
)

type Info struct {
	adBlockList AdBlockList
}

func NewInfo() *Info {
	return &Info{adBlockList : make(AdBlockList)}
}

func (info *Info) findAdBlock(id int) (adBlock *AdBlock) {
	block, ok := info.adBlockList[id]
	if !ok {
		fmt.Println("Cannot find block with id: ", id)
		return nil
	}

	return &block
}

func chooseBanner(block *AdBlock) (*Banner) {
	count := len(block.bannerList)

	if count == 0 {
		fmt.Println("There are no banner")
		return nil
	}

	//-->
	/*max := count*10
	res := make([]int, max)

	for i := 0; i < 300; i++ {
		current := (rand.Intn(max) + rand.Intn(max))%max
		res[current] = res[current] + 1
	}

	for i := 0; i < max; i++ {
		fmt.Println(res[i])

	return nil
	}*/

	//массив id-шников
	ids := make([]int, count)
	i := 0
	for _, v := range block.bannerList {
		ids[i] = v.id
		i++
		//fmt.Println("k=", k, " v=", v)
	}

	current := (rand.Intn(count) + rand.Intn(count))%count
	fmt.Println(current)

	banner, ok := block.bannerList[ ids[current] ]
	if !ok {
		fmt.Println("nil")
		return nil
	}

	return &banner
}