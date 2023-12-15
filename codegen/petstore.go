package main

import (
	"kbtg-bootcamp-petstore/petstore"

	"github.com/labstack/echo/v4"
)

type serverWrapper struct {
	Handler petstore.ServerInterface
}

// (POST /pet)
func (w *serverWrapper) AddPet(ctx echo.Context) error {
	return ctx.String(200, "Added")
}

// Update an existing pet
// (PUT /pet)
func (w *serverWrapper) UpdatePet(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Finds Pets by status
// (GET /pet/findByStatus)
func (w *serverWrapper) FindPetsByStatus(ctx echo.Context, params petstore.FindPetsByStatusParams) error {
	return echo.ErrNotImplemented
}

// Finds Pets by tags
// (GET /pet/findByTags)
func (w *serverWrapper) FindPetsByTags(ctx echo.Context, params petstore.FindPetsByTagsParams) error {
	return echo.ErrNotImplemented
}

// Deletes a pet
// (DELETE /pet/{petId})
func (w *serverWrapper) DeletePet(ctx echo.Context, petId int64, params petstore.DeletePetParams) error {
	return echo.ErrNotImplemented
}

// Find pet by ID
// (GET /pet/{petId})
func (w *serverWrapper) GetPetById(ctx echo.Context, petId int64) error {
	return echo.ErrNotImplemented
}

// Updates a pet in the store with form data
// (POST /pet/{petId})
func (w *serverWrapper) UpdatePetWithForm(ctx echo.Context, petId int64, params petstore.UpdatePetWithFormParams) error {
	return echo.ErrNotImplemented
}

// uploads an image
// (POST /pet/{petId}/uploadImage)
func (w *serverWrapper) UploadFile(ctx echo.Context, petId int64, params petstore.UploadFileParams) error {
	return echo.ErrNotImplemented
}

// Returns pet inventories by status
// (GET /store/inventory)
func (w *serverWrapper) GetInventory(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Place an order for a pet
// (POST /store/order)
func (w *serverWrapper) PlaceOrder(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Delete purchase order by ID
// (DELETE /store/order/{orderId})
func (w *serverWrapper) DeleteOrder(ctx echo.Context, orderId int64) error {
	return echo.ErrNotImplemented
}

// Find purchase order by ID
// (GET /store/order/{orderId})
func (w *serverWrapper) GetOrderById(ctx echo.Context, orderId int64) error {
	return echo.ErrNotImplemented
}

// Create user
// (POST /user)
func (w *serverWrapper) CreateUser(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Creates list of users with given input array
// (POST /user/createWithList)
func (w *serverWrapper) CreateUsersWithListInput(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Logs user into the system
// (GET /user/login)
func (w *serverWrapper) LoginUser(ctx echo.Context, params petstore.LoginUserParams) error {
	return echo.ErrNotImplemented
}

// Logs out current logged in user session
// (GET /user/logout)
func (w *serverWrapper) LogoutUser(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Delete user
// (DELETE /user/{username})
func (w *serverWrapper) DeleteUser(ctx echo.Context, username string) error {
	return echo.ErrNotImplemented
}

// Get user by user name
// (GET /user/{username})
func (w *serverWrapper) GetUserByName(ctx echo.Context, username string) error {
	return echo.ErrNotImplemented
}

// Update user
// (PUT /user/{username})
func (w *serverWrapper) UpdateUser(ctx echo.Context, username string) error {
	return echo.ErrNotImplemented
}
