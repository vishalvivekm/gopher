package main

import "fmt"

func main() {
	/* go maps are implemented by HashTables and provide efficient add, get and delete
	operations.
	// go maps are unordered collections. meaning while priting a map, we may get the key-values pairs of map
	// in different order

	*/
	// declaring a map
	//var codes map[string]string // var <name> map[<datatype-of-key]<value-datatype
	// codes is a nil map at this point i.e. zero value of map in go is nil ( doesn't contain any key)
	// the map is nil, and will cause a panic error if we try to add new keys.
	// codes["en"] = "English" // run-time error : assignment to entry in nil map
	//fmt.Println(codes)

	// to create map with key value pair we have to initialize it first
	// how?
	// using make func // <map_name> := map[<key_data_type>]<value_data_type>{<key-value-pairs>}
	//codes = map[string]string{"Engineering Maths": "BSMA-501", "Operating System": "PCCS-521"}
	// mathsCodes := codes["BSMA-501"]
	// fmt.Println(mathsCodes)
	//fmt.Println(codes)

	// Declaring and initializing go maps using make() which returns a initialized map
	/*
		<name> := make(map[<key_data_type>]<value_data_type>, <initial_capacity-(optional)>)
	*/
	rollNumbers := make(map[string]int)
	fmt.Println(rollNumbers)
	rollNumbers["Bob"] = 2140124
	rollNumbers["James"] = 2140125
	fmt.Println(rollNumbers)
	fmt.Println(len(rollNumbers))

	value, found := rollNumbers["bob"]
	fmt.Printf("%v %t\n", value, found) // 0 false
	value1, found2 := rollNumbers["Bob"]
	fmt.Println(value1, " ", found2) //2140124   true

	// adding new key value pair to a map
	rollNumbers["Sassy"] = 2140105
	rollNumbers["Suna"] = 2140128
	fmt.Println("|-----Studnets ", "and their Roll numbers -------x")

	for key, value := range rollNumbers {
		fmt.Println(key, " -> ", value)
	}
	// update key-value pair
	rollNumbers["James"] = 2140128
	delete(rollNumbers, "James")

	fmt.Println("|---- after deleting James ------>")
	fmt.Println(rollNumbers)

	// how to truncate a map
}

// set a map to an empty map or iteratively delete the key-value pairs in it
