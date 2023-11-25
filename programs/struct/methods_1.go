// methods are functions with defined receiver

package main

import "fmt"

type Circle struct {
	radius    float32
	area      float32
	perimeter float32
}
type friends struct {
	name1 string
	name2 string
	name3 string
}

func (cir *Circle) calcArea() float32 {
	area := 3.14 * cir.radius * cir.radius
	cir.area = area
	return area

}
func main() {
	cir1 := Circle{
		radius: 2.50,
	}
	fmt.Println(cir1.calcArea())
	fmt.Println(cir1)

	friends_list := friends{name1: "vivek"} // comma is not necessary here
	//ojh := friends{
	//	name1: "vivek",
	//}
	//fmt.Println(ojh)
	/*
		friends{
		name: "vivek", // comma is necessary here
		}
	*/
	fmt.Println(friends_list)         //{vivek  }
	fmt.Printf("%+v\n", friends_list) //{name1:vivek name2: name3:}
}
