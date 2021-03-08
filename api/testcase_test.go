package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/omaciel/gobutsu/db/mock"
	db "github.com/omaciel/gobutsu/db/sqlc"
	"github.com/omaciel/gobutsu/util"
	"github.com/stretchr/testify/require"
)

func TestGetTestcaseAPI(t *testing.T) {
	testcase := randomTestCase()

	testCases := []struct {
		name string
		testcaseID int64
		buildStubs func(app *mockdb.MockApp)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			testcaseID: testcase.ID,
			buildStubs: func(app *mockdb.MockApp) {
				app.EXPECT().
	  			GetTestCase(gomock.Any(), gomock.Eq(testcase.ID)).
	  			Times(1).
	  			Return(testcase, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTestcase(t, recorder.Body, testcase)
			},		
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
		
			app := mockdb.NewMockApp(ctrl)
			tc.buildStubs(app)
			
			// start test server and send request
			server := newTestServer(t, app)
			recorder := httptest.NewRecorder()
		
			url := fmt.Sprintf("/testcases/%d", tc.testcaseID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
		
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func randomTestCase() db.Testcase {
	return db.Testcase{
		ID:           int64(util.RandomInt(1000)),
		Classname:    util.RandomClassName(),
		Filename:     util.RandomFileName(),
		Linenumber:   util.RandomLineNumber(),
		Testcasename: util.RandomTestCaseName(),
		Duration:     util.RandomDuration(),
		Username:     util.RandomString(8),
	}
}

func requireBodyMatchTestcase(t *testing.T, body *bytes.Buffer, testcase db.Testcase) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotTestcase db.Testcase
	err = json.Unmarshal(data, &gotTestcase)
	require.NoError(t, err)
	require.Equal(t, testcase, gotTestcase)
}
