package main

import "fmt"

func main() {
	// ages := map[string]int{
	// 	"Devita": 30,
	// 	"Aditya": 32,
	// }
	// for key, value := range ages {
	// 	fmt.Println("Key:", key, "Value:", value)
	// 	fmt.Println("---------------")
	// }
	// fmt.Println("---------------")

	// fmt.Println("Initial map:", ages)
	// fmt.Println("---------------")

	// ages["Rama"] = 40
	// fmt.Println("After adding Rama", ages)
	// fmt.Println("---------------")

	// fmt.Println("Age of Rama:", ages["Rama"])
	// fmt.Println("---------------")

	// ages["Rama"] = 45
	// fmt.Println("After update Rama:", ages["Rama"])
	// fmt.Println("---------------")

	// delete(ages, "Aditya")
	// fmt.Println("After deleting Aditya:", ages)
	// fmt.Println("---------------")

	// age, exists := ages["Aditya"]
	// if exists {
	// 	fmt.Println("Aditya's age exists and is:", age)
	// } else {
	// 	fmt.Println("Aditya not found")
	// }
	// fmt.Println("---------------")
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)
	fmt.Println(a["brand"])
	fmt.Println("---------------")

}
