package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Sisi struct {
	JenisBangun string `json:"jenis bangun"`
	Panjang int `json:"panjang"`
	Lebar int `json:"lebar"`
}

type Hasil struct {
	JenisBangun string `json:"jenis bangun"`
	Luas int `json:"luas"`
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/api/hitung-luas", Luas)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Luas(w http.ResponseWriter, r *http.Request)  {
	var sisi Sisi

	// proses deteksi method http
	if r.Method != "POST"{
		WarpAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//Proses Read Body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		WarpAPIError(w, r, "cant read Body", http.StatusBadRequest)
		return
	}
	//	Proses unmasrsal body input string
	err = json.Unmarshal(body, &sisi); if err != nil{
		WarpAPIError(w, r, "Error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	WarpAPIData(w, r, Hasil{
		JenisBangun: sisi.JenisBangun,
		Luas	: sisi.RumusLuas(),
	}, "Success", http.StatusOK)
}

func (s *Sisi) RumusLuas() int {
	return s.Panjang * s.Lebar
}

func WarpAPIError(w http.ResponseWriter, r *http.Request, message string, code int)  {
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(code)
	result,err := json.Marshal(map[string]interface{}{
		"code":code,
		"error_type":http.StatusText(code),
		"error_detail":message,
	})
	if err == nil{
		w.Write(result)
	}else {
		log.Println(fmt.Println("error Wrap API : %s", err))
	}
}
func WarpAPIData(w http.ResponseWriter, r *http.Request, data interface{},message string, code int)  {
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(code)
	result,err := json.Marshal(map[string]interface{}{
		"code":code,
		"status":message,
		"data":data,
	})
	if err == nil{
		w.Write(result)
	}else {
		log.Println(fmt.Println("error Wrap API Data : %s", err))
	}
}

