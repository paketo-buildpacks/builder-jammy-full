package smoke_test

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/onsi/gomega/format"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var Builder string

var BuildpackURIs map[string]string

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSmoke(t *testing.T) {
	format.MaxLength = 0
	Expect := NewWithT(t).Expect

	flag.Parse()

	Builder = "paketobuildpacks/builder-jammy-buildpackless-full:latest"

	Expect(Builder).NotTo(Equal(""))

	data, err := os.ReadFile("../builder.toml")
	Expect(err).NotTo(HaveOccurred())

	re := regexp.MustCompile(`uri = "docker://docker.io/paketobuildpacks/([^:]+):([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(data), -1)
	Expect(matches).NotTo(BeEmpty())

	BuildpackURIs = make(map[string]string, len(matches))
	for _, m := range matches {
		BuildpackURIs[m[1]] = fmt.Sprintf("docker://docker.io/paketobuildpacks/%s:%s", m[1], m[2])
	}

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Smoke", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("Dotnet", testDotnet)
	suite("Go", testGo)
	suite("Java Native Image", testJavaNativeImage)
	suite("Java", testJava)
	suite("Node.js", testNodejs)
	suite("Procfile", testProcfile)
	suite("PHP", testPhp)
	suite("Python", testPython)
	suite("Ruby", testRuby)
	suite("Web Servers", testWebServers)
	suite.Run(t)
}
