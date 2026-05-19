// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"io"
	"io/fs"
	"regexp"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"

	templatesrepo "github.com/go-swagger/go-swagger/generator/internal/templates-repo"
)

const (
	// default generation targets structure.
	defaultModelsTarget         = "models"
	defaultServerTarget         = "restapi"
	defaultClientTarget         = "client"
	defaultOperationsTarget     = "operations"
	defaultClientName           = "rest"
	defaultServerName           = "swagger"
	defaultScheme               = "http"
	defaultImplementationTarget = "implementation"

	winOS                    = "windows"
	readAllFile  fs.FileMode = 0o644 & fs.ModePerm
	readAllDir   fs.FileMode = 0o755 & fs.ModePerm
	readableFile fs.FileMode = 0o600 & fs.ModePerm
	readableDir  fs.FileMode = 0o700 & fs.ModePerm

	sensibleDefaultMapAlloc = 50
)

func init() {
	// all initializations for the generator package
	debugOptions()
	initTemplateRepo()
	initTypes()
}

// DefaultSectionOpts for a given opts, this is used when no config file is passed
// and uses the embedded templates when no local override can be found.
func DefaultSectionOpts(gen *GenOpts) { _ = "STUB: not implemented"; return }

// For CLI with default formatter (goimports), we needed to postpone the generation of model-supporting source,
// in order for go imports to run properly in all cases.
// If we completely migrate own custom formatter, we don't need to postpone.

// include a commandline tool app

// Use auto configure template

// MarkdownOpts for rendering a spec as markdown.
func MarkdownOpts() *LanguageOpts { _ = "STUB: not implemented"; return nil }

// MarkdownSectionOpts for a given opts and output file.
func MarkdownSectionOpts(gen *GenOpts, output string) { _ = "STUB: not implemented"; return }

// TemplateOpts allows for codegen customization.
type TemplateOpts struct {
	Name       string `mapstructure:"name"`
	Source     string `mapstructure:"source"`
	Target     string `mapstructure:"target"`
	FileName   string `mapstructure:"file_name"`
	SkipExists bool   `mapstructure:"skip_exists"`
	SkipFormat bool   `mapstructure:"skip_format"` // not a feature, but for debugging. generated code before formatting might not work because of unused imports.
}

// SectionOpts allows for specifying options to customize the templates used for generation.
type SectionOpts struct {
	Application     []TemplateOpts `mapstructure:"application"`
	Operations      []TemplateOpts `mapstructure:"operations"`
	OperationGroups []TemplateOpts `mapstructure:"operation_groups"`
	Models          []TemplateOpts `mapstructure:"models"`
	PostModels      []TemplateOpts `mapstructure:"post_models"`
}

// GenOptsCommon the options for the generator.
type GenOptsCommon struct {
	IncludeModel               bool
	IncludeValidator           bool
	IncludeHandler             bool
	IncludeParameters          bool
	IncludeResponses           bool
	IncludeURLBuilder          bool
	IncludeMain                bool
	IncludeSupport             bool
	IncludeCLi                 bool
	ExcludeSpec                bool
	DumpData                   bool
	ValidateSpec               bool
	FlattenOpts                *analysis.FlattenOpts
	IsClient                   bool
	defaultsEnsured            bool
	PropertiesSpecOrder        bool
	StrictAdditionalProperties bool
	AllowTemplateOverride      bool

	Spec                   string
	APIPackage             string
	ModelPackage           string
	ServerPackage          string
	ClientPackage          string
	CliPackage             string
	CliAppName             string // name of cli app. For example "dockerctl"
	ImplementationPackage  string
	Principal              string
	PrincipalCustomIface   bool   // user-provided interface for Principal (non-nullable)
	Target                 string // dir location where generated code is written to
	Sections               SectionOpts
	LanguageOpts           *LanguageOpts
	TypeMapping            map[string]string
	Imports                map[string]string
	DefaultScheme          string
	DefaultProduces        string
	DefaultConsumes        string
	WithXML                bool
	TemplateDir            string
	Template               string
	RegenerateConfigureAPI bool
	Operations             []string
	Models                 []string
	Tags                   []string
	StructTags             []string
	Name                   string
	FlagStrategy           string
	CompatibilityMode      string
	ExistingModels         string
	Copyright              string
	SkipTagPackages        bool
	MainPackage            string
	IgnoreOperations       bool
	AllowEnumCI            bool
	StrictResponders       bool
	AcceptDefinitionsOnly  bool
	WantsRootedErrorPath   bool
	ReturnErrors           bool
	WithCustomFormatter    bool

	templates *templatesrepo.Repository // a shallow clone of the global template repository
}

// CheckOpts carries out some global consistency checks on options.
func (g *GenOpts) CheckOpts() error { _ = "STUB: not implemented"; return nil }

// ensure spec path is absolute

// whenever opting for the custom formatter, we leave the basic formatting to the standard
// imports.Process and focus on a custom handling of imports.

// TargetPath returns the target generation path relative to the server package.
// This method is used by templates, e.g. with {{ .TargetPath }}
//
// Errors cases are prevented by calling CheckOpts beforehand.
//
// Example:
// Target: ${PWD}/tmp
// ServerPackage: abc/efg
//
// Server is generated in ${PWD}/tmp/abc/efg
// relative TargetPath returned: ../../../tmp.
func (g *GenOpts) TargetPath() string { _ = "STUB: not implemented"; return "" }

// That's for windows

// SpecPath returns the path to the spec relative to the server package.
// If the spec is remote keep this absolute location.
//
// If spec is not relative to server (e.g. lives on a different drive on windows),
// then the resolved path is absolute.
//
// This method is used by templates, e.g. with {{ .SpecPath }}
//
// Errors cases are prevented by calling CheckOpts beforehand.
func (g *GenOpts) SpecPath() string { _ = "STUB: not implemented"; return "" }

// Local specifications

// That's for windows

// PrincipalIsNullable indicates whether the principal type used for authentication
// may be used as a pointer.
func (g *GenOpts) PrincipalIsNullable() bool { _ = "STUB: not implemented"; return false }

// EnsureDefaults for these gen opts.
func (g *GenOpts) EnsureDefaults() error { _ = "STUB: not implemented"; return nil }

// set defaults for flattening options

// always include validator with models

func (g *GenOpts) location(t *TemplateOpts, data any) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (g *GenOpts) render(t *TemplateOpts, data any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try to load from repository (and enable dependencies)

// try to load template from disk, in TemplateDir if specified
// (dependencies resolution is limited to preloaded assets)

// Render template and write generated source code
// generated code is reformatted ("linted"), which gives an
// additional level of checking. If this step fails, the generated
// code is still dumped, for template debugging purposes.
func (g *GenOpts) write(t *TemplateOpts, data any) error { _ = "STUB: not implemented"; return nil }

// Directory settings consistent with file privileges.
// Environment's umask may alter this setup

// Conditionally format the code, unless the user wants to skip

// #nosec

// #nosec

func fileName(in string) string { _ = "STUB: not implemented"; return "" }

func (g *GenOpts) shouldRenderApp(t *TemplateOpts, _ *GenApp) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *GenOpts) shouldRenderOperations() bool { _ = "STUB: not implemented"; return false }

func (g *GenOpts) renderApplication(app *GenApp) error { _ = "STUB: not implemented"; return nil }

func (g *GenOpts) renderOperationGroup(gg *GenOperationGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GenOpts) renderOperation(gg *GenOperation) error { _ = "STUB: not implemented"; return nil }

func (g *GenOpts) renderDefinition(gg *GenDefinition) error { _ = "STUB: not implemented"; return nil }

func (g *GenOptsCommon) setTemplates() error { _ = "STUB: not implemented"; return nil }

// set contrib templates

// set custom templates

// defaultImports produces a default map for imports with models.
func (g *GenOpts) defaultImports() map[string]string { _ = "STUB: not implemented"; return nil }

// generated models

// external models

// resolve model representing an authenticated principal

// if principal is specified with the models generation package, do not import any extra package

// if principal is specified with a path, assume this is a fully qualified package and generate this import

// if principal is specified with a relative path (no "/", e.g. internal.Principal), assume it is located in generated target

// initImports produces a default map for import with the specified root for operations.
func (g *GenOpts) initImports(operationsPackage string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// PrincipalAlias returns an aliased type to the principal.
func (g *GenOpts) PrincipalAlias() string { _ = "STUB: not implemented"; return "" }

var ifaceRex = regexp.MustCompile(`^interface\s\{\s*\}$`)

func (g *GenOpts) resolvePrincipal() (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// handle possible conflicts with injected principal package
// NOTE(fred): we do not check here for conflicts with packages created from operation tags, only standard imports

func fileExists(target, name string) bool { _ = "STUB: not implemented"; return false }

func gatherModels(specDoc *loads.Document, modelNames []string) (map[string]spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// titleOrDefault infers a name for the app from the title of the spec.
func titleOrDefault(specDoc *loads.Document, name, defaultName string) string {
	_ = "STUB: not implemented"
	return ""
}

func mainNameOrDefault(specDoc *loads.Document, name, defaultName string) string {
	_ = "STUB: not implemented"
	// *_test won't do as main server name
	return ""
}

func appNameOrDefault(specDoc *loads.Document, name, defaultName string) string {
	_ = "STUB: not implemented"
	// *_test won't do as app names
	return ""
}

type opRef struct {
	Method string
	Path   string
	Key    string
	ID     string
	Op     *spec.Operation
}

type opRefs []opRef

func (o opRefs) Len() int           { _ = "STUB: not implemented"; return 0 }
func (o opRefs) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (o opRefs) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func gatherOperations(specDoc *analysis.Spec, operationIDs []string) map[string]opRef {
	_ = "STUB: not implemented"
	return nil
}

func pruneEmpty(in []string) (out []string) { _ = "STUB: not implemented"; return nil }

func trimBOM(in string) string { _ = "STUB: not implemented"; return "" }

const (
	securitySchemeAPIKey = "apikey"
	securitySchemeBasic  = "basic"
	securitySchemeOAuth2 = "oauth2"
)

// gatherSecuritySchemes produces a sorted representation from a map of spec security schemes.
func gatherSecuritySchemes(securitySchemes map[string]spec.SecurityScheme, appName, principal, receiver string, nullable bool) (security GenSecuritySchemes) {
	_ = "STUB: not implemented"
	return *new(GenSecuritySchemes)
}

// from original spec

// securityRequirements just clones the original SecurityRequirements from either the spec
// or an operation, without any modification. This is used to generate documentation.
func securityRequirements(orig []map[string][]string) (result []analysis.SecurityRequirement) {
	_ = "STUB: not implemented"
	return nil
}

// TODO(fred): sort this for stable generation

// gatherExtraSchemas produces a sorted list of extra schemas.
//
// ExtraSchemas are inlined types rendered in the same model file.
func gatherExtraSchemas(extraMap map[string]GenSchema) (extras GenSchemaList) {
	_ = "STUB: not implemented"
	return *new(GenSchemaList)
}

// figure out if top level validations are needed

func getExtraSchemes(ext spec.Extensions) []string { _ = "STUB: not implemented"; return nil }

func gatherURISchemes(swsp *spec.Swagger, operation spec.Operation) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dumpData(w io.Writer, data any) error { _ = "STUB: not implemented"; return nil }

func importAlias(pkg string) string { _ = "STUB: not implemented"; return "" }

// concatUnique concatenate collections of strings with deduplication.
func concatUnique(collections ...[]string) []string { _ = "STUB: not implemented"; return nil }
