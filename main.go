package main
// Go Program to store carpool info

import (
	"fmt"
	"net/http"
	"strconv"
)

type Person struct {
	_id string
	Name string
	Music string
	Mbti string
	CanDrive bool
}

func isCompletePerson(p *Person) bool {
	return len(p.Name) > 0 && len(p.Music) > 0 && len(p.Mbti) > 0
}

// add obj
func joinCar(w http.ResponseWriter, r *http.Request) {
	// requires parameters name, music, mbti, and group

	for key, value := range r.Form {
		fmt.Printf("%s: %s\n", key, value[0])
	}

	canDrive, err := strconv.ParseBool(r.FormValue("canDrive"))
	if err != nil {
		fmt.Printf("bad bool value")
	}

	// extract values into new Person struct
	newPerson := &Person{
		Name: r.FormValue("name"),
		Music: r.FormValue("music"),
		Mbti: r.FormValue("mbti"),
		CanDrive: canDrive,
	}

	if (isCompletePerson(newPerson)) {
		print("adding person!")
		addPersonToGroup(newPerson, r.FormValue("group"))
	} else {
		// error if fields not present
		print("bad input >:(")
		http.Error(w, "bad input, please make sure all fields (name, music, mbti, canDrive) are included", 400)
	}
}

// delete obj
func leaveCar(w http.ResponseWriter, req *http.Request) {

}

// get all objs 
func getRiders(w http.ResponseWriter, r *http.Request) {
	
}

// delete group
func crashCar(w http.ResponseWriter, r *http.Request) {
		
}

func main() {

	daniel := &Person{
		Name: "Daniel",
		Music: "Pop",
		Mbti: "INFP",
		CanDrive: false,
	}

	// christy := &Person{
	// 	name: "Christy",
	// 	music: "Indie Rock",
	// 	mbti: "ISFP",
	// 	canDrive: true ,
	// }
	
	addPersonToGroup(daniel, "berk")
	// deleteGroup("groupShouldBeDeleted")
	
	http.HandleFunc("/join", joinCar)
	http.HandleFunc("/dropoff", leaveCar)
	http.HandleFunc("/riders", getRiders)
	http.HandleFunc("/crash", crashCar)
	http.ListenAndServe(":8090", nil)

}