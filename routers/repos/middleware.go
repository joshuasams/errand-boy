package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NormalizePRPayload turns a bitbucket-specific PR payload into a general one.
func NormalizePRPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(r.URL.Path)
		// NOTICE: url.Host == "" when on localhost.
		switch url.Host {
		case "github.com":
			prPayload := &gitHubPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGenericPR()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
		case "bitbucket.com":
			prPayload := &bitBucketPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGenericPR()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
		}
		next.ServeHTTP(w, r)
	})
}