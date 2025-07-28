package handler

import "fmt"

// Create Opening

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s: (%s string) is required", name, typ)
}

type CreateOpeningRequest struct {
	Company  string `json:"company"`
	Role     string `json:"role"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Company == "" && r.Role == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Company  string `json:"company"`
	Role     string `json:"role"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Company != "" || r.Role != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
