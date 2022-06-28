package errors

type B2Code uint16

// General Codes.
const (
	ResourceNotFound      B2Code = 20404
)

var Templates = map[B2Code]ErrorResponse{
  ResourceNotFound: {
		Code: 20404,
		Message: "The requested resource {{.Resource}} was not found",
		MoreInfo: "https://www.twilio.com/docs/errors/20404",
		Status: 404,
	},
}
