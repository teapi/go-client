package teapi

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type ConfigurationTests struct{}

func Test_Configuration(t *testing.T) {
	Expectify(new(ConfigurationTests), t)
}

func (_ ConfigurationTests) NormalizesTheHost() {
	for _, h := range []string{"m01.teapi.io", "m01.teapi.io/", "https://m01.teapi.io", "https://m01.teapi.io/"} {
		Expect(Configure(h, "", "").host).To.Equal("https://m01.teapi.io")
	}
}
