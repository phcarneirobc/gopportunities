package handler

import "fmt"

// CreateOpening

func errParamIsRequired(name,typ string) error{
	return fmt.Errorf("param: %s(type %s) is required", name, typ)
}

type CreateOpeningRequest struct{
	Name string `json:"name"`
	CPF string `json:"cpf"`
	Location string `json:"location"`
	Priority *bool `json:"priority"`
	Number string `json:"Number"`
	//Salary int64 `json:"salary"`

}

func(r *CreateOpeningRequest) Validate() error{
	if r.CPF == "" && r.Name == "" && r.Location == "" && r.Priority == nil /*&& r.Salary <= 0*/{
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.CPF == ""{
		return errParamIsRequired("cpf","string")
	}
	if r.Name == ""{
		return errParamIsRequired("name","string")
	}
	if r.Location == ""{
		return errParamIsRequired("location","string")
	}
	if r.Number == ""{
		return errParamIsRequired("number","string")
	}
	if r.Priority == nil{
		return errParamIsRequired("priority","bool")
	}
	/*
	if r.Salary <= 0{
		return errParamIsRequired("salary","int64")
	}
	*/
	return nil
}

type UpdateOpeningRequest struct{
	Name string `json:"name"`
	CPF string `json:"cpf"`
	Location string `json:"location"`
	Priority *bool `json:"priority"`
	Number string `json:"Number"`
	//Salary int64 `json:"salary"`
}

func(r*UpdateOpeningRequest) Validate() error{
	// If any is provided, validation is truthy

	if r.CPF != "" || r.Name != "" || r.Location != "" || r.Priority != nil || r.Number != "" /*|| r.Salary >0*/{
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")

}