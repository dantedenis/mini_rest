package web

import (
	"net/http"
	"strconv"
)

func (s *Server) NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("/debug", midddleLogger(http.HandlerFunc(s.debug)))
	router.HandleFunc("/create_user", midddleLogger(http.HandlerFunc(s.create)))
	router.HandleFunc("/get_balance", midddleLogger(http.HandlerFunc(s.balance)))
	router.HandleFunc("/transfer", midddleLogger(http.HandlerFunc(s.transfer)))
	
	return router
}

func (s *Server) balance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "invalid method: " + r.Method, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	user := s.cache.GetUser(id)
	
	if user == nil {
		writeError(w, "user not found", http.StatusBadRequest)
		return
	}
	
	writeResult(w, user)	
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "invalid method: " + r.Method, http.StatusBadRequest)
		return
	}
	
	user, err := s.jsonDecode(r)
	if err != nil {
		writeError(w, "error encoding", http.StatusBadRequest)
		return
	}
	s.cache.Add(user)
}

func (s *Server) transfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "invalid method: " + r.Method, http.StatusBadRequest)
		return
	}
	
	err := s.jsonTrans(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	writeDebug(w, "SUCCESS TRANSFER")	
}
