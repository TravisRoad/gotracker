package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"travisroad/gotracker/auth"
	"travisroad/gotracker/config"
	"travisroad/gotracker/models"
	"travisroad/gotracker/route"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	preTest()
	prepareData()
	m.Run()
}

func preTest() {
	dbFile := "/tmp/sqlmock_db.sqlite"
	exec.Command("rm", "-f", dbFile).Run()

	config.Conf = &config.Config{}
	config.Conf.SetDefaults()

	db, err := models.ConnectSqlite(dbFile)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}
	models.DB = db

}

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
	r := route.RouteInit()

	req, err := http.NewRequest("POST",
		"/api/auth/login",
		bytes.NewReader([]byte(`{"username": "foo", "password": "bar"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal("response body is not a json")
	}
	tokenString, ok := data["token"].(string)
	if !ok {
		t.Fatal("there is no \"token\" field")
	}

	if err := auth.TokenStringValid(tokenString); err != nil {
		t.Fatal(err)
	}
}

func TestLoginFail(t *testing.T) {
	r := route.RouteInit()

	req, err := http.NewRequest("POST",
		"/api/auth/login",
		bytes.NewReader([]byte(`{"username": "foofoo", "password": "barbar"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal("response body is not a json")
	}
	t.Log(data)
	errorMsg, ok := data["error"].(string)
	if !ok {
		t.Fatal("there is no \"srrorMsg\" field")
	}

	assert.Equal(t, errorMsg, "record not found")

}

func TestRegister(t *testing.T) {
	r := route.RouteInit()

	req, err := http.NewRequest("POST",
		"/api/auth/register",
		bytes.NewReader([]byte(`{"username": "foofoo", "password": "barbar"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var count int64
	if err := models.DB.Model(&models.User{}).Where("username = ?", "foofoo").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count == 0 {
		t.Fatal("there is no user")
	}

	var data map[string]interface{}

	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal("response body is not a json")
	}

	message, ok := data["message"].(string)
	if !ok {
		t.Fatal("there is no \"message\" field")
	}
	assert.Equal(t, message, "success")
}

func TestRegisterUsernameTaken(t *testing.T) {
	r := route.RouteInit()

	req, err := http.NewRequest("POST",
		"/api/auth/register",
		bytes.NewReader([]byte(`{"username": "foo", "password": "bar"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var count int64
	if err := models.DB.Model(&models.User{}).Where("username = ?", "foo").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatal("there are duplicated users")
	}

	var data map[string]interface{}

	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal("response body is not a json")
	}

	error, ok := data["error"].(string)
	if !ok {
		t.Fatal("there is no \"error\" field")
	}
	assert.Equal(t, error, "username is already taken")
}

func TestRegisterUsernameNotValid(t *testing.T) {
	r := route.RouteInit()

	req, err := http.NewRequest("POST",
		"/api/auth/register",
		bytes.NewReader([]byte(`{"username": "12", "password": "bar"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var count int64
	if err := models.DB.Model(&models.User{}).Where("username = ?", "foo").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatal("there are duplicated users")
	}

	var data map[string]interface{}

	if err := json.NewDecoder(rr.Body).Decode(&data); err != nil {
		t.Fatal("response body is not a json")
	}

	error, ok := data["error"].(string)
	if !ok {
		t.Fatal("there is no \"error\" field")
	}
	assert.Equal(t, error, "username is invalid")
}
