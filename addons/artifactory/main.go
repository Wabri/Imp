package artifactory

import (
	"encoding/json"
	"imp/utils/http"
	"strconv"
	"time"
)

var RequestHandler http.RequestHandler = http.RequestHandler{ 
    Url: base_url + api_prefix,
    TokenHeader: "X-JFrog-Art-Api",
    Token: token,
}

func GetRepositories() []Item {
    var repositories []Item
    RequestHandler.Url = RequestHandler.Url + "/repositories"

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &repositories) 

    RequestHandler.Url = base_url + api_prefix
    return repositories
}

func GetArtifacts(path string) []Item{
    RequestHandler.Url = RequestHandler.Url + "/storage/" + path

    type Response struct {
        Item []Item `json:"children"`
    }
    var response Response

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &response) 

    RequestHandler.Url = base_url + api_prefix
    return response.Item
}

func GetItemInfos(repository string, path string) []Item {
    RequestHandler.Url = RequestHandler.Url + "/search/aql"
    if len(path) == 0 {
	RequestHandler.Data = "items.find({\"repo\":\"" + repository + "\"})"
    } else {
        RequestHandler.Data = "items.find({\"repo\":\"" + repository + "\", \"path\":\"" + path + "\"})"
    }

    type Response struct {
        Item []Item `json:"results"`
    }
    var response Response

    raw := http.PostRequestTextPlain(RequestHandler)
    json.Unmarshal(raw, &response) 

    RequestHandler.Url = base_url + api_prefix
    return response.Item
}

func GetArtifactNotUsedSinceForRepository(repository string, date string) []Item {
    since, err := time.Parse("2006-01-02", date)
    if err != nil {
        panic(err)
    }

    type Response struct {
        Repositories []Item `json:"results"`
    }
    var response Response

    RequestHandler.Url = RequestHandler.Url + "/search/usage?notUsedSince=" + strconv.FormatInt(since.UnixMilli(), 10) + "&repos=" + repository

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &response) 

    RequestHandler.Url = base_url + api_prefix
    return response.Repositories
}

func DeleteArtifact(repository string, artifact string) bool {
    RequestHandler.Url = base_url + "/" + repository + "/" + artifact

    result := http.DeleteRequest(RequestHandler)

    RequestHandler.Url = base_url + api_prefix
    return result
}
