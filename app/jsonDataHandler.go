package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go/types"
	"log"
	"net/http"
	"reflect"
	"sync"
	"time"
)

type JsonDataHandlers struct {
}

type RegisterEndpointRequest struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
}

type RegisterActiveEndpointRequest struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
	Options      ActiveEndpointOptions  `json:"activeEndpointOptions"`
	TargetURI    string                 `json:"targetURI"`
	UseProxy     bool                   `json:"useProxy"`
}

type RunActiveEndpointCommandRequest struct {
	EndpointName string                 `json:"endpointName"`
	Command string `json:"command"`
}

type ActiveEndpointActionRequest struct {
	EndpointName string `json:"endpointName"`
	Action       string `json:"action"`
}

type ActiveEndpointActionResponse struct {
	EndpointName string `json:"endpointName"`
	Action       string `json:"action"`
	Status string `json:"status"`
}

type RegisterEndpointResponse struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
}

type RegisterEndpointOptions struct {
}

type ActiveEndpointHandler struct {
	ActiveEndpoints map[string]ActiveEndpoint
}

type ActiveEndpointOptions struct {
	Burst     bool `json:"burst"`
	BurstSize *int `json:"burstSize"`
}

type ActiveEndpoint struct {
	TargetURI         string                 `json:"targetURI"`
	Name              string                 `json:"name"`
	ReflectionPattern map[string]interface{} `json:"reflectionPattern"`
	Options           ActiveEndpointOptions  `json:"activeEndpointOptions"`
	BroadcastChannel  chan string
	OperationChannel  chan string
	WaitGroup *sync.WaitGroup
}

type EndpointPersistence struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:"_id,omitempty`
	URI         string                 `json:"URI" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	CreatedAt   string                 `json:"createdAt" bson:"created_at"`
	Active      bool                   `json:"active" bson:"active"`
	Running     bool                   `json:"running" bson:"running"`
	DataPattern map[string]interface{} `json:"dataPattern" bson:"dataPattern"`
	// reflection, mirror
	Kind string `json:"kind" bson:"kind"`
	// active, passive
	Mode        string `json:"kind" bson:"kind"`
	TargetURI   string `json:"targetURI" bson:"targetURI"`
	HeartbeatMS string `json:"heartbeatMS" bson:"heartbeatMS"`
	BatchSize   string `json:"batchSize" bson:"batchSize"`
	SpyHttpCode string `json:"spyHttpCode" bson:"spyHttpCode"`
}

var activeEndpointHandler = ActiveEndpointHandler{
	ActiveEndpoints: make(map[string]ActiveEndpoint),
}

func (aeh ActiveEndpointHandler) setHandler(ae ActiveEndpoint) {
	aeh.ActiveEndpoints[ae.Name] = ae
}

func (aeh ActiveEndpointHandler) getHandler(n string) ActiveEndpoint {
	return aeh.ActiveEndpoints[n]
}

func (raer RegisterActiveEndpointRequest) toActiveEndpoint(rp map[string]interface{}) ActiveEndpoint {
	return ActiveEndpoint{
		TargetURI:         raer.TargetURI,
		Name:              raer.EndpointName,
		ReflectionPattern: rp,
		BroadcastChannel:  make(chan string),
		OperationChannel:  make(chan string),
		WaitGroup: new(sync.WaitGroup),
	}
}

func (h JsonDataHandlers) RegisterEndpoint(ginSrv *gin.Engine, uri string) func(*gin.Context) {
	return func(c *gin.Context) {
		var request RegisterEndpointRequest
		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			jsdh := JsonDataHandlers{}

			ginSrv.GET(uri, jsdh.ServeFromReflection)
			c.JSON(http.StatusCreated, "true")
		}
	}
}

func (h JsonDataHandlers) RegisterActiveEndpoint() func(*gin.Context) {
	return func(c *gin.Context) {
		var request RegisterActiveEndpointRequest
		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			rp := reflectAndDescribe(request.DataPattern)

			activeEndpoint := request.toActiveEndpoint(rp)

			activeEndpointHandler.setHandler(activeEndpoint)

			fmt.Print(activeEndpoint)

			c.JSON(http.StatusCreated, "true")
		}
	}
}

func (h JsonDataHandlers) HandleActiveEndpointAction(c *gin.Context) {
		var request ActiveEndpointActionRequest
		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			UseActiveEndpointAction(request)
			c.JSON(http.StatusCreated, "true")
		}
}

func UseActiveEndpointAction(r ActiveEndpointActionRequest) {
	switch r.Action {
	case "run":
		fmt.Println("run")
		go routineFn(&r, &activeEndpointHandler)
	case "pause":
		fmt.Println("pause")
	case "kill":
		fmt.Println("kill")
	}
}

func (h JsonDataHandlers) ServeFromReflection(c *gin.Context) {
	var request RegisterEndpointRequest
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		reflection := reflectAndDescribe(request.DataPattern)

		data := generateReflectionData(reflection, 5)

		fmt.Println("NEW")
		fmt.Println(data)
		fmt.Println("NEW")

		res := RegisterEndpointResponse{
			EndpointName: request.EndpointName,
			Description:  request.Description,
			DataPattern:  reflection,
		}

		//routineFn :=

		c.JSON(http.StatusCreated, res)
	}
}


func routineFn(r *ActiveEndpointActionRequest, bc *ActiveEndpointHandler) {

	//client := &http.Client{}
	payloadBuf := new(bytes.Buffer)

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			bc.ActiveEndpoints[r.EndpointName].BroadcastChannel <- "every 2 seconds"
			fmt.Println("EVERY 2")
		}
	}()

	select {
	case proc := <-bc.ActiveEndpoints[r.EndpointName].BroadcastChannel:
		fmt.Println("obdrzeno")
		fmt.Println(proc)

		json.NewEncoder(payloadBuf).Encode(proc)
		req, err := http.NewRequest("POST", ProxyURL + "/eventLog", payloadBuf)

		fmt.Println(ProxyURL + "/eventLog")
		if err != nil {
			fmt.Println(err)
			fmt.Println("chyba")
		}

		client := &http.Client{}
		resp, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(resp.Status)

		fmt.Println(proc)
	case <-time.After(10 * time.Second):
		// somehow send tooLate <- true
		//so that we can stop the go routine running
		fmt.Println("too late")
	}
}

func Register(r *ActiveEndpointActionRequest, bc *ActiveEndpointHandler) {

	currentEndpoint := bc.ActiveEndpoints[r.EndpointName]
	client := &http.Client{}


	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	defer currentEndpoint.WaitGroup.Done()


	var activeCommand = "Pause"
	for {
		select {
		case cmd := <- currentEndpoint.OperationChannel:
			fmt.Println(cmd)
			switch cmd {
			case "Kill":
				return
			case "Pause":
				activeCommand = "Pause"
			default:
				activeCommand = "Run"
			}
		default:
			if activeCommand == "Run" {
				for {
					time.Sleep(time.Second * 2)
					SendHttpRequest(r,client, true)
				}
			}
		}
	}

	currentEndpoint.WaitGroup.Wait()
}

func RunCmd(r *RunActiveEndpointCommandRequest, bc *ActiveEndpointHandler){
	activeEndpoint := bc.getHandler(r.EndpointName)
	activeEndpoint.OperationChannel <- r.Command
}

func SendHttpRequest(r *ActiveEndpointActionRequest, c *http.Client, prx bool) (*http.Response, *error){
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("pes")

	var url string
	var suffix string

	if prx {
		url = ProxyURL
		suffix = "eventLog"
	}else {
		url = r.EndpointName
		suffix = "efefe"
	}
	req, _ := http.NewRequest("POST", url + suffix, payloadBuf)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, &err
	}
	return resp, nil
}


type JsonNumberRepresentative = float64
type JsonStringRepresentative = string
type JsonBoolRepresentative = bool
type JsonNullRepresentative = types.Nil
type JsonArrayRepresentative = []string
type JsonObjectRepresentativeS = types.Struct
type JsonObjectRepresentativeM = types.Map

type PropertyKind int64

const (
	single PropertyKind = iota
	slice
)

type PropertyPattern struct {
	Reflect         reflect.Kind `json:"reflect"`
	ReflectionLabel string       `json:"reflectionLabel"`
	Kind            PropertyKind `json:"-"`
	Len             int          `json:"len"`
}

func generateReflectionData(t map[string]interface{}, l int) []map[string]interface{} {
	data := make([]map[string]interface{}, l)
	for i := 0; i < l; i++ {
		copy := generateReflectionCopy(t)
		data = append(data, copy)
	}
	return data
}

func generateReflectionCopy(t map[string]interface{}) map[string]interface{} {
	for key, v := range t {
		// switch by reflected kind
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			// we need cast v to map[string]interface{} for correct recursion call
			if mv, ok := v.(map[string]interface{}); ok {
				generateReflectionCopy(mv)
			} else {
				fmt.Println("unknown property from reflect.Map")
			}

		default:
			if mv, ok := v.(PropertyPattern); ok {
				fmt.Println(mv)
				if mv.Kind == single {
					generateValueFromReflection(mv.Reflect, t, key)
				} else {
					//TODO array gen
				}
			} else {
				fmt.Println("unknown property from reflect.Map")
			}
		}
	}
	return t
}

func generateValueFromReflection(k reflect.Kind, t map[string]interface{}, key string) {
	switch k {
	case reflect.String:
		t[key] = "pes"
	case reflect.Float64:
		t[key] = 999
	case reflect.Bool:
		t[key] = true
	}
}

func reflectAndDescribe(t map[string]interface{}) map[string]interface{} {
	for key, v := range t {
		rv := reflect.ValueOf(v)

		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			t[key] = PropertyPattern{
				Reflect:         reflect.String,
				ReflectionLabel: reflect.String.String(),
				Kind:            single,
				Len:             rv.Len(),
			}
			fmt.Println(reflect.String)
		case reflect.Bool:
			t[key] = PropertyPattern{
				Reflect:         reflect.Bool,
				ReflectionLabel: reflect.Bool.String(),
				Kind:            single,
			}
		case reflect.Float64:
			t[key] = PropertyPattern{
				Reflect:         reflect.Float64,
				ReflectionLabel: reflect.Float64.String(),
				Kind:            single,
			}
		case reflect.Map:
			if mv, ok := v.(map[string]interface{}); ok {
				reflectAndDescribe(mv)
			} else {
				fmt.Println("unknown property from reflect.Map")
			}

		// reflexe array/slice, vytvorenim ret array z puvodniho interface elementu pro ziskani itemu 0 pro reflexi povahy pole (nelze pouzit pro tuples)
		case reflect.Array, reflect.Slice:
			ret := make([]interface{}, rv.Len())
			for i := range ret {
				// set current value as interface
				ret[i] = rv.Index(i).Interface()
			}
			kind := reflect.TypeOf(ret[0]).Kind()
			t[key] = PropertyPattern{
				Reflect:         kind,
				ReflectionLabel: reflect.Float64.String(),
				Kind:            slice,
				Len:             rv.Len(),
			}
		}
	}

	fmt.Println(t)
	return t
}
