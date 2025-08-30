package stat

type GetStatResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}

type GetStatSqlResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
