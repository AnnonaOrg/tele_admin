package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/internal/response"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/apis/v1|/login|/admin/api)")
		if !reg.MatchString(path) {
			return
		}

		if path == "/user/login" || path == "/admin/index" || path == "/admin/logout" || path == "/admin/404" || path == "/admin/res" {
			return
		}

		//Read the Body context
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		//Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		//The basic informations
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("new request come in, path:%s,Method:%s,body `%s`",path,method,string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		//continue
		c.Next()

		//calculates the latency
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		//get code and message
		var resp response.Response
		if err := json.Unmarshal(blw.body.Bytes(), &resp); err != nil {
			log.Errorf("%v response body can not unmarshal to model.Response struct,body:`%s`",
				err, blw.body.Bytes(),
			)
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = resp.Code
			message = resp.Message
		}

		log.Infof("%-13s|%-12s|%s %s|{code:%d,message:%s}", latency, ip, pad.Right(method, 5, ""), path, code, message)

	}
}
