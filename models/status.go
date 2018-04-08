package models

/*StatusCode Type*/
type StatusCode int

/*PandaStatusCode extra status codes for polyPanda*/
type PandaStatusCode int

/*HTTP Return Status Code*/
const (
	St200OK                  StatusCode = 200
	St201Created             StatusCode = 201
	St202Accepted            StatusCode = 202
	St203                    StatusCode = 203
	St204NoContent           StatusCode = 204
	St300MultipleChoices     StatusCode = 300
	St304NotModified         StatusCode = 304
	St400BadRequest          StatusCode = 400
	St401UnAuthorized        StatusCode = 401
	St404NotFound            StatusCode = 404
	St405MethodNotAllowed    StatusCode = 405
	St408RequestTimeout      StatusCode = 408
	St409Conflict            StatusCode = 409
	St500InternalServerError StatusCode = 500
	St503ServiceUnavailable  StatusCode = 503
)
const ()
