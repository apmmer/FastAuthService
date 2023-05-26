package schemas

// HealthCheckResponse godoc
// HealthCheckResponse represents the response of a health check request
// @Schema
// @ID HealthCheckResponse
// @Required status
// @Property status string "Status" example("Healthy")
type HealthCheckResponse struct {
	Status string `json:"status"`
}
