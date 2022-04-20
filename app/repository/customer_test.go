package repository

import (
	"github.com/cockroachdb/copyist"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Repository", func() {
	copyist.Register("postgres")

	Describe("Getting list of customers", func() {
		It("should list all customers", func() {
			defer copyist.Open(GinkgoT()).Close()

			customerRepository, err := newCustomerRepositoryWithMockedDB()
			Expect(err).To(BeNil())

			_, err = customerRepository.List()

			Expect(err).To(BeNil())
		})
	})

	Describe("Getting customer", func() {
		It("should find by id", func() {
			defer copyist.Open(GinkgoT()).Close()

			customerRepository, err := newCustomerRepositoryWithMockedDB()
			Expect(err).To(BeNil())

			var id uint = 1
			_, err = customerRepository.FindByID(id)

			Expect(err).To(BeNil())
		})
	})
})
