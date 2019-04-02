package envs_test

import (
	"github.com/mapreal19/beemiel/envs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("envs", func() {
	Describe("IsProduction", func() {
		It("uses BEEGO_ENV key by default", func() {
			_ = os.Setenv("BEEGO_ENV", "production")

			Expect(envs.IsProduction()).To(BeTrue())
		})

		It("returns true if we are in production", func() {
			envs.Init("MY_APP_ENV")
			_ = os.Setenv("MY_APP_ENV", "production")

			Expect(envs.IsProduction()).To(BeTrue())
		})

		It("returns false if we aren't in production", func() {
			envs.Init("MY_APP_ENV")
			_ = os.Setenv("MY_APP_ENV", "test")

			Expect(envs.IsProduction()).To(BeFalse())
		})
	})

	Describe("IsTest", func() {
		It("returns true if we are in test or false otherwise", func() {
			envs.Init("MY_APP_ENV")
			_ = os.Setenv("MY_APP_ENV", "test")

			Expect(envs.IsTest()).To(BeTrue())

			_ = os.Setenv("MY_APP_ENV", "development")

			Expect(envs.IsTest()).To(BeFalse())
		})
	})

	Describe("IsDevelopment", func() {
		It("returns true if we are in test or false otherwise", func() {
			envs.Init("MY_APP_ENV")
			_ = os.Setenv("MY_APP_ENV", "development")

			Expect(envs.IsDevelopment()).To(BeTrue())

			_ = os.Setenv("MY_APP_ENV", "test")

			Expect(envs.IsDevelopment()).To(BeFalse())
		})
	})

	Describe("IsCI", func() {
		It("returns true if we are in test or false otherwise", func() {
			envs.Init("MY_APP_ENV")
			_ = os.Setenv("MY_APP_ENV", "ci")

			Expect(envs.IsCI()).To(BeTrue())

			_ = os.Setenv("MY_APP_ENV", "test")

			Expect(envs.IsCI()).To(BeFalse())
		})
	})
})
