package controller

import (
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type mockedTodoRepository struct {
	todos *[]model.Todo
}

func (m *mockedTodoRepository) List() (*[]model.Todo, error) {
	return m.todos, nil
}

func (m *mockedTodoRepository) Create(todo *model.Todo) error {
	*m.todos = append(*m.todos, *todo)
	return nil
}

func (m *mockedTodoRepository) FindByID(todoID uint) (*model.Todo, error) {
	for _, todo := range *m.todos {
		if todo.ID == todoID {
			return &todo, nil
		}
	}

	return nil, nil
}

func newMockedTodoController(repo *mockedTodoRepository) *TodoController {
	lg := logger.New()

	return &TodoController{repo, lg}
}

var _ = Describe("Todo Controller", func() {
	todos := &[]model.Todo{}
	todoController := newMockedTodoController(&mockedTodoRepository{todos})

	Describe("Listing todos", func() {
		When("the todo repository is empty", func() {
			It("should be an empty list", func() {
				todos, err := todoController.List()

				Expect(len(*todos)).To(Equal(0))
				Expect(err).To(BeNil())
			})
		})

		When("the todo repository has todos", func() {
			BeforeEach(func() {
				*todos = []model.Todo{
					{Name: "Buy an apple"},
					{Name: "Buy an orange"},
				}
			})

			It("should be 2 todos in the list", func() {
				todos, err := todoController.List()

				Expect(len(*todos)).To(Equal(2))
				Expect((*todos)[0].Name).To(Equal("Buy an apple"))
				Expect((*todos)[1].Name).To(Equal("Buy an orange"))
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Creating todo", func() {
		BeforeEach(func() {
			*todos = []model.Todo{}
		})

		It("should create todo", func() {
			err := todoController.Create(&model.Todo{Name: "Buy a doughnut"})
			Expect(err).To(BeNil())

			todos, err := todoController.List()
			Expect(len(*todos)).To(Equal(1))
			Expect((*todos)[0].Name).To(Equal("Buy a doughnut"))
		})
	})

	Describe("Find todo by ID", func() {
		BeforeEach(func() {
			*todos = []model.Todo{
				{
					BaseModel: model.BaseModel{
						ID: 1,
					},
					Name: "Buy a house",
				},
				{
					BaseModel: model.BaseModel{
						ID: 2,
					},
					Name: "Buy a car",
				},
			}
		})

		It("should find a todo", func() {
			todo, err := todoController.FindByID(2)

			Expect(todo.Name).To(Equal("Buy a car"))
			Expect(err).To(BeNil())
		})
	})
})
