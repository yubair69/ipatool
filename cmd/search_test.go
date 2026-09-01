package cmd

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Search command", func() {
	It("is registered on the root command", func() {
		cmd, _, err := rootCmd().Find([]string{"search"})

		Expect(err).ToNot(HaveOccurred())
		Expect(cmd.Name()).To(Equal("search"))
	})

	DescribeTable("rejects invalid limit values",
		func(value string, errorSubstring string) {
			cmd := searchCmd()
			Expect(cmd.Flags().Set("limit", value)).To(Succeed())

			err := cmd.PreRunE(cmd, nil)
			Expect(err).To(MatchError(ContainSubstring(errorSubstring)))
		},
		Entry("limit zero", "0", "greater than 0"),
		Entry("limit negative", "-5", "greater than 0"),
	)

	It("accepts positive limit values", func() {
		cmd := searchCmd()
		Expect(cmd.Flags().Set("limit", "5")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("accepts maximum limit for visionOS", func() {
		cmd := searchCmd()
		Expect(cmd.Flags().Set("limit", "12")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})
})
