package artifactory

type Repository struct {
    Key string `json:"key"`
    Type string `json:"type"`
    Url string `json:"url"`
    PackageType string `json:"packageType"`
}

type RepositoryNotUsedSince struct {
    Uri string `json:"uri"`
    DownloadCount int `json:"downloadCount"`
    LastDownloaded string `json:"lastDownloaded"`
}

type Artifact struct {
    Uri string `"json: uri"`
    Repo string `"json: repo"`
    Path string `"json: path"`
    RemoteUrl string `"json: remoteUrl"`
    Created string `"json: created"`
    CreatedBy string `"json: createdBy"`
    LastModified string `"json: lastModified"`
    ModifiedBy string `"json: modifiedBy"`
    LastUpdated string `"json: lastUpdated"`
    Size string `"json: size"`
}
