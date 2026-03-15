package cmd

import (
	"embed"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

type Options struct {
	StaticFs map[string]embed.FS
	App      interfaces.AppProvider
}

type Option func(*Options)

func WithTemplateFs(fs embed.FS) Option {
	return func(o *Options) {
		if o.StaticFs == nil {
			o.StaticFs = make(map[string]embed.FS)
		}
		o.StaticFs["templates"] = fs
	}
}

func WithWebStaticFs(fs embed.FS) Option {
	return func(o *Options) {
		if o.StaticFs == nil {
			o.StaticFs = make(map[string]embed.FS)
		}
		o.StaticFs["web.static"] = fs
	}
}

func WithApp(app interfaces.AppProvider) Option {
	return func(o *Options) { o.App = app }
}
