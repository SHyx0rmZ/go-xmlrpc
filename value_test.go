package xmlrpc_test

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Value", func() {
	var (
		server *ghttp.Server
		client xmlrpc.Client
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = xmlrpc.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	It("Can encode slices", func() {
		request := `<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><array><data><value><string>foo</string></value><value><string>bar</string></value></data></array></value></param></params></methodCall>`
		response := `<?xml version="1.0"?><methodResponse><params><param><boolean>true</boolean></param></params></methodResponse>`

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyBody([]byte(request)),
				ghttp.RespondWith(200, []byte(response)),
			),
		)

		_, err := client.Call("test", []interface{}{"foo", "bar"})

		Expect(err).To(BeNil())
	})

	It("Can decode slices", func() {
		request := `<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`
		response := `<?xml version="1.0"?><methodResponse><params><param><value><array><data><value><string>foo</string></value><value><string>bar</string></value></data></array></value></param></params></methodResponse>`

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyBody([]byte(request)),
				ghttp.RespondWith(200, []byte(response)),
			),
		)

		val, err := client.Call("test")

		Expect(err).To(BeNil())
		Expect(len(val.AsArray())).To(Equal(2))
		Expect(val.AsArray()[0].AsString()).To(Equal("foo"))
		Expect(val.AsArray()[1].AsString()).To(Equal("bar"))
	})
})
