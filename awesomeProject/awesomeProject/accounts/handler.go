package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, name)

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto.PatchAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	name := request.Name

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Name = name

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	name := request.Name
	amount := request.Amount

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount = amount

	return c.NoContent(http.StatusOK)
}

// Написать клиент консольный, который делает запросы
