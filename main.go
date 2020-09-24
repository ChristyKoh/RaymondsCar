package main
// Go Program to store carpool info

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

type Person struct {
	_id string
	Name string
	Music string
	Mbti string
	CanDrive bool
}

type Response struct {
	Status string
	Message string
}

type DataResponse struct {
	Status string
	Message string
	Data Person
}

func isCompletePerson(p *Person) bool {
	return len(p.Name) > 0 && len(p.Music) > 0 && len(p.Mbti) > 0
}

// add rider to a group
func joinCar(w http.ResponseWriter, r *http.Request) {
	// requires parameters name, music, mbti, and group

	canDrive, err := strconv.ParseBool(r.FormValue("canDrive"))
	if err != nil {
		http.Error(w, "bad bool value >:(", 400)
		return
	}

	// extract values into new Person struct
	newPerson := &Person{
		Name: r.FormValue("name"),
		Music: r.FormValue("music"),
		Mbti: r.FormValue("mbti"),
		CanDrive: canDrive,
	}

	var resp DataResponse
	if (isCompletePerson(newPerson) && len(r.FormValue("group")) > 0) {
		addPerson(newPerson, r.FormValue("group"))
		resp = DataResponse{"success", "rider was added to group", *newPerson}
	} else {
		// error if fields not present
		resp.Status = "error"
		resp.Message = "please make sure all fields (name, music, mbti, canDrive, and group) are included"
	}
	
	response_bytes, err := json.Marshal(resp)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response_bytes)
}

// remove rider from group
func dropoff(w http.ResponseWriter, req *http.Request) {

	name := req.FormValue("name")
	group := req.FormValue("group")
	if (len(name) == 0 || len(group) == 0) {
		http.Error(w, "bad input, please make sure name + group parameter is included", 400)
		return
	}
	
	var resp DataResponse
	var deleted Person
	deletePerson(name, group, &deleted)
	if (deleted.Name != "") {
		resp.Status = "success"
		resp.Message = name + " deleted."
		resp.Data = deleted
	} else {
		resp.Status = "error"
		resp.Message = "could not find person " + name + " in group " + group + ". >:("
	}

	response_bytes, err := json.Marshal(resp)
	check(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response_bytes)
}

// get all riders in a group
func getRiders(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data []Person
		Status string
		Message string
	}

	var resp Response

	group := r.FormValue("group")
	if (len(group) == 0) {
		resp.Message = "bad input, please put in the group that you want to get the riders from"
		resp.Status = "failure"
	} else {
		people := readGroup(group)
		if len(people) == 0 {
			resp = Response{people, "error", "group " + group + " does not exist. >:("}
		} else {
			resp = Response{people, "success", "listing riders from group" + group}
		}
	}
	
	bytes, err := json.Marshal(resp)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// delete group
func crashCar(w http.ResponseWriter, r *http.Request) {
	
	var resp Response
	group := r.FormValue("group")
	status_bool := deleteGroup(group)

	if (status_bool) {
		resp.Status = "success"
		resp.Message = "BREAKING NEWS: car crashed :'("
	} else {
		resp.Status = "error"
		resp.Message = "Assassination attempt averted. Could not crash car."
	}
	
	bytes, err := json.Marshal(resp)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {

	deleteGroup("berk")

	daniel := &Person{
		Name: "Daniel",
		Music: "Pop",
		Mbti: "INFP",
		CanDrive: false,
	}

	christy := &Person{
		Name: "Christy",
		Music: "Indie Rock",
		Mbti: "ISFP",
		CanDrive: true ,
	}
	
	people := readGroup("berk")
	fmt.Println(people)
	addPerson(daniel, "berk")
	addPerson(christy, "berk")
	
	http.HandleFunc("/join", joinCar)
	http.HandleFunc("/dropoff", dropoff)
	http.HandleFunc("/riders", getRiders)
	http.HandleFunc("/crash", crashCar)
	http.ListenAndServe(":8090", nil)

}