package corehttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

var Codes []string

//temporary function for demo data, get rid of when real database implemented
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func HandleRoutes() {
	Router.HandleFunc("/getCandidates", GetCandidatesHandler).Methods("GET") // located in candidates.go
	Router.HandleFunc("/getElections", GetElectionsHandler).Methods("GET")   // located in elections.go
	Router.HandleFunc("/verifyCode", VerifyEmailHandler).Methods("GET")      // located in verification.go
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")  // located in register_election.go
	http.Handle("/", Router)
}
