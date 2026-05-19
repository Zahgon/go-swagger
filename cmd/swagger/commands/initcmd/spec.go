// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package initcmd

// Spec represents a command for initializing a new swagger application.
type Spec struct {
	Format      string   `choice:"yaml"                            choice:"json"                                                                   default:"yaml"  description:"the format for the spec document" long:"format"` //nolint:staticcheck // false positive detecting duplicate tags (it works fine on other files with the same pattern)
	Title       string   `description:"the title of the API"       long:"title"`
	Description string   `description:"the description of the API" long:"description"`
	Version     string   `default:"0.1.0"                          description:"the version of the API"                                            long:"version"`
	Terms       string   `description:"the terms of services"      long:"terms"`
	Consumes    []string `default:"application/json"               description:"add a content type to the global consumes definitions, can repeat" long:"consumes"`
	Produces    []string `default:"application/json"               description:"add a content type to the global produces definitions, can repeat" long:"produces"`
	Schemes     []string `default:"http"                           description:"add a scheme to the global schemes definition, can repeat"         long:"scheme"`
	Contact     struct {
		Name  string `description:"name of the primary contact for the API"  long:"contact.name"`
		URL   string `description:"url of the primary contact for the API"   long:"contact.url"`
		Email string `description:"email of the primary contact for the API" long:"contact.email"`
	}
	License struct {
		Name string `description:"name of the license for the API" long:"license.name"`
		URL  string `description:"url of the license for the API"  long:"license.url"`
	}
}

// Execute this command.
func (s *Spec) Execute(args []string) error { _ = "STUB: not implemented"; return nil }
