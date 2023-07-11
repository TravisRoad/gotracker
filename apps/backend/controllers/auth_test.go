package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"travisroad/gotracker/auth"
	"travisroad/gotracker/controllers"
	"travisroad/gotracker/di"
	"travisroad/gotracker/models"
	"travisroad/gotracker/utils"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func prepareData() {
	// init data
	var users []*models.User
	for i := gofakeit.IntRange(15, 25); i > 0; i-- {
		u := &models.User{
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, gofakeit.IntRange(8, 16)),
		}
		users = append(users, u)
	}
	users = append(users, &models.User{
		Username: "foo",
		Password: "bar",
	}) // for login test
	if err := models.DB.Save(&users).Error; err != nil {
		log.Fatal(err)
	}
	models.DB.Model(&models.User{}).Where("username = ?", "foofoo").Delete(&models.User{})      // login fail case
	models.DB.Model(&models.User{}).Where("username = ?", "fooRegister").Delete(&models.User{}) // register case
}

func TestLogin(t *testing.T) {
	utils.PreTest()
	user := &models.User{
		Username: "testlogin",
		Password: "bar",
	}
	user.Save()

	tests := []struct {
		name string
		body controllers.LoginInput
		code int
	}{
		{
			name: "success",
			body: controllers.LoginInput{
				Username: "testlogin",
				Password: "bar",
			},
			code: http.StatusOK,
		},
		{
			name: "password wrong",
			body: controllers.LoginInput{
				Username: "testlogin",
				Password: "wrongpassword",
			},
			code: http.StatusBadRequest,
		},
		{
			name: "no such username",
			body: controllers.LoginInput{
				Username: "no_such_username",
				Password: "bar",
			},
			code: http.StatusBadRequest,
		},
	}

	di.C.Invoke(func(jh *auth.JWTAuthHelper) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)

				bodyBytes, err := json.Marshal(tt.body)
				if err != nil {
					t.Errorf(err.Error())
				}

				c.Request, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")

				controllers.Login(c)

				assert.Equal(t,
					tt.code,
					w.Code,
					"handler returned wrong status code: got %v want %v", w.Code, tt.code)

				if w.Code != http.StatusOK {
					return
				}

				var data map[string]interface{}
				if err := json.NewDecoder(w.Body).Decode(&data); err != nil {
					t.Fatal("response body is not a json")
				}
				tokenString, ok := data["token"].(string)
				if !ok {
					t.Fatal("there is no \"token\" field")
				}
				if err := jh.TokenStringValid(tokenString); err != nil {
					t.Fatal(err)
				}
			})
		}
	})
	utils.PostTest()
}

func TestRegister(t *testing.T) {
	utils.PreTest()
	user := &models.User{
		Username: "testregister-duplicated",
		Password: "bar",
	}
	user.Save()

	tests := []struct {
		name    string
		body    controllers.RegisterInput
		code    int
		wantErr bool
		errMsg  string
	}{
		{
			name: "success",
			body: controllers.RegisterInput{
				Username: "testregister",
				Password: "bar",
			},
			code:    http.StatusOK,
			wantErr: false,
		},
		{
			name: "duplicated",
			body: controllers.RegisterInput{
				Username: "testregister-duplicated",
				Password: "bar",
			},
			code:    http.StatusBadRequest,
			wantErr: true,
			errMsg:  "username is already taken",
		},
		{
			name: "invalid username",
			body: controllers.RegisterInput{
				Username: "f**k",
				Password: "bar",
			},
			code:    http.StatusBadRequest,
			wantErr: true,
			errMsg:  "username is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			bodyBytes, err := json.Marshal(tt.body)
			if err != nil {
				t.Errorf(err.Error())
			}

			c.Request, _ = http.NewRequest("POST", "/api/register", bytes.NewBuffer(bodyBytes))
			c.Request.Header.Set("Content-Type", "application/json")

			controllers.Register(c)

			if w.Code != tt.code {
				t.Errorf("handler returned wrong status code: got %v want %v", w.Code, tt.code)
			}

			if tt.wantErr {
				var data map[string]interface{}

				if err := json.NewDecoder(w.Body).Decode(&data); err != nil {
					t.Fatal("response body is not a json")
				}
				error, ok := data["error"].(string)
				if !ok {
					t.Fatal("there is no \"error\" field")
				}

				assert.Equal(t, error, tt.errMsg)
			}

		})
	}
	utils.PostTest()
}
