package usercontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	routerFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/router"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth/dto"
	"github.com/stretchr/testify/assert"
)

var url string

func TestMain(m *testing.M) {
	url = startTestServer()

	code := m.Run()

	os.Exit(code)
}

func makeUrl(uri string) string {
	return url + uri
}

func TestCreateUser(t *testing.T) {
	userDto := makeNewUserDto()
	assert := assert.New(t)

	t.Run("Should successfully create a new user", func(t *testing.T) {
		response := makeRequest[models.User](http.MethodPost, makeUrl("/user"), makeNewUserDto(), nil)
		user := response.Content

		addressDto := userDto.Address
		address := user.AddressList[0]

		assert.Equal(200, response.HttpStatus)

		assert.Equal(int64(1), user.Id)
		assert.Equal(userDto.Name, user.Name)
		assert.Equal(userDto.Age, user.Age)
		assert.Equal(userDto.Email, user.Email)

		assert.Equal(int64(1), address.Id)
		assert.Equal(addressDto.Alias, address.AddressAlias)
		assert.Equal(addressDto.ZipCode, address.ZipCode)
		assert.Equal(addressDto.StreetName, address.StreetName)
		assert.Equal(addressDto.Number, address.Number)
		assert.Equal(addressDto.State, address.State)
		assert.Equal(addressDto.Country, address.Country)
	})

	t.Run("Should fail to create new user due to already in use email", func(t *testing.T) {
		response := makeRequest[string](http.MethodPost, makeUrl("/user"), makeNewUserDto(), nil)

		assert.Equal(409, response.HttpStatus)
	})

}

func TestAuthentication(t *testing.T) {
	assert := assert.New(t)
	user := makeNewUserDto()

	t.Run("Should successfully authenticate and get JWT token", func(t *testing.T) {
		response := makeRequest[string](http.MethodPost, makeUrl("/auth"), dto.AuthDto{
			Email:    user.Email,
			Password: user.Password,
		}, nil)

		assert.Equal(200, response.HttpStatus)
		assert.NotEmpty(response.Content)
	})

	t.Run("Should fail to create new user due to already in use email", func(t *testing.T) {
		response := makeRequest[string](http.MethodPost, makeUrl("/auth"), dto.AuthDto{
			Email:    user.Email,
			Password: "wrong password",
		}, nil)

		assert.Equal(401, response.HttpStatus)
		assert.Empty(response.Content)
	})

}

func TestGetUserById(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should successfully fetch user by id", func(t *testing.T) {
		jwt := getAuthToken()

		response := makeRequest[any](http.MethodGet, makeUrl("/user/1"), nil, map[string]string{
			"Authorization": jwt,
		})

		assert.Equal(200, response.HttpStatus)
	})
}

func getAuthToken() string {
	user := makeNewUserDto()

	response := makeRequest[string](http.MethodPost, makeUrl("/auth"), dto.AuthDto{
		Email:    user.Email,
		Password: user.Password,
	}, nil)

	return response.Content
}

func startTestServer() string {
	routerFactory.BootControllers()
	router := routerFactory.GetRouter()

	srv := httptest.NewServer(router)
	return srv.URL
}

func makeNewUserDto() dto.CreateUserDto {
	return dto.CreateUserDto{
		Name:     "John Doe",
		Age:      36,
		Email:    "johndoe@genericmail.com",
		Password: "password+123",
		Address: dto.CreateAddressDto{
			Alias:      "Home",
			ZipCode:    "00000",
			StreetName: "Generic Street",
			Number:     "007",
			State:      "GS",
			Country:    "GC",
		},
	}
}

func makeRequest[T any](method, url string, rawBody any, h map[string]string) httputils.HttpResponse[T] {
	body, err := json.Marshal(rawBody)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	reader := bytes.NewReader(body)
	req, _ := http.NewRequest(method, url, reader)

	for k, v := range h {
		req.Header.Set(k, v)
	}

	res, _ := client.Do(req)

	content := new(T)
	response := httputils.HttpResponse[T]{
		Content: *content,
	}

	json.NewDecoder(res.Body).Decode(&response)

	return response
}

func insertUserInDb[T any](url string, createUserDto dto.CreateUserDto) httputils.HttpResponse[T] {
	body, err := json.Marshal(createUserDto)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/user", url), "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	bytes, _ := io.ReadAll(resp.Body)
	log.Println(string(bytes))

	content := new(T)
	response := httputils.HttpResponse[T]{
		Content: *content,
	}

	if err := json.Unmarshal(bytes, &response); err != nil {
		panic(err.Error())
	}

	return response
}
