package ecpayEcho

import (
	"bytes"
	"github.com/Laysi/go-ecpay-sdk"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func ECPayCheckMacValueHandler(client *ecpay.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Request
			body := []byte{}
			if c.Request().Body != nil { // Read
				body, _ = ioutil.ReadAll(c.Request().Body)
			}

			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
			err := c.Request().ParseForm()
			if err != nil {
				return c.NoContent(http.StatusBadRequest)

			}

			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

			params := ecpay.ECPayValues{c.Request().PostForm}.ToMap()
			c.Request().Form = nil
			c.Request().PostForm = nil

			senderMac := params["CheckMacValue"]
			delete(params, "CheckMacValue")
			mac := client.GenerateCheckMacValue(params)
			if mac != senderMac {
				return c.String(http.StatusBadRequest, "0|Error")
			}
			return next(c)
		}
	}

}
