package gitlab

type Project struct {
    Id int `json:"id"`
    Name string `json:"name"`
    PathWithNamespace string `json:"path_with_namespace"`
}

type Hook struct {
    Id int `json:"id"`
    Url string `json:"url"`
    ProjectId int `json:"project_id"`
    SslEnable bool `json:"enable_ssl_verification"`
}

type SearchResult struct {
    Basename string `json:"basename"`
    Data string `json:"data"`
    Path string `json:"path"`
    Filename string `json:"filename"`
    Id string `json:"id"`
    Ref string `json:"ref"`
    Startline string `json:"startline"`
    Project_id string `json:"project_id"`
}

