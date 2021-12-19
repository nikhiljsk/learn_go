package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"` // - won't return this value in Response
	Author      *Author `json:"author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Simulating Database
var courses []Course

// Middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Starting up the API - Nikhil JSK")
	r := mux.NewRouter()

	// Seeding data
	seedData := []Course{
		{
			CourseId:    "412",
			CourseName:  "Golang!",
			CoursePrice: 129,
			Author: &Author{
				Fullname: "Nikhil JSK",
				Website:  "nikhiljsk.com",
			},
		},
		{
			CourseId:    "4112",
			CourseName:  "Python!",
			CoursePrice: 119,
			Author: &Author{
				Fullname: "Nikhil JSK",
				Website:  "nikhiljsk.com",
			},
		},
	}
	courses = append(courses, seedData...)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/delete", deleteAllCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":8089", r))

}

// Controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Nikhil JSK here!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all the courses!")
	// Adding some errors in the response to send
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching one course!")
	w.Header().Set("Content-Type", "application/json")

	// Get user sent parameters
	params := mux.Vars(r)
	fmt.Printf("Here are the params: %v and its type %T", params, params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found for the given id" + string(params["id"]))
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading one course!")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No body found to upload!")
		return
	}

	// Check if the sent data is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Found {} empty data in body")
		return
	}

	// Body OK. Generate id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(131313))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Find the course
	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Done updated the course:" + params["id"])
			return
		}
	}
	json.NewEncoder(w).Encode("Did not find the requested course to update" + params["id"])

	// Update the course

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("Deleted the course:" + params["id"])
			return
		}
	}
	json.NewEncoder(w).Encode("Did not find the requested course to delete" + params["id"])
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all courses !!!!!")
	w.Header().Set("Content-Type", "application/json")

	courses = []Course{}
	json.NewEncoder(w).Encode("Deleted all. See ya :(")
}
