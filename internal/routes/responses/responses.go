package responses

type VersionResponse struct {
	Current string `json:"current"`
	Error   string `json:"error"`
}

func NewVersionResponse(currentVersion string) *VersionResponse {
	return &VersionResponse{
		Current: currentVersion,
	}
}

func NewErrorResponse(error string) *VersionResponse {
	return &VersionResponse{
		Error: error,
	}
}
