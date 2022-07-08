package web

import (
	"net/http"
	"log"
	"encoding/json"
	"wb_test/pkg/model"
	"errors"
	"strconv"
)

// Хорошим тоном бы объеденить эти методы в 1

func (s *Server) jsonTrans(r *http.Request) error {
	trans := struct {
		Sender int		`json:"sender"`
		Recipient int	`json:"recipient"`
		Amount int		`json:"amount"`
	}{}
	
	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		return err
	}
	
	send := s.cache.GetUser(trans.Sender)
	rec := s.cache.GetUser(trans.Recipient)
	
	if rec == nil {
		return errors.New("Recipient ID: " + strconv.Itoa(trans.Recipient) + " not found")
	} else if send == nil {
		return errors.New("Sender ID: " + strconv.Itoa(trans.Sender) + " not found")
	}
	
	err = send.TransferTo(rec, trans.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) jsonDecode(r *http.Request) (*model.User, error) {
	mod := struct {
		Balance int	`json:"balance"`
		Name string `json:"name"`
	}{}
	
	err := json.NewDecoder(r.Body).Decode(&mod)
	if err != nil {
		return nil, err
	}
	return model.NewUser(mod.Name, mod.Balance), nil
}


func writeError(w http.ResponseWriter, str string, status int) {
	err := struct {
		Err string `json:"error"`
	}{str}
	
	input, _ := json.MarshalIndent(&err, "", "  ")
	http.Error(w, string(input), status)
}


func writeResult(w http.ResponseWriter, user *model.User) {
	request := struct{
		ID int	`json:"id"`
		Name string	`json:"name"`
		Balance int `json:"balance"`
	}{
		ID: user.GetID(),
		Name: user.GetName(),
		Balance: user.GetBalance(),
	}
	
	input, _ := json.MarshalIndent(&request, "", "  ")
	_, err := w.Write(input)
	if err !=nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}
}


// Log and Debug
func (s *Server) debug(w http.ResponseWriter, r *http.Request) {
	writeDebug(w, "Helloworld")
}

func writeDebug(w http.ResponseWriter, str string) {
	res := struct {
		Msg string `json:"request"`
	}{str}
	
	input, _ := json.MarshalIndent(&res, "", "  ")
	_ , err := w.Write(input)
	if err != nil {
		log.Println(err)
	}
}

