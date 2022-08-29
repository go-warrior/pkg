/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/20 17:59
 */
package response

import (
	"fmt"
	"net/http"

	"github.com/go-warrior/pkg/v2/encoding"
	"github.com/go-warrior/pkg/v2/errors"
	"github.com/go-warrior/pkg/v2/internal/httputil"
)

type response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DefaultResponseEncoder encodes the object to the HTTP response.
func Encoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	codec, _ := CodecForRequest(r, "Accept")

	data, err := codec.Marshal(v)

	if err != nil {
		return err
	}

	body := fmt.Sprintf("{\"success\":true,\"code\":\"0\",\"message\":\"\",\"data\":%s}", string(data))

	w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	w.Write([]byte(body))
	return nil
}

// CodecForRequest get encoding.Codec via http.Request
func CodecForRequest(r *http.Request, name string) (encoding.Codec, bool) {
	for _, accept := range r.Header[name] {
		codec := encoding.GetCodec(httputil.ContentSubtype(accept))
		if codec != nil {
			return codec, true
		}
	}
	return encoding.GetCodec("json"), false
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := CodecForRequest(r, "Accept")

	resp := response{
		Success: false,
		Code:    se.Reason,
		Message: se.Message,
		Data:    nil,
	}

	body, err := codec.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	w.Write(body)
}
