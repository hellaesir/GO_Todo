package entities

type Task struct{
	Id string `json:"id,omitempty"`
    Description string `json:"description,omitempty"`
	SupposedStartDate string `json:"supposedStartDate,omitempty"`
	SupposedEndDate string `json:"supposedEndDate,omitempty"`
	RealStartDate string `json:"realStartDate,omitempty"`
	RealEndDate string `json:"realEndDate,omitempty"`
	Checked bool
}