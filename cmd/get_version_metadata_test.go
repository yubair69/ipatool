package cmd

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Version Metadata command", func() {
	It("is registered on the root command", func() {
		cmd, _, err := rootCmd().Find([]string{"get-version-metadata"})

		Expect(err).ToNot(HaveOccurred())
		Expect(cmd.Name()).To(Equal("get-version-metadata"))
	})

	DescribeTable("rejects invalid app ID values when no bundle identifier given",
		func(appIDValue string, errorSubstring string) {
			cmd := getVersionMetadataCmd()
			Expect(cmd.Flags().Set("app-id", appIDValue)).To(Succeed())

			err := cmd.PreRunE(cmd, nil)
			Expect(err).To(MatchError(ContainSubstring(errorSubstring)))
		},
		Entry("app ID zero with no bundle ID", "0", "either the app ID or the bundle identifier"),
		Entry("app ID negative", "-1", "positive number"),
	)

	It("accepts positive app ID values", func() {
		cmd := getVersionMetadataCmd()
		Expect(cmd.Flags().Set("app-id", "555555")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("accepts zero app ID when bundle identifier is provided", func() {
		cmd := getVersionMetadataCmd()
		Expect(cmd.Flags().Set("app-id", "0")).To(Succeed())
		Expect(cmd.Flags().Set("bundle-identifier", "com.example.app")).To(Succeed())

		err := cmd.PreRunE(cmd, nil)
		Expect(err).ToNot(HaveOccurred())
	})
})
