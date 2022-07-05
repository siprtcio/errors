package errors

type ApiError struct {
	Resource  string `json:"-"`
	ErrorCode B2Code `json:"-"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	MoreInfo  string `json:"more_info"`
	Status    int    `json:"status"`
}

func (ae *ApiError) Generate() error {
	ae.setDefaults()

	if err := ae.SetMessage(); err != nil {
		return err
	}

	if err := ae.SetMoreInfo(); err != nil {
		return err
	}

	if err := ae.SetStatus(); err != nil {
		return err
	}

	if err := ae.SetCode(); err != nil {
		return err
	}

	return nil
}


func (ae *ApiError) SetMessage() error {
	var err error
	ae.Message, err = parseErrorField(ErrorTemplates[ae.ErrorCode].Message, ae)
	return err
}

func (ae *ApiError) SetMoreInfo() error {
	var err error
	ae.Message, err = parseErrorField(ErrorTemplates[ae.ErrorCode].MoreInfo, ae)
	return err
}

func (ae *ApiError) SetStatus() error {
	var err error
	ae.Status = ErrorTemplates[ae.ErrorCode].Status
	return err
}

func (ae *ApiError) SetCode() error {
	var err error
	ae.Code = ErrorTemplates[ae.ErrorCode].Code
	return err
}

func (ae *ApiError) setDefaults() {
	if ae.ErrorCode == 0 {
		ae.ErrorCode = ResourceNotFound
	}
}
