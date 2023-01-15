// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package security_test

import (
	"crypto/tls"
	"github.com/synnaxlabs/synnax/pkg/security"
	"github.com/synnaxlabs/synnax/pkg/security/cert"
	"github.com/synnaxlabs/synnax/pkg/security/mock"
	"github.com/synnaxlabs/x/config"
	xfs "github.com/synnaxlabs/x/io/fs"
	"os"
)

var _ = Describe("Provider", func() {
	Describe("Secure", func() {
		Describe("TLS Configuration", func() {
			It("Should load and return the correct TLS configuration", func() {
				fs := xfs.NewMem()
				mock.GenerateCerts(fs)
				prov := MustSucceed(security.NewProvider(security.ProviderConfig{
					LoaderConfig: cert.LoaderConfig{FS: fs},
					KeySize:      mock.SmallKeySize,
					Insecure:     config.BoolPointer(false),
				}))
				config := prov.TLS()
				Expect(config).ToNot(BeNil())
				Expect(config.GetCertificate).ToNot(BeNil())
				c := MustSucceed(config.GetCertificate(&tls.ClientHelloInfo{}))
				Expect(c.Certificate).To(HaveLen(1))
				Expect(config.GetClientCertificate).ToNot(BeNil())
				c = MustSucceed(config.GetClientCertificate(&tls.CertificateRequestInfo{}))
				Expect(c.Certificate).To(HaveLen(1))
				Expect(config.RootCAs).ToNot(BeNil())
				Expect(config.ClientAuth).To(Equal(tls.VerifyClientCertIfGiven))
				Expect(config.MinVersion).To(Equal(uint16(tls.VersionTLS12)))
				Expect(config.ClientCAs).ToNot(BeNil())
			})
			It("Should return an error if the node certificate is not found", func() {
				fs := xfs.NewMem()
				_, err := security.NewProvider(security.ProviderConfig{
					LoaderConfig: cert.LoaderConfig{FS: fs},
					KeySize:      mock.SmallKeySize,
					Insecure:     config.BoolPointer(false),
				})
				Expect(err).To(HaveOccurredAs(os.ErrNotExist))
			})
		})
		Describe("Node Private", func() {
			It("Should return the node private key", func() {
				fs := xfs.NewMem()
				mock.GenerateCerts(fs)
				prov := MustSucceed(security.NewProvider(security.ProviderConfig{
					LoaderConfig: cert.LoaderConfig{FS: fs},
					KeySize:      mock.SmallKeySize,
					Insecure:     config.BoolPointer(false),
				}))
				Expect(prov.NodePrivate()).ToNot(BeNil())
			})

		})
	})
	Describe("Insecure", func() {
		Describe("TLS Configuration", func() {
			It("Should return an empty TLS configuration", func() {
				prov := MustSucceed(security.NewProvider(security.ProviderConfig{
					Insecure: config.BoolPointer(true),
					KeySize:  mock.SmallKeySize,
				}))
				Expect(prov.TLS()).To(BeNil())
			})
		})
		Describe("Node Private", func() {
			It("Should return the randomly generated private key", func() {
				prov := MustSucceed(security.NewProvider(security.ProviderConfig{
					Insecure: config.BoolPointer(true),
					KeySize:  mock.SmallKeySize,
				}))
				Expect(prov.NodePrivate()).ToNot(BeNil())
			})
		})
	})
})
