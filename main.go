package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/zhaizhonghao/explorerTool/services/connection"
)

type Success struct {
	Payload string `json:"Payload"`
	Message string `json:"Message"`
}

var tpl *template.Template

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/display/explorer", setupExplorer).Methods("POST", http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))

	fmt.Println("Server is listenning on localhost:8282")

	http.ListenAndServe(":8282", router)
}

func setupExplorer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start setuping the explorer")
	setHeader(w)
	if (*r).Method == "OPTIONS" {
		fmt.Println("Options request discard!")
		return
	}
	var channel = connection.Channel{}
	err := json.NewDecoder(r.Body).Decode(&channel)
	if err != nil {
		fmt.Println("parse channel error", err)
		return
	}
	if channel.ChannelName == "" {
		return
	}
	fmt.Println("Get channel", channel)
	//Generate connection file
	tpl = template.Must(template.ParseGlob("templates/explorer/*.json"))
	file, err := os.Create("connection-profile/first-network_2.2.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Fail to create file!")
	}
	err = connection.GenerateConnectionTemplate(channel, tpl, file)
	if err != nil {
		fmt.Println("Fail to generate firt-network_2.2.json", err)
	}
	//copy crypto-config folder into the explorer folder
	fmt.Println("Copying the crypto-config folder into the the explorer folder")
	out, err1 := exec.Command("cp", "-r", "../BasicNetwork-2.0/artifacts/channel/crypto-config/", "crypto-config").Output()
	if err1 != nil {
		fmt.Println("Copy the crypto-config failed:" + err1.Error())
		fmt.Println(string(out))
		return
	}
	//up the explorer docker
	fmt.Println("uping the explorer docker")
	out, err1 = exec.Command("docker-compose", "up", "-d").Output()
	if err1 != nil {
		fmt.Println("Copy the crypto-config failed:" + err1.Error())
		fmt.Println(string(out))
		return
	}
	success := Success{
		Payload: "setup the explorer successfully!",
		Message: "200 OK",
	}
	json.NewEncoder(w).Encode(success)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("X-Powered-By", "3.2.1")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
}
