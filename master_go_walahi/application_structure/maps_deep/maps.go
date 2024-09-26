package main

import (
	"fmt"
	"sort"
)

func main() {
	var firstMap map[string]string
	fmt.Println(firstMap)

	//cant assign value to uninitialized key, so use these:normal way,using make,
	realMap := map[string]string{
		"Ogun":  "APC",
		"Ondo":  "APC",
		"Oyo":   "APC",
		"Lagos": "LP",
	}
	fmt.Println(realMap)
	realMap1 := make(map[string]string)
	realMap1["Ogun"] = "APC"
	realMap1["Ondo"] = "APC"
	realMap1["Oyo"] = "APC"
	realMap1["Lagos"] = "LP"

	fmt.Println(realMap1)
	realMap1["Ondo"] = "Former LP"
	fmt.Println(realMap1)

	v, ok := realMap1["Ondo"]

	if ok {
		fmt.Println("Ondo is an APC state", v)
	} else {
		fmt.Println("Ondo is not an APC state")
	}
	//iterating over a map
	for key, value := range realMap1 {
		fmt.Println(key, value)
	}
	//comparing maps
	//realMap1==realMap1 (can only compare map to nil)

	r := fmt.Sprintf("%s", realMap)
	r1 := fmt.Sprintf("%s", realMap1)
	fmt.Println(r, r1)
	fmt.Println(r == r1)

	realMap2 := realMap1
	realMap2["Oyo"] = "APGA"
	fmt.Println(realMap2, realMap1) //Both changed

	//copy by iterating instead
	realMap3 := make(map[string]string)
	for key, value := range realMap1 {
		realMap3[key] = value
		fmt.Println(realMap3, realMap1)
	}
	realMap3["Ogun"] = "PDP"
	fmt.Printf("realMap1 is still >>>>>>%#v\n\trealMap3 changes to>>>>>%#v\n", realMap1, realMap3)
	delete(realMap3, "Oyo")
	fmt.Println(realMap3)

	//sorting maps by iterating over keys/values
	basket := map[string]int{
		"apple": 7, "mango": 3, "strawberry": 9, "orange": 5,
	}
	keys := make([]string, 0, len(basket)) //using keys
	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
	unSortedMap := map[int]string{
		30: "kim", 50: "jim", 40: "Alex", 20: "diego", 10: "jack"}
	keyS := make([]int, 0, len(unSortedMap))
	for ke := range unSortedMap {
		keyS = append(keyS, ke)
	}
	sort.Ints(keyS)
	for _, ke := range keyS {
		fmt.Println(ke, unSortedMap[ke])
	}
	//clearing out the map using for loop/make
	for index, value := range unSortedMap {
		fmt.Println(index, value)
	}
	for i := range unSortedMap {
		delete(unSortedMap, i)
	}
	fmt.Println(unSortedMap)
	unSortedMap = make(map[int]string)
	fmt.Println(unSortedMap)

	//merge maps
	map1 := map[int]string{10: "Okon", 20: "Abasi", 30: "Uwem"}
	map2 := map[int]string{40: "Sam", 50: "Praise", 60: "Ade", 70: "Bayo"}
	for key, value := range map2 {
		map1[key] = value
	}
	for key, value := range map1 {
		fmt.Println("key: value:\n", key, value)
	}
}
