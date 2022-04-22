package metering

type ActiveEndpointSurvey struct {
	Requests      int64 `json:"requests"`
	Avgms         int64 `json:"avgms"`
	Slowest       int64 `json:"slowest"`
	Fastest       int64 `json:"fastest"`
	Trend         int64 `json:"trend"`
	Errors        int64 `json:"errors"`
	Reqsec        int64 `json:"reqsec"`
	TotalDuration int64 `json:"totalDuration"`
}

type ActiveEndpointMetering struct {
	Duration int64  `json:"requests"`
	Status   string `json:"status"`
}

func (aes *ActiveEndpointSurvey) MergeNext(m ActiveEndpointMetering) *ActiveEndpointSurvey {
	aes.Requests++
	aes.TotalDuration += m.Duration
	aes.Avgms = aes.Avgms / aes.Requests

	if aes.Slowest < m.Duration {
		aes.Slowest = m.Duration
	}
	if aes.Fastest > m.Duration {
		aes.Fastest = m.Duration
	}

	if m.Status == "200" {
		aes.Errors++
	}

	aes.Reqsec = aes.TotalDuration / 1000

	// todo
	aes.Trend = 100

	return aes
}
