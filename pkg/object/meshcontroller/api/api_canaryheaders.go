package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/megaease/easegress/pkg/api"
	"github.com/megaease/easegress/pkg/object/meshcontroller/spec"
)

func (a *API) getGlobalCanaryHeaders(meta *partMeta) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := a.service.GetGlobalCanaryHeaders()
		if headers == nil {
			api.HandleAPIError(w, r, http.StatusNotFound,
				fmt.Errorf("global canary not found"))
			return
		}

		buff, err := json.Marshal(headers)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(buff)
		return
	})
}

func (a *API) putGlobalCanaryHeaders(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.HandleAPIError(w, r, http.StatusInternalServerError,
			fmt.Errorf("read body failed: %v", err))
	}

	headers := spec.GlobalCanaryHeaders{}
	err = json.Unmarshal(body, &headers)
	if err != nil {
		api.HandleAPIError(w, r, http.StatusInternalServerError,
			fmt.Errorf("unmarshal %s to pb spec %#v failed: %v", string(body), headers, err))
	}

	a.service.PutGlobalCanaryHeaders(&headers)

	return
}
