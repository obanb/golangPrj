package activeEndpoint

import (
	"awesomeProject/app/mock/metering"
	"awesomeProject/app/mock/reflection"
	"awesomeProject/app/mock/socket"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)



/*
ActiveEndpointService
*/

var activeEndpointHandler = ActiveEndpointHandler{
	ActiveEndpoints: make(map[string]ActiveEndpoint),
}

func Register(r *RegisterActiveEndpointRequest) map[string]interface{}{
	rp := reflection.ReflectAndDescribe(r.DataPattern)

	activeEndpoint := r.toActiveEndpoint(rp)

	ae := activeEndpointHandler.setEndpoint(activeEndpoint)

	go AttachCmdListener(r, &activeEndpointHandler)
	go AttachBroadcastListener(r, &activeEndpointHandler)
	go AttachMeteringListener(r, &activeEndpointHandler)

	return ae
}


func GetRegistered() map[string]ActiveEndpoint {
	return activeEndpointHandler.ActiveEndpoints
}


func AttachCmdListener(r *RegisterActiveEndpointRequest, bc *ActiveEndpointHandler) {

	currentEndpoint := bc.ActiveEndpoints[r.EndpointName]
	client := http.Client{}


	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	defer currentEndpoint.WaitGroup.Done()

	currentEndpoint.WaitGroup.Add(1)

	survey := &metering.ActiveEndpointSurvey{}

	var activeCommand = "Pause"
	for {
		select {
		case cmd := <- currentEndpoint.OperationChannel:
			fmt.Println(cmd)
			switch cmd {
			case "Kill":
				return
			case "Pause":
				fmt.Println("PAUZA")
				activeCommand = "Pause"
			default:
				activeCommand = "Run"
			}
		case <-time.After(1 * time.Second):
		if activeCommand == "Run" {
			fmt.Println("will be run")

			start := time.Now()

			res, _ := sendHttpRequest(r,client, true, currentEndpoint.ReflectionPattern)

			end := time.Since(start)

			metering := metering.ActiveEndpointMetering{
				Status: res.Status,
				Duration: int64(end / time.Millisecond),
			}
			fmt.Println("DURATION")
			fmt.Println(end)

			fmt.Println("send")

			currentEndpoint.MeteringChannel <- survey.MergeNext(metering)
			currentEndpoint.BroadcastChannel <- res.Status
			}
		}
	}

	currentEndpoint.WaitGroup.Wait()
}


func AttachBroadcastListener(r *RegisterActiveEndpointRequest, bc *ActiveEndpointHandler) {
	currentEndpoint := bc.ActiveEndpoints[r.EndpointName]

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	for {
		e := <- currentEndpoint.BroadcastChannel
		room := socket.WsConnections.Rooms[currentEndpoint.Name]
		fmt.Println("NAME")
		fmt.Println(room)

		room.Broadcast(e)
	}
}

func AttachMeteringListener(r *RegisterActiveEndpointRequest, bc *ActiveEndpointHandler) {
	currentEndpoint := bc.ActiveEndpoints[r.EndpointName]

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	for {
		e := <- currentEndpoint.MeteringChannel
		room := socket.WsConnections.Rooms[currentEndpoint.Name]
		fmt.Println(e)
		fmt.Println(room)

		room.Broadcast(e)
	}
}


func RunCmd(r *RunActiveEndpointCommandRequest){
	activeEndpoint := activeEndpointHandler.getEndpoint(r.EndpointName)
	activeEndpoint.OperationChannel <- r.Command
}



/*
ActiveEndpointService - helpers
*/


func sendHttpRequest(r *RegisterActiveEndpointRequest, c http.Client, prx bool, rp map[string]interface{}) (*http.Response, error){
	fmt.Println("sendHttpRequest")
	payloadBuf := new(bytes.Buffer)

	reflectedData := reflection.GenerateReflectionCopy(rp)
	json.NewEncoder(payloadBuf).Encode(reflectedData)

	var url string
	var suffix string

	if prx {
		url = "http://localhost:8082/"
		suffix = "eventLog"
	}else {
		url = r.EndpointName
		suffix = "efefe"
	}
	fmt.Println("sendHttpRequest")

	fmt.Println(url + suffix)

	req, err := http.NewRequest(http.MethodPost, url + suffix, payloadBuf)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(req.URL)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
