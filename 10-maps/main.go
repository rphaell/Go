package main

import "fmt"

func main() {
	fmt.Println("welcome to mapsin go! =)")

	//criando maps
	Languages := make(map[string]string)

	Languages["JS"] = "Javascript"
	Languages["RB"] = "Ruby"
	Languages["Po"] = "Pytn"
	Languages["PY"] = "Pytn"
	Languages["PY"] = "Pton"
	Languages["PY"] = "Pyton"
	Languages["PY"] = "Pthon"

	//pegando apenas um valor da key entre os colchetes
	fmt.Println(Languages)
	fmt.Println(Languages["RB"])

	//delete dados from maps usando a key
	delete(Languages, "JS")
	fmt.Println(Languages)

	//loops are interesting in go

	for key, value := range Languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
