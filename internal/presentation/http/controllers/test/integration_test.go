package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	routerFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/router"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/dto"
	"github.com/stretchr/testify/assert"
)

const (
	CREATE = iota + 1
	READ
	UPDATE
	DELETE
)

const (
	EMAIL_1 = "johndoe@genericmail.com"
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

		assert.Equal(http.StatusOK, response.HttpStatus)

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

		assert.Equal(http.StatusConflict, response.HttpStatus)
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

		assert.Equal(http.StatusOK, response.HttpStatus)
		assert.NotEmpty(response.Content)
	})

	t.Run("Should fail wrong credentials", func(t *testing.T) {
		response := makeRequest[string](http.MethodPost, makeUrl("/auth"), dto.AuthDto{
			Email:    user.Email,
			Password: "wrong password",
		}, nil)

		log.Println(response.Message)
		assert.Equal(http.StatusUnauthorized, response.HttpStatus)
		assert.Empty(response.Content)
	})

}

func TestGetUserById(t *testing.T) {
	assert := assert.New(t)
	sampleUser := makeNewUserDto()
	sampleAddress := sampleUser.Address

	t.Run("Should fail to fetch user by id due to forbidden", func(t *testing.T) {
		jwt := getAuthToken()

		response := makeRequest[any](http.MethodGet, makeUrl("/user/2"), nil, map[string]string{
			"Authorization": jwt,
		})

		assert.Equal(http.StatusForbidden, response.HttpStatus)
	})

	t.Run("Should successfully fetch user by id", func(t *testing.T) {
		grantRole(getAuthToken(), READ)

		response := findUserById("1", getAuthToken())
		user := response.Content

		assert.Equal(http.StatusOK, response.HttpStatus)

		assert.Equal(int64(1), user.Id)
		assert.Equal(sampleUser.Name, user.Name)
		assert.Equal(sampleUser.Age, user.Age)
		assert.Equal(sampleUser.Email, user.Email)

		assert.Equal(1, len(user.AddressList))
		address := user.AddressList[0]

		assert.Equal(int64(1), address.Id)
		assert.Equal(sampleAddress.Alias, address.AddressAlias)
		assert.Equal(sampleAddress.ZipCode, address.ZipCode)
		assert.Equal(sampleAddress.StreetName, address.StreetName)
		assert.Equal(sampleAddress.Number, address.Number)
		assert.Equal(sampleAddress.State, address.State)
		assert.Equal(sampleAddress.Country, address.Country)
	})

	t.Run("Should fail due to user not found", func(t *testing.T) {
		response := findUserById("2", getAuthToken())

		assert.Equal(http.StatusNotFound, response.HttpStatus)
	})
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should successfully update user", func(t *testing.T) {
		grantRole(getAuthToken(), UPDATE)
		jwt := getAuthToken()
		user := findUserById("1", jwt).Content
		updateDto := user
		updateDto.Age = 40

		response := makeRequest[models.User](http.MethodPut, makeUrl("/user"), updateDto, makeAuthenticationHeader(jwt))
		updatedUser := response.Content

		assert.Equal(http.StatusOK, response.HttpStatus)
		assert.Equal(user.Id, updatedUser.Id)
		assert.NotEqual(user.Age, updatedUser.Age)
		assert.NotNil(updatedUser.UpdatedAt)
	})

	t.Run("Should fail due to user not found", func(t *testing.T) {
		jwt := getAuthToken()
		user := findUserById("1", jwt).Content
		updateDto := user
		updateDto.Id = 2
		updateDto.Age = 40
		updateDto.Email = "johndoe@notsogeneric.com"

		response := makeRequest[models.User](http.MethodPut, makeUrl("/user"), updateDto, makeAuthenticationHeader(getAuthToken()))
		assert.Equal(http.StatusNotFound, response.HttpStatus)
	})

	t.Run("Should fail due no new data found", func(t *testing.T) {
		jwt := getAuthToken()
		user := findUserById("1", jwt).Content

		response := makeRequest[models.User](http.MethodPut, makeUrl("/user"), user, makeAuthenticationHeader(getAuthToken()))
		assert.Equal(http.StatusNotFound, response.HttpStatus)
	})
}

func TestCreateAddress(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should successfully add a new address to the user", func(t *testing.T) {
		grantRole(getAuthToken(), CREATE)

		jwt := getAuthToken()

		newAddressDto := makeNewAddressDto()
		response := makeRequest[models.Address](http.MethodPost, makeUrl("/address"), newAddressDto, makeAuthenticationHeader(jwt))

		assert.Equal(http.StatusOK, response.HttpStatus)
	})

	t.Run("User now should have two active addresses", func(t *testing.T) {
		user := findUserById("1", getAuthToken())

		assert.Equal(2, len(user.Content.AddressList))
	})

	t.Run("Should fail due to user not found", func(t *testing.T) {
		jwt := getAuthToken()

		newAddressDto := makeNewAddressDto()
		newAddressDto.UserId = 99
		response := makeRequest[models.Address](http.MethodPost, makeUrl("/address"), newAddressDto, makeAuthenticationHeader(jwt))

		assert.Equal(http.StatusNotFound, response.HttpStatus)
	})
}

var remainerAddressId int64

func TestDeleteAddress(t *testing.T) {
	assert := assert.New(t)
	var deletedAddressId int64

	t.Run("Should delete the address with no errors", func(t *testing.T) {
		jwt := getAuthToken()
		grantRole(jwt, DELETE)
		jwt = getAuthToken()

		user := findUserById("1", jwt).Content
		address := user.AddressList[0]

		response := deleteAddress(jwt, address.Id)
		assert.Equal(http.StatusOK, response.HttpStatus)

		user = findUserById("1", jwt).Content
		assert.Equal(1, len(user.AddressList))

		deletedAddressId = address.Id
		remainerAddressId = user.AddressList[0].Id
	})

	t.Run("Should fail due to address inactive/not found", func(t *testing.T) {
		jwt := getAuthToken()

		response := deleteAddress(jwt, deletedAddressId)
		assert.Equal(http.StatusNotFound, response.HttpStatus)

		response2 := deleteAddress(jwt, 99)
		assert.Equal(http.StatusNotFound, response2.HttpStatus)
	})
}

func TestUpdateAddress(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should successfully update the address", func(t *testing.T) {
		jwt := getAuthToken()
		user := findUserById("1", jwt)
		address := user.Content.AddressList[0]
		addressUpdateDto := address
		addressUpdateDto.AddressAlias = "Work"

		response := makeRequest[models.Address](http.MethodPut, makeUrl("/address"), addressUpdateDto, makeAuthenticationHeader(jwt))
		updatedAddress := response.Content

		assert.Equal(http.StatusOK, response.HttpStatus)
		assert.Equal(address.Id, updatedAddress.Id)
		assert.NotEqual(address.AddressAlias, updatedAddress.AddressAlias)
		assert.NotNil(updatedAddress.CreatedAt)
	})
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should fail due to user not found", func(t *testing.T) {
		jwt := getAuthToken()
		resp := deleteUser("2", jwt)

		assert.Equal(http.StatusNotFound, resp.HttpStatus)
	})

	t.Run("Should successfully delete the requested user", func(t *testing.T) {
		jwt := getAuthToken()

		resp := deleteUser("1", jwt)
		resp2 := findUserById("1", jwt)

		assert.Equal(http.StatusOK, resp.HttpStatus)
		assert.Equal(http.StatusNotFound, resp2.HttpStatus)
	})
}

func deleteUser(id, jwt string) httputils.HttpResponse[any] {
	return makeRequest[any](http.MethodDelete, makeUrl("/user/"+id), nil, makeAuthenticationHeader(jwt))
}

func findUserById(id, jwt string) httputils.HttpResponse[models.User] {
	return makeRequest[models.User](http.MethodGet, makeUrl("/user/"+id), nil, makeAuthenticationHeader(jwt))
}

func grantRole(jwt string, roleId int64) {
	makeRequest[any](http.MethodPost, makeUrl("/auth/authorize"), dto.AuthorizeUserDto{
		UserId: 1,
		RoleId: roleId,
	}, makeAuthenticationHeader(jwt))
}

func makeAuthenticationHeader(token string) map[string]string {
	return map[string]string{
		"Authorization": token,
	}
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
		Email:    EMAIL_1,
		Password: "password+123",
		Address:  makeNewAddressDto(),
	}
}

func makeNewAddressDto() dto.CreateAddressDto {
	return dto.CreateAddressDto{
		Alias:      "Home",
		ZipCode:    "00000",
		StreetName: "Generic Street",
		Number:     "007",
		State:      "GS",
		Country:    "GC",
		UserId:     1,
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

func deleteAddress(jwt string, addressId int64) httputils.HttpResponse[any] {
	return makeRequest[any](http.MethodDelete, makeUrl("/address/"+fmt.Sprint(addressId)), nil, makeAuthenticationHeader(jwt))
}
