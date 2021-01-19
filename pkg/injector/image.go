// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package injector

import (
	"fmt"

	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type options struct {
	name     []name.Option
	remote   []remote.Option
	platform *v1.Platform
}

func makeOptions(opts ...Option) options {
	opt := options{
		remote: []remote.Option{
			remote.WithAuthFromKeychain(authn.DefaultKeychain),
		},
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Option is a functional option for crane.
type Option func(*options)

// WithTransport is a functional option for overriding the default transport
// for remote operations.
func WithTransport(t http.RoundTripper) Option {
	return func(o *options) {
		o.remote = append(o.remote, remote.WithTransport(t))
	}
}

// Insecure is an Option that allows image references to be fetched without TLS.
func Insecure(o *options) {
	o.name = append(o.name, name.Insecure)
}

// WithPlatform is an Option to specify the platform.
func WithPlatform(platform *v1.Platform) Option {
	return func(o *options) {
		if platform != nil {
			o.remote = append(o.remote, remote.WithPlatform(*platform))
		}
		o.platform = platform
	}
}

// WithAuthFromKeychain is a functional option for overriding the default
// authenticator for remote operations, using an authn.Keychain to find
// credentials.
//
// By default, crane will use authn.DefaultKeychain.
func WithAuthFromKeychain(keys authn.Keychain) Option {
	return func(o *options) {
		// Replace the default keychain at position 0.
		o.remote[0] = remote.WithAuthFromKeychain(keys)
	}
}

// WithAuth is a functional option for overriding the default authenticator
// for remote operations.
//
// By default, crane will use authn.DefaultKeychain.
func WithAuth(auth authn.Authenticator) Option {
	return func(o *options) {
		// Replace the default keychain at position 0.
		o.remote[0] = remote.WithAuth(auth)
	}
}

// WithUserAgent adds the given string to the User-Agent header for any HTTP
// requests.
func WithUserAgent(ua string) Option {
	return func(o *options) {
		o.remote = append(o.remote, remote.WithUserAgent(ua))
	}
}

func getImage(r string, opt ...Option) (v1.Image, name.Reference, error) {
	o := makeOptions(opt...)
	ref, err := name.ParseReference(r, o.name...)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing reference %q: %v", r, err)
	}
	img, err := remote.Image(ref, o.remote...)
	if err != nil {
		return nil, nil, fmt.Errorf("reading image %q: %v", ref, err)
	}
	return img, ref, nil
}

func getManifest(r string, opt ...Option) (*remote.Descriptor, error) {
	o := makeOptions(opt...)
	ref, err := name.ParseReference(r, o.name...)
	if err != nil {
		return nil, fmt.Errorf("parsing reference %q: %v", r, err)
	}
	return remote.Get(ref, o.remote...)
}

func head(r string, opt ...Option) (*v1.Descriptor, error) {
	o := makeOptions(opt...)
	ref, err := name.ParseReference(r, o.name...)
	if err != nil {
		return nil, err
	}
	return remote.Head(ref, o.remote...)
}
