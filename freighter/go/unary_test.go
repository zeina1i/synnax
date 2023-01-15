// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package freighter_test

import (
	"context"
	"github.com/synnaxlabs/freighter"
	"github.com/synnaxlabs/freighter/fhttp"
	"github.com/synnaxlabs/x/address"
	"github.com/synnaxlabs/x/httputil"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type (
	unaryServer = freighter.UnaryServer[request, response]
	unaryClient = freighter.UnaryClient[request, response]
)

type unaryImplementation interface {
	start(host address.Address, logger *zap.SugaredLogger) (unaryServer, unaryClient)
	stop() error
}

var unaryImplementations = []unaryImplementation{
	&httpUnaryImplementation{},
}

var _ = Describe("Unary", Ordered, Serial, func() {
	Describe("Implementation Tests", func() {
		for _, impl := range unaryImplementations {
			impl := impl
			var (
				addr   address.Address
				server unaryServer
				client unaryClient
			)
			BeforeAll(func() {
				addr = "localhost:8080"
				l := zap.NewNop()
				server, client = impl.start(addr, l.Sugar())
			})
			AfterAll(func() {
				Expect(impl.stop()).To(Succeed())
			})
			Describe("Normal Operation", func() {
				It("should send a request", func() {
					server.BindHandler(func(ctx context.Context, req request) (response, error) {
						return response{
							ID:      req.ID,
							Message: req.Message,
						}, nil
					})
					req := request{ID: 1, Message: "hello"}
					res, err := client.Send(context.TODO(), addr, req)
					Expect(err).To(Succeed())
					Expect(res).To(Equal(response{ID: 1, Message: "hello"}))
				})
			})
			Describe("Err Handling", func() {
				It("Should correctly return a custom error to the client", func() {
					server.BindHandler(func(ctx context.Context, req request) (response, error) {
						return response{}, myCustomError
					})
					req := request{ID: 1, Message: "hello"}
					_, err := client.Send(context.TODO(), addr, req)
					Expect(err).To(Equal(myCustomError))
				})
			})
			Describe("Middleware", func() {
				It("Should correctly call the middleware", func() {
					c := 0
					server.Use(freighter.MiddlewareFunc(func(ctx context.Context, md freighter.MD, next freighter.Next) (freighter.MD, error) {
						c++
						oMd, err := next(ctx, md)
						c++
						return oMd, err
					}))
					server.BindHandler(func(ctx context.Context, req request) (response, error) {
						return response{}, nil
					})
					req := request{ID: 1, Message: "hello"}
					_, err := client.Send(context.TODO(), addr, req)
					Expect(err).To(Succeed())
					Expect(c).To(Equal(2))
				})
			})
		}
	})
})

type httpUnaryImplementation struct {
	app *fiber.App
}

func (h *httpUnaryImplementation) start(host address.Address, logger *zap.SugaredLogger) (unaryServer, unaryClient) {
	h.app = fiber.New(fiber.Config{DisableStartupMessage: true})
	router := fhttp.NewRouter(fhttp.RouterConfig{Logger: logger})
	factory := fhttp.NewClientFactory(fhttp.ClientFactoryConfig{
		EncoderDecoder: httputil.JSONEncoderDecoder,
		Logger:         logger,
	})
	server := fhttp.UnaryPostServer[request, response](router, "/")
	client := fhttp.UnaryPostClient[request, response](factory)
	router.BindTo(h.app)
	h.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	go func() {
		if err := h.app.Listen(host.PortString()); err != nil {
			logger.Error(err)
		}
	}()
	Eventually(func(g Gomega) {
		_, err := http.Get("http://" + host.String() + "/health")
		g.Expect(err).ToNot(HaveOccurred())
	}).WithPolling(1 * time.Millisecond).Should(Succeed())
	return server, client
}

func (h *httpUnaryImplementation) stop() error {
	return h.app.Shutdown()
}
