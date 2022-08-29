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

// Model for course - file

type Course struct{
	CourseId 	  string   `json:"courseid"`
	CourseName 	  string   `json:"coursename"`
	CoursePrice   int      `json:"price"`
	Author 		  *Author  `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


// fake DB
var courses []Course

// middleware , helper - file 

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main(){

	fmt.Println("API - VivekCodeOnline.in")
	r := mux.NewRouter()

	// seeding
	courses = append(courses , Course{CourseId: "2" , CourseName: "ReactJS" , CoursePrice:  299 , Author: &Author{Fullname: "Vivek Bhawsar" , Website: "Vib.dev"} })

	courses = append(courses , Course{CourseId: "4" , CourseName: "MERN-stack" , CoursePrice:  199 , Author: &Author{Fullname: "Vivek Bhawsar" , Website: "go.dev"} })

	//routing
	r.HandleFunc("/" , serveHome).Methods("GET")
	r.HandleFunc("/courses" , getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}" , getOneCourse).Methods("GET")
	r.HandleFunc("/course" , createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}" , updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}" , deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":3000" , r))


}

func serveHome(w http.ResponseWriter , r *http.Request){

	w.Write([]byte("<h1> Welcome To API by Vivek </h1>"))

}

func getAllCourses(w http.ResponseWriter , r *http.Request){

	fmt.Println(("Get all courses"))
	w.Header().Set("Content-Type " , "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type" , "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop thruough courses , find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}


	json.NewEncoder(w).Encode("NO course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter , r *http.Request){

	fmt.Println("Create one course")
	w.Header().Set("Content-Type" , "application/json")

	//what if body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id , string
	//apend couse into course 

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses , course)
	return

}

func updateOneCourse( w http.ResponseWriter , r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type" , "application/json")


	//first - grab id from req
	params := mux.Vars(r)

	for index , course := range courses {
		if course.CourseId == params["id"]{

			courses = append(courses[:index] , courses[index+1:]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse( w http.ResponseWriter , r *http.Request){

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type" , "application/json")

	// loop ,id , remove (index , index+1)
	params := mux.Vars(r)

	for index , course := range courses {
		if course.CourseId == params["id"]{

			courses = append(courses[:index] , courses[index+1:]... )
			break
		}

	}
}
