package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/logger"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type mockedTodoController struct {
	todos *[]model.Todo
}

func (m *mockedTodoController) List() (*[]model.Todo, error) {
	return m.todos, nil
}

func (m *mockedTodoController) Create(todo *model.Todo) error {
	*m.todos = append(*m.todos, *todo)
	return nil
}

func (m *mockedTodoController) FindByID(todoID uint) (*model.Todo, error) {
	for _, todo := range *m.todos {
		if todo.ID == todoID {
			return &todo, nil
		}
	}

	return nil, nil
}

func newMockedTodoHandler(con *mockedTodoController) *Handler {
	lg := logger.New()

	return &Handler{con, nil, lg}
}

var _ = Describe("Todo Handler", func() {
	todos := &[]model.Todo{}

	h := newMockedTodoHandler(&mockedTodoController{todos})

	Describe("Creating Todo", func() {
		It("should create todo", func() {
			todoJSON := `{"name":"Buy a cup of mocha"}`

			e := echo.New()
			e.Validator = NewRequestValidator()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(todoJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.CreateTodo(c)

			Expect(err).To(BeNil())

			var r createTodoRequest
			json.Unmarshal([]byte(todoJSON), &r)

			var createdTodo model.Todo
			json.Unmarshal([]byte(rec.Body.String()), &createdTodo)

			Expect(createdTodo.ID).To(BeNumerically(">=", 0))
			Expect(createdTodo.Name).To(Equal(r.Name))
		})
	})
})
