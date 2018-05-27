package customError

import (
	"github.com/labstack/echo"
)


type CustomError int

// Request methods
const (
    NoResource CustomError = iota
    InvalidInput
	TestError
)

func (self CustomError) Error() string {
    switch self {
    case NoResource:
        return "GET"
    case InvalidInput:
        return "POST"
    case TestError:
        return "test error"
    default:
        return "Unknown"
    }
}

func (self CustomError) Render(c echo.Context) string {
    switch self {
    case NoResource:
        return "GET"
    case InvalidInput:
        return "POST"
    case TestError:
        return "POST"
    default:
        return "Unknown"
    }
}
