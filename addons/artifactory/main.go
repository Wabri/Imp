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

func GetRepositories() []Repository {
    var repositories []Repository
    RequestHandler.Url = RequestHandler.Url + "/repositories"

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &repositories) 

    RequestHandler.Url = base_url + api_prefix
    return repositories
}

func GetArtifactNotUsedSinceForRepository(repository string, date string) []RepositoryNotUsedSince {
    since, err := time.Parse("2006-01-02", date)
    if err != nil {
        panic(err)
    }

    type RepositoryResponse struct {
        Repository []RepositoryNotUsedSince `json:"results"`
    }
    var response RepositoryResponse

    RequestHandler.Url = RequestHandler.Url + "/search/usage?notUsedSince=" + strconv.FormatInt(since.UnixMilli(), 10) + "&repos=" + repository

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &response) 

    RequestHandler.Url = base_url + api_prefix
    return response.Repository
}

func DeleteArtifact(repository string, artifact string) bool {
    RequestHandler.Url = base_url + "/" + repository + "/" + artifact

    result := http.DeleteRequest(RequestHandler)

    RequestHandler.Url = base_url + api_prefix
    return result
}
