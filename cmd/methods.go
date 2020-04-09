package main


func (r *HealthResponse) New() *HealthResponse {
	if r == nil {
		return new(HealthResponse).Init()
	}
	return r
}


func (r *HealthResponse) Init() *HealthResponse {
	r.Status = true

	return r
}
