package repository

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testPostgres *gorm.DB
var err error

func Test_GetAccountById(t *testing.T) {

	urlStr := os.Getenv("TEST_DB_URL")
	if urlStr == "" {
		urlStr = "postgres://testuser:testpass@localhost:5432/testdb"
		const format = "env TEST_DB_URL is empty, used default value: %s"
		t.Logf(format, urlStr)
	}

	testPostgres, err = InitDB(urlStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Successfully initialized")

	var dest interface{}
	res := testPostgres.Model(&models.Account{}).First(&dest)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("Printing dest:\n %s", dest)

	type args struct {
		Text   string
		Target string
		Source string
	}

	tests := []struct {
		name                 string
		args                 args
		expectedStatusCode   int
		expectedErrorMessage string
		want                 string
		wantErr              bool
	}{
		{
			name: "Ok",
			args: args{
				"car",
				"en",
				"ru",
			},
			expectedStatusCode:   http.StatusOK,
			expectedErrorMessage: "",
			want:                 "машина",
			wantErr:              false,
		},
		{
			name: "Error",
			args: args{
				"car",
				"en",
				"ru",
			},
			expectedStatusCode: http.StatusBadGateway,
			want:               "",
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// func call here

			var err error
			var got string
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}

}
