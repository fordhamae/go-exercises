package main

import (
	"fmt"
	"sort"
)

type numbers sort.IntSlice

func (p numbers) Len() int {
	return len(p)
}

func (p numbers) Swap(i, j int){
	p[i],p[j]=p[j],p[i]
}

func (p numbers) Less(i, j int) bool {
	return p[i] < p[j]
}

func main(){
	s := []int{7,4,8,2,9,19,12,32,3}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	i := numbers{7,4,8,2,9,19,12,32,3}
	fmt.Println(i)
	sort.Sort(i)
	fmt.Println(i)
}