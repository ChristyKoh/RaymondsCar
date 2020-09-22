package main


import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"log"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
	return !info.IsDir()
}

func readGroup(group string) []Person {
	var people []Person

	// open file
	f, err := os.OpenFile(group,
		os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// read line by line
	// if _, err := f.ReadString(string(personBytes) + "\n"); err != nil {
	// 	fmt.Println(err)
	// }
	
	// var nextPerson *[]byte
	// var counint
	// for b, err := f.ReadAt(&nextPerson + , off int64)
	// 
	//json.Unmarshal()

	return people
}

func deleteObj(group string) {
	// people := readGroup(group)
	// wait lol ok ill just let u do it
	// var filtered_people [len(people)]Person
	// filtered_people = make([]Person, len(people))
	// i := 0
	
}

func deleteGroup(group string) bool {
	// delete file
	
	if !fileExists(group) {
		return false // file never existed to begin with
	}

    var err = os.Remove(group)
    if err != nil {
		fmt.Println("error deleting")
		panic(err)
	}

	fmt.Println("File Deleted")
	return true // file existed, and was deleted
}

func addPersonToGroup(person *Person, group string) string {

	// var people_already_in_group = "[]"
	if fileExists(group) {
        fmt.Println("group file exists")
    } else {
		fmt.Println("file with name %s does not exist  (or is a directory). Will be writing.", group)
		err := ioutil.WriteFile(group, []byte(""), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
	
	personBytes, err := json.Marshal(person)
	// os.Stdout.Write(personBytes)
	fmt.Println("personBytes as string: ", string(personBytes))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// append to group file
	f, err := os.OpenFile(group,
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(personBytes) + "\n"); err != nil {
		fmt.Println(err)
	}	
	

	return "success writing person to group"
}