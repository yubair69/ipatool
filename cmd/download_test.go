package cmd

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Download command", func() {
	It("is registered on the root command", func() {
		cmd, _, err := rootCmd().Find([]string{"download"})

		Expect(err).ToNot(HaveOccurred())
		Expect(cmd.Name()).To(Equal("download"))
	})

	DescribeTable("rejects invalid app ID values when no bundle identifier given",
		func(appIDValue string, errorSubstring string) {
			cmd := downloadCmd()
			Expect(cmd.Flags().Set("app-id", appIDValue)).To(Succeed())

			err := cmd.PreRunE(cmd, nil)
			Expect(err).To(MatchError(ContainSubstring(errorSubstring)))
		},
		Entry("app ID zero with no bundle ID", "0", "either the app ID or the bundle identifier"),
		Entry("app ID negative", "-42", "positive number"),
	)

	It("accepts positive app ID values", func() {
		cmd := downloadCmd()
		Expect(cmd.Flags().Set("app-id", "123456")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("accepts zero app ID when bundle identifier is provided", func() {
		cmd := downloadCmd()
		Expect(cmd.Flags().Set("app-id", "0")).To(Succeed())
		Expect(cmd.Flags().Set("bundle-identifier", "com.example.app")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("allows bundle identifier to override app ID", func() {
		cmd := downloadCmd()
		Expect(cmd.Flags().Set("app-id", "0")).To(Succeed())
		Expect(cmd.Flags().Set("bundle-identifier", "com.example.app")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})
})
