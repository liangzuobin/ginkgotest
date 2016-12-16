package cache

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("start testing cache", func() {
	BeforeSuite(func() {
		Expect(StartRedis()).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		Expect(StopRedis()).NotTo(HaveOccurred())
	})

	It("run Config and no error occorred", func() {
		Expect(Config("localhost:6379", 10)).NotTo(HaveOccurred())
	})

	It("ping with client", func() {
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})
})

var _ = Describe("just before each test", func() {
	var (
		book Book
		err  error
		json string
	)

	JustBeforeEach(func() {
		book, err = NewBookFromJSON(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			BeforeEach(func() {
				json = `{
				    "title":"Les Miserables",
				    "author":"Victor Hugo",
				    "pages":1488
				}`
			})

			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				json = `{
				    "title":"Les Miserables",
				    "author":"Victor Hugo",
				    "pages":1488oops
				}`
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
