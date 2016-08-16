package helpers

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var SkipSSLValidation bool

func Curl(args ...string) *gexec.Session {
	return CurlSkipSSL(SkipSSLValidation, args...)
}

func CurlSkipSSL(skip bool, args ...string) *gexec.Session {
	curlArgs := append([]string{"-s"}, args...)
	if skip {
		curlArgs = append([]string{"-k"}, curlArgs...)
	}
	cmdStarter := runner.NewCommandStarter()
	request, err := cmdStarter.Start("curl", curlArgs...)
	Expect(err).NotTo(HaveOccurred())

	return request
}