package activeEndpoint

import (
	"awesomeProject/app/mock/metering"
	"awesomeProject/app/mock/socket"
	"sync"
)

/*
Requests
*/

type RegisterActiveEndpointRequest struct {
	EndpointName string                 `json:"endpointName"`
	Description  *string                `json:"description,omitempty"`
	DataPattern  map[string]interface{} `json:"dataPattern"`
	Options      ActiveEndpointOptions  `json:"activeEndpointOptions"`
	TargetURI    string                 `json:"targetURI"`
	UseProxy     bool                   `json:"useProxy"`
}

type ActiveEndpointOptions struct {
	Burst     bool `json:"burst"`
	BurstSize *int `json:"burstSize"`
}

func (raer RegisterActiveEndpointRequest) toActiveEndpoint(rp map[string]interface{}) ActiveEndpoint {
	return ActiveEndpoint{
		TargetURI:         raer.TargetURI,
		Name:              raer.EndpointName,
		ReflectionPattern: rp,
		BroadcastChannel:  make(chan string),
		MeteringChannel:  make(chan *metering.ActiveEndpointSurvey),
		OperationChannel:  make(chan string),
		WaitGroup: new(sync.WaitGroup),
	}
}


type RunActiveEndpointCommandRequest struct {
	EndpointName string                 `json:"endpointName"`
	Command string `json:"command"`
}

/*
Impl
*/


type ActiveEndpoint struct {
	TargetURI         string                 `json:"targetURI"`
	Name              string                 `json:"name"`
	ReflectionPattern map[string]interface{} `json:"reflectionPattern"`
	Options           ActiveEndpointOptions  `json:"activeEndpointOptions"`
	BroadcastChannel  chan string `json:"-"`
	OperationChannel  chan string `json:"-"`
	MeteringChannel  chan *metering.ActiveEndpointSurvey `json:"-"`
	WaitGroup *sync.WaitGroup `json:"-"`
	Connections []socket.WebSocketConnection `json:"-"`
}


type ActiveEndpointHandler struct {
	ActiveEndpoints map[string]ActiveEndpoint
}

func (aeh ActiveEndpointHandler) setEndpoint(ae ActiveEndpoint) map[string]interface{} {
	aeh.ActiveEndpoints[ae.Name] = ae
	return aeh.ActiveEndpoints[ae.Name].ReflectionPattern
}

func (aeh ActiveEndpointHandler) getEndpoint(n string) ActiveEndpoint {
	return aeh.ActiveEndpoints[n]
}