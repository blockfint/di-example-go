package repository

import (
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Todo Repository", func() {
	gormDB, mock, err := newMockedDB()
	lg := logger.New()

	todoRepository := NewTodoRepository(gormDB, lg)

	if err != nil {
		panic(err)
	}

	Describe("Getting todo", func() {
		It("should find by id", func() {
			var (
				id   uint   = 1
				name string = "test-todo"
			)

			mock.
				ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "todo"
					WHERE "todo"."id" = $1 AND "todo"."deleted_at" IS NULL
					ORDER BY "todo"."id" LIMIT 1`,
				)).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).AddRow(id, time.Now(), time.Now(), nil, name))

			_, err := todoRepository.FindByID(id)
			Expect(err).To(BeNil())
		})
	})

	Describe("Creating todo", func() {
		It("should insert", func() {
			todo := model.Todo{
				Name: "Buy a cup of latte",
			}

			mock.ExpectBegin()
			mock.
				ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "todo" ("created_at","updated_at","deleted_at","name")
			   	VALUES ($1,$2,$3,$4)
				  RETURNING "id"`,
				)).
				WithArgs(AnyTime{}, AnyTime{}, nil, todo.Name).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			mock.ExpectCommit()

			err := todoRepository.Create(&todo)
			Expect(err).To(BeNil())

			err = mock.ExpectationsWereMet()
			Expect(err).To(BeNil())
		})
	})
})
