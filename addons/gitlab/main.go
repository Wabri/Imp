package gitlab

import (
	"encoding/json"
	"strconv"

	"imp/utils/http"
)

var requestHandler http.RequestHandler = http.RequestHandler{ 
    Url: base_url + api_prefix,
    TokenHeader: "PRIVATE-TOKEN",
    Token: token,
}

func GetProjectById(id int) Project {
    var project Project
    requestHandler.Url = requestHandler.Url + "/projects/" + strconv.Itoa(id)

    raw := http.GetRequest(requestHandler)
    json.Unmarshal(raw, &project) 

    requestHandler.Url = base_url + api_prefix
    return project
}

func GetProjectHooksById(id int) []Hook {
    var hooks []Hook
    requestHandler.Url = requestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks"

    raw := http.GetRequest(requestHandler)
    json.Unmarshal(raw, &hooks) 
    
    requestHandler.Url = base_url + api_prefix
    return hooks
}

func PutProjectHooksById(id int, hook Hook) bool {
    requestHandler.Url = requestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks/" + strconv.Itoa(hook.Id) 

    payload := make([][2]string, 2)
    payload[0][0] = "url"
    payload[0][1] = hook.Url
    payload[1][0] = "enable_ssl_verification"
    payload[1][1] = strconv.FormatBool(hook.SslEnable)
    valid := http.PutRequest(requestHandler, payload)

    requestHandler.Url = base_url + api_prefix
    return valid
}
