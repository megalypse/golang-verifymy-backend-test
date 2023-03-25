package address

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/address"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/dto"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/roles"
)

type AddressController struct {
	CreateAddressUsecase address.CreateAddress
	UpdateAddressUsecase address.UpdateAddress
	DeleteAddressUsecase address.DeleteAddress
}

func (ac AddressController) GetHandlers() []httputils.RouteDefinition {
	return []httputils.RouteDefinition{
		{
			Method:        http.MethodPost,
			Route:         "/address",
			RequiredRoles: []string{roles.UPDATE, roles.READ},
			HandlingFunc:  ac.createAddress,
		},
		{
			Method:        http.MethodPut,
			Route:         "/address",
			RequiredRoles: []string{roles.CREATE, roles.READ},
			HandlingFunc:  ac.updateAddress,
		},
		{
			Method:        http.MethodDelete,
			Route:         "/address/{addressId}",
			RequiredRoles: []string{roles.DELETE},
			HandlingFunc:  ac.deleteAddress,
		},
	}
}

// @Summary Updates the selected address
// @Tags Address
// @Success 200 {object} httputils.HttpResponse[models.Address]
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body models.Address true "Address model"
// @Router /address [put]
// @Security ApiKeyAuth
func (ac AddressController) updateAddress(w http.ResponseWriter, r *http.Request) {
	req, err := httputils.ParseRequest[models.Address](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	address, err := ac.UpdateAddressUsecase.Update(r.Context(), req.Body)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[models.Address]{
		HttpStatus: http.StatusOK,
		Message:    "Address successfully updated",
		Content:    *address,
	})
}

// @Summary Deletes the selected address
// @Tags Address
// @Success 200 {object} httputils.HttpResponse[string]
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param addressId path int true "Address id"
// @Router /address/{addressId} [delete]
// @Security ApiKeyAuth
func (ac AddressController) deleteAddress(w http.ResponseWriter, r *http.Request) {
	param := "addressId"
	req, err := httputils.ParseRequest[httputils.Void](r, &[]string{param})
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	addressId, err := httputils.ParseId(req.Params[param])
	if err != nil {
		httputils.WriteError(w, &models.CustomError{
			Code:    err.Code,
			Message: err.Message,
			Source:  err.Source,
		})
		return
	}

	if err = ac.DeleteAddressUsecase.Delete(r.Context(), addressId); err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[string]{
		HttpStatus: http.StatusOK,
		Message:    "Address successfully deleted",
	})
}

// @Summary Creates a new address
// @Tags Address
// @Success 201 {object} httputils.HttpResponse[models.Address]
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body dto.CreateAddressDto true "Create address dto"
// @Router /address [post]
// @Security ApiKeyAuth
func (ac AddressController) createAddress(w http.ResponseWriter, r *http.Request) {
	req, err := httputils.ParseRequest[dto.CreateAddressDto](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	updatedAddress, err := ac.CreateAddressUsecase.Create(r.Context(), req.Body.ToAddress())
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[models.Address]{
		HttpStatus: http.StatusOK,
		Content:    *updatedAddress,
		Message:    "Address successfully updated",
	})
}
