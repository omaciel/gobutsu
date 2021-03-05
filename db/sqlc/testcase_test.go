package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/omaciel/gobutsu/util"
	"github.com/stretchr/testify/require"
)

func createRandomTestCase(t *testing.T) Testcase {

	arg := CreateTestCaseParams{
		Classname:    util.RandomClassName(),
		Filename:     util.RandomFileName(),
		Linenumber:   util.RandomLineNumber(),
		Testcasename: util.RandomTestCaseName(),
		Duration:     util.RandomDuration(),
	}

	tc, err := testQueries.CreateTestCase(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tc)

	require.NotZero(t, tc.Testcaseid)

	return tc
}

func TestCreateTestCase(t *testing.T) {
	createRandomTestCase(t)
}

func TestGetTestCase(t *testing.T) {
	tc1 := createRandomTestCase(t)
	tc2, err := testQueries.GetTestCase(context.Background(), tc1.Testcaseid)

	require.NoError(t, err)
	require.NotEmpty(t, tc2)

	require.Equal(t, tc1.Classname, tc2.Classname)
	require.Equal(t, tc1.Duration, tc2.Duration)
	require.Equal(t, tc1.Filename, tc2.Filename)
	require.Equal(t, tc1.Linenumber, tc2.Linenumber)
	require.Equal(t, tc1.Testcasename, tc2.Testcasename)
}

func TestDeleteTestCase(t *testing.T) {
	tc1 := createRandomTestCase(t)
	err := testQueries.DeleteTestCase(context.Background(), tc1.Testcaseid)
	require.NoError(t, err)

	tc2, err := testQueries.GetTestCase(context.Background(), tc1.Testcaseid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, tc2)
}

func TestListTestCases(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTestCase(t)
	}

	arg := ListTestCasesParams {
		Limit: 5,
		Offset: 5,
	}

	testcases, err := testQueries.ListTestCases(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, testcases, 5)

	for _, tc := range testcases {
		require.NotEmpty(t, tc)
	}
}