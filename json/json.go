package main

import (
	"encoding/json"
	"fmt"
)

type user struct { //decode json ke var
	FullName4 string `json:"Name4"` // Field yang sesuai dengan JSON
	Age4      int    `json:"Age4"`  // Field yang sesuai dengan JSON

	FullName string `json:"Name"`
	Age      int
}

func main() {
	//interface
	var jsonString2 = `{"Name1": "Json", "Age1":27}`
	var jsonData2 = []byte(jsonString2)
	var data1 map[string]interface{}
	json.Unmarshal(jsonData2, &data1) //konversi json dan simpat ke data1 type interface
	fmt.Println("user:", data1["Name1"])
	fmt.Println("user:", data1["Age1"])

	var jsonString3 = `{"Name1": "Son", "Age1":27}`
	var jsonData3 = []byte(jsonString3)
	var data2 interface{}
	json.Unmarshal(jsonData3, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user:", decodedData["Name1"])
	fmt.Println("user:", decodedData["Age1"])

	//json dalam array
	var jsonString4 = `[{"Name4": "Son", "Age4":27}, {"Name4": "Sonah", "Age4":27}]`
	var data4 []user
	var err4 = json.Unmarshal([]byte(jsonString4), &data4)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}
	fmt.Println("4user 1 :", data4[0].FullName4)
	fmt.Println("4user 2 :", data4[1].FullName4)

	//struk
	var jsonString = `{"Name": "Jon", "Age":27}`
	var jsonData = []byte(jsonString)         //mengubah jsonstring ke byte untuk umarshal
	var data user                             // var data dengan tipe user
	var err = json.Unmarshal(jsonData, &data) //konversi json dan simpan di var data
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user:", data.FullName)
	fmt.Println("user:", data.Age)
}
