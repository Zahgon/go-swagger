// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

// ServeCmd to serve a swagger spec with docs ui.
type ServeCmd struct {
	BasePath string `description:"the base path to serve the spec and UI at"                           long:"base-path"`
	Flavor   string `choice:"redoc"                                                                    choice:"swagger"                                                    default:"redoc" description:"the flavor of docs, can be swagger or redoc" long:"flavor" short:"F"`
	DocURL   string `description:"override the url which takes a url query param to render the doc ui" long:"doc-url"`
	NoOpen   bool   `description:"when present won't open the browser to show the url"                 long:"no-open"`
	NoUI     bool   `description:"when present, only the swagger spec will be served"                  long:"no-ui"`
	Flatten  bool   `description:"when present, flatten the swagger spec before serving it"            long:"flatten"`
	Port     int    `description:"the port to serve this site"                                         env:"PORT"                                                          long:"port"     short:"p"`
	Host     string `default:"0.0.0.0"                                                                 description:"the interface to serve this site, defaults to 0.0.0.0" env:"HOST"      long:"host"`
	Path     string `default:"docs"                                                                    description:"the uri path at which the docs will be served"         long:"path"`
}

// Execute the serve command.
func (s *ServeCmd) Execute(args []string) error { _ = "STUB: not implemented"; return nil }

//nolint:noctx // that's ok for a demo server
