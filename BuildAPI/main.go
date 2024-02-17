package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Course - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database
var courses []Course

// middleware, helper function
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

// Controllers

// Home Route
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Building API in Golang </h1>"))
}

// AllCourses route
func AllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// API controller to GetOneCourse
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The course requested")
	// grab id from request
	params := mux.Vars(r) // user inserted id
	// loop in the courses array and find matching id and return the response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found")

}

// Create Course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course")
	w.Header().Set("Content-Type", "application/json")

	// What if data is nil (empty)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide some data")
	}
	// what if data is {} like this
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found insid JSON")
		return
	}

	// Converting id string into id int
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func main() {
	fmt.Println("Welcome to Building API Section of golang")
}
