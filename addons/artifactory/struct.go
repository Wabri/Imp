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
