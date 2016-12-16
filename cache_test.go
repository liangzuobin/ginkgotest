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
