package service

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"

	"github.com/bitrise-io/bitrise-workflow-editor/apiserver/tools"
	"github.com/stretchr/testify/require"
)

func TestGetSpecHandler(t *testing.T) {
	t.Log("empty request body")
	{
		req, err := http.NewRequest("POST", "/api/spec", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(PostSpecHandler)

		handler.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code, rr.Body.String())

		var response PostSpecResponseModel
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &response))
	}

	t.Log("request for default steplib")
	{
		defaultSteplib := "https://github.com/bitrise-io/bitrise-steplib.git"
		require.NoError(t, tools.StepmanSetupLibrary(defaultSteplib))

		body := PostSpecRequestBodyModel{
			Libraries: []string{defaultSteplib},
		}
		bytes, err := json.Marshal(body)
		require.NoError(t, err)

		reader := strings.NewReader(string(bytes))

		req, err := http.NewRequest("POST", "/api/spec", reader)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(PostSpecHandler)

		handler.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code, rr.Body.String())

		var response PostSpecResponseModel
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &response))

		require.Equal(t, 1, len(response.LibraryMap))

		_, found := response.LibraryMap[defaultSteplib]
		require.Equal(t, true, found)
	}
}
