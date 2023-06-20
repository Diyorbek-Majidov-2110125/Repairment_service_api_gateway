package handlers

import (
	"projects/Repairment_service/Repairment_service_api_gateway/api/http"
	"projects/Repairment_service/Repairment_service_api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// Create user godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create user
// @Description Create user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=corporate_service.CreateAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var user user_service.CreateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get user by id
// @ID get_user_by_id
// @Router /user/{id} [GET]
// @Summary Get user by id
// @Description Get user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=corporate_service.GetAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.UserService().GetById(
		c.Request.Context(),
		&user_service.UserPrimaryKey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get user list
// @ID get_user_list
// @Router /user [GET]
// @Summary Get user list
// @Description Get user list
// @Tags User
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param sort_by query string false "sort_by"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param order query string false "order"
// @Success 200 {object} http.Response{data=corporate_service.GetAllAgentsResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllUsers(c *gin.Context) {
	resp, err := h.services.UserService().GetList(
		c.Request.Context(),
		&user_service.GetListUserRequest{
			Offset:  0,
			Limit: 0,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete user by id
// @ID delete_user_by_id
// @Router /user/{id} [DELETE]
// @Summary Delete user by id
// @Description Delete user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.UserService().Delete(
		c.Request.Context(),
		&user_service.UserPrimaryKey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body user_service.Update true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=user_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var user user_service.UpdateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Update(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
