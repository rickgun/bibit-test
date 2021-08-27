package repository

import (
	"bibit-test/lib/customerror"
	mocks "bibit-test/mocks/repository"
	"bibit-test/models"
	"bibit-test/testdata"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func TestInsertSearch(t *testing.T) {
	mockListSearch := make([]*models.MoviesLogSearch, 0)
	testdata.GoldenJSONUnmarshal(t, "movieslogsearch", &mockListSearch)
	var mockSearch = mockListSearch[0]

	tests := map[string]struct {
		expectInsert testdata.FuncCaller
		expectError  error
	}{
		"success": {
			expectInsert: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mockSearch},
				Output:   []interface{}{nil},
			},
			expectError: nil,
		},
		"error": {
			expectInsert: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mockSearch},
				Output:   []interface{}{fmt.Errorf("error")},
			},
			expectError: customerror.ErrFailedCommunicateWithRepository,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			repoMock := new(mocks.MoviesLogSearch)

			dbMock, _, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
			}
			defer dbMock.Close()

			db, err := gorm.Open("mysql", dbMock)
			require.NoError(t, err)

			if test.expectInsert.IsCalled {
				repoMock.On("Search", test.expectInsert.Input...).
					Return(test.expectInsert.Output...).
					Once()
			}

			moviesLogSearchUsecase := NewMoviesLogRepository(db)
			err = moviesLogSearchUsecase.InsertSearch(mockSearch)
			if err != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestInsertDetail(t *testing.T) {
	mockListDetail := make([]*models.MoviesLogDetail, 0)
	testdata.GoldenJSONUnmarshal(t, "movieslogdetail", &mockListDetail)
	var mockDetail = mockListDetail[0]

	tests := map[string]struct {
		expectInsert testdata.FuncCaller
		expectError  error
	}{
		"success": {
			expectInsert: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mockDetail},
				Output:   []interface{}{nil},
			},
			expectError: nil,
		},
		"error": {
			expectInsert: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mockDetail},
				Output:   []interface{}{fmt.Errorf("error")},
			},
			expectError: customerror.ErrFailedCommunicateWithRepository,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			repoMock := new(mocks.MoviesLogDetail)

			dbMock, _, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
			}
			defer dbMock.Close()

			db, err := gorm.Open("mysql", dbMock)
			require.NoError(t, err)

			if test.expectInsert.IsCalled {
				repoMock.On("Detail", test.expectInsert.Input...).
					Return(test.expectInsert.Output...).
					Once()
			}

			moviesLogDetailUsecase := NewMoviesLogRepository(db)
			err = moviesLogDetailUsecase.InsertDetail(mockDetail)
			if err != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
