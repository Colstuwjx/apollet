package apollet

type GetStringRequest struct {
	AppId     string `form:"app_id" json:"app_id"`
	Cluster   string `form:"cluster" json:"cluster"`
	Namespace string `form:"namespace" json:"namespace"`
	Key       string `form:"key" json:"key"`
}

type GetStringResponse struct {
	Data string `json:"data"`
}
