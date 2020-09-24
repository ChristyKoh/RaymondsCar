package main


import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
	return !info.IsDir()
}

// returns array of person structs, otherwise empty array
func readGroup(group string) []Person {
	var people []Person

	if !fileExists(group) {
		return people
	}

	// open file
	bytes_read, err := ioutil.ReadFile(group)
	if (err != nil) {
		fmt.Println("0 bytes read")
		return people
	}

	// unmarshal data
	err = json.Unmarshal(bytes_read, &people)
	check(err)
	fmt.Printf("%+v", people)	
	
	return people
}


// write person to group file
func addPerson(person *Person, group string) string {

	var people []Person
	
	// check if file exists
	if fileExists(group) {
		people = readGroup(group)
	}
	
	people = append(people, *person)			// add person to Person array
	personBytes, err := json.Marshal(people)	// convert to JSON
	check(err)
	
	// write to group file
	f, err := os.OpenFile(group, os.O_WRONLY|os.O_CREATE, 0755)
	check(err)
	defer f.Close()
	if _, err := f.Write(personBytes); err != nil {
		fmt.Println(err)
	}

	return "success writing person to group"
}

// deletes person from group, returns number of people left.
func deletePerson(name string, group string, deleted *Person) {

	var people []Person
	var remaining []Person
	// var deleted *Person

	// check if file exists
	if fileExists(group) {
		people = readGroup(group)
	} else {
		fmt.Println("Group did not exist!")
		return
	}

	// copy all people except person to be deleted to new array
	for _, person := range people {
		if (person.Name != name) {
			remaining = append(remaining, person)
		} else {
			*deleted = person
			// fmt.Println("\n person deleted: \n", *deleted)
			continue
		}
	}

	fmt.Println("\n remaining people \n", remaining)
	// fmt.Println("\n person deleted: \n", *deleted)

	if deleted == nil {
		fmt.Println("person doesn't seem to exist in group")
		return
	} else if len(remaining) == 0 {
		fmt.Println("Everyone has been deleted from this group. Will remove group.")
		deleteGroup(group)
		return
	}

	personBytes, err := json.Marshal(remaining)
	check(err)

	// overwrite group file
	f, err := os.OpenFile(group, os.O_TRUNC|os.O_WRONLY, 0755)
	check(err)
	defer f.Close()
	
	if _, err := f.Write(personBytes); err != nil {
		fmt.Println(err)
	}
}

// delete file corresponding to group
func deleteGroup(group string) bool {
	
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
