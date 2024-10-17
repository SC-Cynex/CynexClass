package router

import (
	"github.com/SC-Cynex/cynex-class-service/internal/api/teachers"
	"github.com/gorilla/mux"
)

func NewRouter(teacherHandler *teachers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/teachers", teacherHandler.CreateTeacher).Methods("POST")
	r.HandleFunc("/teachers/{id}", teacherHandler.GetTeacher).Methods("GET")

	return r
}
