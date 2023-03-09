package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/arkhipovkm/go.neose-mini.firmata-client/neose_mini"
)

//go:embed build/*
var content embed.FS

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
func LedOn(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.LedOn()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}

func LedOff(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.LedOff()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}

func LcsOn(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.LCShutterOn()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}

func LcsOff(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.LCShutterOff()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}


func FanOn(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.FanOn()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}

func FanOff(neose *neose_mini.NeoseMini) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w = setHeaders(w)
		err := neose.FanOff()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("NOT OK"))
			return
		}
		w.Write([]byte("OK"))
	}
}

func FileServerWithDefaultFile(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/build" + r.URL.Path
		log.Println(r.URL.Path)
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {
	var err error

	neose := neose_mini.NewNeoseMini()
	err = neose.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer neose.Disconnect()

	http.HandleFunc("/led/on", LedOn((neose)))
	http.HandleFunc("/led/off", LedOff((neose)))
	http.HandleFunc("/lcs/on", LcsOn((neose)))
	http.HandleFunc("/lcs/off", LcsOff((neose)))
	http.HandleFunc("/fan/on", FanOn((neose)))
	http.HandleFunc("/fan/off", FanOff((neose)))

	iface := "localhost:4567"
	
	staticIface := "localhost:3001"
	log.Println("Dynamic server is listening on", iface)
	go http.ListenAndServe("localhost:4567", nil)

	log.Printf("Static server is listening on http://%s. Please click on the link to open the UI!", staticIface)
	http.ListenAndServe(staticIface, FileServerWithDefaultFile(http.FS(content)))
	
}