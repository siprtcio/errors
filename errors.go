package errors


type ApiError struct {
  BaseUrl  string `json:"-"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}


// Paginate handles the pagination calculation.
func (ae *ApiError) Error() error {
	ae.setDefaults()

	if err := c.Message(); err != nil {
		return err
	}

	if err := c.doPaginate(); err != nil {
		return err
	}

	if err := c.checkLinks(); err != nil {
		return err
	}

	return nil
}

func (ae *ApiError) Message() error {
  if c.BaseUrl == "" {
		return errors.New("BaseUrl value is missing")
	}
}
// Sets the defaults values for current page
// and limit if none of them were provided.
func (ae *ApiError) setDefaults() {
	if ae.Code == 0 {
		ae.Code = 20404
    ae.Status = 404
	}
}

{
    "code": 20404,
    "message": "The requested resource /2010-04-01/Accounts/ACb83eb454cb3275d7b1200902935fed75/Conferences/1.json was not found",
    "more_info": "https://www.twilio.com/docs/errors/20404",
    "status": 404
}
