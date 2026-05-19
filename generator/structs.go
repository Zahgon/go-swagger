// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
)

// GenCommon contains common properties needed across
// definitions, app and operations
// TargetImportPath may be used by templates to import other (possibly
// generated) packages in the generation path (e.g. relative to GOPATH).
// TargetImportPath is NOT used by standard templates.
type GenCommon struct {
	Copyright        string
	TargetImportPath string
	RootedErrorPath  bool // wants array and map types to have a path corresponding to their type in reported errors
}

// GenDefinition contains all the properties to generate a
// definition from a swagger spec.
type GenDefinition struct {
	GenCommon
	GenSchema

	Package        string
	CliPackage     string
	Imports        map[string]string
	DefaultImports map[string]string
	ExtraSchemas   GenSchemaList
	DependsOn      []string
	External       bool
}

// GenDefinitions represents a list of operations to generate
// this implements a sort by operation id.
type GenDefinitions []GenDefinition

func (g GenDefinitions) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenDefinitions) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (g GenDefinitions) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// GenSchemaList is a list of schemas for generation.
//
// It can be sorted by name to get a stable struct layout for
// version control and such.
type GenSchemaList []GenSchema

// GenSchema contains all the information needed to generate the code
// for a schema.
type GenSchema struct {
	resolvedType
	sharedValidations

	Example                    string
	OriginalName               string
	Name                       string
	Suffix                     string
	Path                       string
	ValueExpression            string
	IndexVar                   string
	KeyVar                     string
	Title                      string
	Description                string
	Location                   string
	ReceiverName               string
	Items                      *GenSchema
	AllowsAdditionalItems      bool
	HasAdditionalItems         bool
	AdditionalItems            *GenSchema
	Object                     *GenSchema
	XMLName                    string
	CustomTag                  string
	Properties                 GenSchemaList
	AllOf                      GenSchemaList
	HasAdditionalProperties    bool
	IsAdditionalProperties     bool
	AdditionalProperties       *GenSchema
	StrictAdditionalProperties bool
	ReadOnly                   bool
	IsVirtual                  bool
	IsBaseType                 bool
	HasBaseType                bool
	IsSubType                  bool
	IsExported                 bool
	IsElem                     bool // IsElem gives some context when the schema is part of an array or a map
	IsProperty                 bool // IsProperty gives some context when the schema is a property of an object
	DiscriminatorField         string
	DiscriminatorValue         string
	Discriminates              map[string]string
	Parents                    []string
	IncludeValidator           bool
	IncludeModel               bool
	Default                    any
	WantsMarshalBinary         bool // do we generate MarshalBinary interface?
	StructTags                 []string
	ExtraImports               map[string]string // non-standard imports detected when using external types
	ExternalDocs               *spec.ExternalDocumentation
	WantsRootedErrorPath       bool
}

// PrintTags takes care of rendering tags for a struct field.
func (g GenSchema) PrintTags() string { _ = "STUB: not implemented"; return "" }

// Add extra struct tags, only if the tag hasn't already been set, i.e. example.
// Extra struct tags have the same value has the `json` tag.

// dedupe

// only add example tag if it's contained in the struct tags
// json representation of the example object

// Assemble the tags in key value pairs with the value properly quoted.

// Join the key value pairs by a space.

// If the values contain a backtick, we cannot render the tag using backticks because Go does not support
// escaping backticks in raw string literals.

// We have to escape the tag again to put it in a literal with double quotes as the tag format uses double quotes.

// UnderlyingType tells the go type or the aliased go type.
func (g GenSchema) UnderlyingType() string { _ = "STUB: not implemented"; return "" }

// ToString returns a string conversion expression for the schema.
func (g GenSchema) ToString() string { _ = "STUB: not implemented"; return "" }

func (g GenSchema) renderMarshalTag() string { _ = "STUB: not implemented"; return "" }

func (g GenSchemaList) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenSchemaList) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenSchemaList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// If both properties have x-order defined, then the one with lower x-order is smaller

// If only the first property has x-order defined, then it is smaller

// If only the second property has x-order defined, then it is smaller

// If neither property has x-order defined, then the one with lower lexicographic name is smaller

type sharedValidations struct {
	spec.SchemaValidations

	HasValidations        bool
	HasContextValidations bool
	Required              bool
	HasSliceValidations   bool
	ItemsEnum             []any

	// NOTE: "patternProperties" and "dependencies" not supported by Swagger 2.0
}

// GenResponse represents a response object for code generation.
type GenResponse struct {
	Package       string
	ModelsPackage string
	ReceiverName  string
	Name          string
	Description   string

	IsSuccess bool

	Code               int
	Method             string
	Path               string
	Headers            GenHeaders
	Schema             *GenSchema
	AllowsForStreaming bool

	Imports        map[string]string
	DefaultImports map[string]string

	Extensions map[string]any

	StrictResponders bool
	OperationName    string
	Examples         GenResponseExamples
	ReturnErrors     bool
}

// GenResponseExamples is a sortable collection []GenResponseExample.
type GenResponseExamples []GenResponseExample

func (g GenResponseExamples) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenResponseExamples) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenResponseExamples) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// GenResponseExample captures an example provided for a response for some mime type.
type GenResponseExample struct {
	MediaType string
	Example   any
}

// GenHeader represents a header on a response for code generation.
type GenHeader struct {
	resolvedType
	sharedValidations

	Package      string
	ReceiverName string
	IndexVar     string

	ID              string
	Name            string
	Path            string
	ValueExpression string

	Title       string
	Description string
	Default     any
	HasDefault  bool

	CollectionFormat string

	Child  *GenItems
	Parent *GenItems

	Converter string
	Formatter string

	ZeroValue string
}

// ItemsDepth returns a string "items.items..." with as many items as the level of nesting of the array.
// For a header objects it always returns "".
func (h *GenHeader) ItemsDepth() string {
	_ = "STUB: not implemented"
	// NOTE: this is currently used by templates to generate explicit comments in nested structures
	return ""
}

// ToString returns a string conversion expression for the header.
func (h GenHeader) ToString() string { _ = "STUB: not implemented"; return "" }

// GenHeaders is a sorted collection of headers for codegen.
type GenHeaders []GenHeader

func (g GenHeaders) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenHeaders) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenHeaders) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// HasSomeDefaults returns true is at least one header has a default value set.
func (g GenHeaders) HasSomeDefaults() bool {
	_ = "STUB: not implemented"
	// NOTE: this is currently used by templates to avoid empty constructs
	return false
}

// GenParameter is used to represent
// a parameter or a header for code generation.
type GenParameter struct {
	resolvedType
	sharedValidations

	ID              string
	Name            string
	ModelsPackage   string
	Path            string
	ValueExpression string
	IndexVar        string
	KeyVar          string
	ReceiverName    string
	Location        string
	Title           string
	Description     string
	Converter       string
	Formatter       string

	Schema *GenSchema

	CollectionFormat string

	CustomTag string

	Child  *GenItems
	Parent *GenItems

	// Unused
	// BodyParam *GenParameter

	Default         any
	HasDefault      bool
	ZeroValue       string
	AllowEmptyValue bool

	// validation strategy for Body params, which may mix model and simple constructs.
	// Distinguish the following cases:
	// - HasSimpleBodyParams: body is an inline simple type
	// - HasModelBodyParams: body is a model objectd
	// - HasSimpleBodyItems: body is an inline array of simple type
	// - HasModelBodyItems: body is an array of model objects
	// - HasSimpleBodyMap: body is a map of simple objects (possibly arrays)
	// - HasModelBodyMap: body is a map of model objects
	HasSimpleBodyParams bool
	HasModelBodyParams  bool
	HasSimpleBodyItems  bool
	HasModelBodyItems   bool
	HasSimpleBodyMap    bool
	HasModelBodyMap     bool

	Extensions map[string]any
}

// IsQueryParam returns true when this parameter is a query param.
func (g *GenParameter) IsQueryParam() bool { _ = "STUB: not implemented"; return false }

// IsPathParam returns true when this parameter is a path param.
func (g *GenParameter) IsPathParam() bool { _ = "STUB: not implemented"; return false }

// IsFormParam returns true when this parameter is a form param.
func (g *GenParameter) IsFormParam() bool { _ = "STUB: not implemented"; return false }

// IsHeaderParam returns true when this parameter is a header param.
func (g *GenParameter) IsHeaderParam() bool { _ = "STUB: not implemented"; return false }

// IsBodyParam returns true when this parameter is a body param.
func (g *GenParameter) IsBodyParam() bool { _ = "STUB: not implemented"; return false }

// IsFileParam returns true when this parameter is a file param.
func (g *GenParameter) IsFileParam() bool { _ = "STUB: not implemented"; return false }

// ItemsDepth returns a string "items.items..." with as many items as the level of nesting of the array.
// For a parameter object, it always returns "".
func (g *GenParameter) ItemsDepth() string {
	_ = "STUB: not implemented"
	// NOTE: this is currently used by templates to generate explicit comments in nested structures
	return ""
}

// UnderlyingType tells the go type or the aliased go type.
func (g GenParameter) UnderlyingType() string {
	_ = "STUB: not implemented"

	// ToString returns a string conversion expression for the parameter.
	return ""
}

func (g GenParameter) ToString() string { _ = "STUB: not implemented"; return "" }

// GenParameters represents a sorted parameter collection.
type GenParameters []GenParameter

func (g GenParameters) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenParameters) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (g GenParameters) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// HasSomeDefaults returns true is at least one parameter has a default value set.
func (g GenParameters) HasSomeDefaults() bool {
	_ = "STUB: not implemented"
	// NOTE: this is currently used by templates to avoid empty constructs
	return false
}

// GenItems represents the collection items for a collection parameter.
type GenItems struct {
	sharedValidations
	resolvedType

	Name             string
	Path             string
	ValueExpression  string
	CollectionFormat string
	Child            *GenItems
	Parent           *GenItems
	Converter        string
	Formatter        string

	Location string
	IndexVar string
	KeyVar   string

	// instructs generator to skip the splitting and parsing from CollectionFormat
	SkipParse bool
	// instructs generator that some nested structure needs an higher level loop index
	NeedsIndex bool
}

// ItemsDepth returns a string "items.items..." with as many items as the level of nesting of the array.
func (g *GenItems) ItemsDepth() string {
	_ = "STUB: not implemented"
	// NOTE: this is currently used by templates to generate explicit comments in nested structures
	return ""
}

// UnderlyingType tells the go type or the aliased go type.
func (g GenItems) UnderlyingType() string {
	_ = "STUB: not implemented"

	// ToString returns a string conversion expression for the item.
	return ""
}

func (g GenItems) ToString() string { _ = "STUB: not implemented"; return "" }

// GenOperationGroup represents a named (tagged) group of operations.
type GenOperationGroup struct {
	GenCommon

	Name       string
	Operations GenOperations

	Summary        string
	Description    string
	Imports        map[string]string
	DefaultImports map[string]string
	RootPackage    string
	GenOpts        *GenOpts
	PackageAlias   string

	ClientOptions *GenClientOptions
}

// GenOperationGroups is a sorted collection of operation groups.
type GenOperationGroups []GenOperationGroup

func (g GenOperationGroups) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenOperationGroups) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenOperationGroups) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// GenStatusCodeResponses a container for status code responses.
type GenStatusCodeResponses []GenResponse

func (g GenStatusCodeResponses) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenStatusCodeResponses) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenStatusCodeResponses) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// MarshalJSON marshals these responses to json
//
// This is used by DumpData.
func (g GenStatusCodeResponses) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// order marshalled output

//nolint:musttag // OK: we leave json fields identical to go fields. Dumpdata is used for debug.

// UnmarshalJSON unmarshals this GenStatusCodeResponses from json.
func (g *GenStatusCodeResponses) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:musttag // OK: we leave json fields identical to go fields. Dumpdata is used for debug.

// GenOperation represents an operation for code generation.
type GenOperation struct {
	GenCommon

	Package      string
	ReceiverName string
	Name         string
	Summary      string
	Description  string
	Method       string
	Path         string
	BasePath     string
	Tags         []string
	UseTags      bool
	RootPackage  string

	Imports        map[string]string
	DefaultImports map[string]string
	ExtraSchemas   GenSchemaList
	PackageAlias   string

	Authorized           bool
	Security             []GenSecurityRequirements // resolved security requirements for the operation
	SecurityDefinitions  GenSecuritySchemes
	SecurityRequirements []analysis.SecurityRequirement // original security requirements as per the spec (for doc)
	Principal            string
	PrincipalIsNullable  bool

	SuccessResponse  *GenResponse
	SuccessResponses []GenResponse
	Responses        GenStatusCodeResponses
	DefaultResponse  *GenResponse

	Params               GenParameters
	QueryParams          GenParameters
	PathParams           GenParameters
	HeaderParams         GenParameters
	FormParams           GenParameters
	HasQueryParams       bool
	HasPathParams        bool
	HasHeaderParams      bool
	HasFormParams        bool
	HasFormValueParams   bool
	HasFileParams        bool
	HasBodyParams        bool
	HasStreamingResponse bool

	Schemes              []string
	ExtraSchemes         []string
	SchemeOverrides      []string // original scheme overrides for operation, as per spec (for doc)
	ExtraSchemeOverrides []string // original extra scheme overrides for operation, as per spec (for doc)
	ProducesMediaTypes   []string
	ConsumesMediaTypes   []string
	TimeoutName          string

	Extensions map[string]any

	StrictResponders bool
	ReturnErrors     bool
	ExternalDocs     *spec.ExternalDocumentation
	Produces         []string // original produces for operation (for doc)
	Consumes         []string // original consumes for operation (for doc)
}

// GenOperations represents a list of operations to generate
// this implements a sort by operation id.
type GenOperations []GenOperation

func (g GenOperations) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenOperations) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (g GenOperations) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// GenApp represents all the meta data needed to generate an application
// from a swagger spec.
type GenApp struct {
	GenCommon

	APIPackage                 string
	ServerPackageAlias         string
	ImplementationPackageAlias string
	APIPackageAlias            string
	Package                    string
	ReceiverName               string
	Name                       string
	Principal                  string
	PrincipalIsNullable        bool
	DefaultConsumes            string
	DefaultProduces            string
	Host                       string
	BasePath                   string
	Info                       *spec.Info
	ExternalDocs               *spec.ExternalDocumentation
	Tags                       []spec.Tag
	Imports                    map[string]string
	DefaultImports             map[string]string
	Schemes                    []string
	ExtraSchemes               []string
	Consumes                   GenSerGroups
	Produces                   GenSerGroups
	SecurityDefinitions        GenSecuritySchemes
	SecurityRequirements       []analysis.SecurityRequirement // original security requirements as per the spec (for doc)
	Models                     []GenDefinition
	Operations                 GenOperations
	OperationGroups            GenOperationGroups
	SwaggerJSON                string
	// Embedded specs: this is important for when the generated server adds routes.
	// NOTE: there is a distinct advantage to having this in runtime rather than generated code.
	// We are not ever going to generate the router.
	// If embedding spec is an issue (e.g. memory usage), this can be excluded with the --exclude-spec
	// generation option. Alternative methods to serve spec (e.g. from disk, ...) may be implemented by
	// adding a middleware to the generated API.
	FlatSwaggerJSON string
	ExcludeSpec     bool
	GenOpts         *GenOpts
}

// UseGoStructFlags returns true when no strategy is specified or it is set to "go-flags".
func (g *GenApp) UseGoStructFlags() bool { _ = "STUB: not implemented"; return false }

// UsePFlags returns true when the flag strategy is set to pflag.
func (g *GenApp) UsePFlags() bool { _ = "STUB: not implemented"; return false }

// UseFlags returns true when the flag strategy is set to flag.
func (g *GenApp) UseFlags() bool { _ = "STUB: not implemented"; return false }

// UseIntermediateMode for https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28default.29
func (g *GenApp) UseIntermediateMode() bool { _ = "STUB: not implemented"; return false }

// UseModernMode for https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
func (g *GenApp) UseModernMode() bool { _ = "STUB: not implemented"; return false }

// GenSerGroups sorted representation of serializer groups.
type GenSerGroups []GenSerGroup

func (g GenSerGroups) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenSerGroups) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenSerGroups) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// NumSerializers yields the total number of serializer entries in this group.
func (g GenSerGroups) NumSerializers() int { _ = "STUB: not implemented"; return 0 }

// GenSerGroup represents a group of serializers: this links a serializer to a list of
// prioritized media types (mime).
type GenSerGroup struct {
	GenSerializer

	// All media types for this serializer. The redundant representation allows for easier use in templates
	AllSerializers GenSerializers
}

// GenSerializers sorted representation of serializers.
type GenSerializers []GenSerializer

func (g GenSerializers) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenSerializers) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenSerializers) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// GenSerializer represents a single serializer for a particular media type.
type GenSerializer struct {
	AppName        string // Application name
	ReceiverName   string
	Name           string   // Name of the Producer/Consumer (e.g. json, yaml, txt, bin)
	MediaType      string   // mime
	Implementation string   // func implementing the Producer/Consumer
	Parameters     []string // parameters supported by this serializer
}

// GenSecurityScheme represents a security scheme for code generation.
type GenSecurityScheme struct {
	AppName             string
	ID                  string
	Name                string
	ReceiverName        string
	IsBasicAuth         bool
	IsAPIKeyAuth        bool
	IsOAuth2            bool
	Scopes              []string
	Source              string
	Principal           string
	PrincipalIsNullable bool

	// from spec.SecurityScheme
	Description      string
	Type             string
	In               string
	Flow             string
	AuthorizationURL string
	TokenURL         string
	Extensions       map[string]any
	ScopesDesc       []GenSecurityScope
}

// GenSecuritySchemes sorted representation of serializers.
type GenSecuritySchemes []GenSecurityScheme

func (g GenSecuritySchemes) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenSecuritySchemes) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenSecuritySchemes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// GenSecurityRequirement represents a security requirement for an operation.
type GenSecurityRequirement struct {
	Name   string
	Scopes []string
}

// GenSecurityScope represents a scope descriptor for an OAuth2 security scheme.
type GenSecurityScope struct {
	Name        string
	Description string
}

// GenSecurityRequirements represents a compounded security requirement specification.
// In a []GenSecurityRequirements complete requirements specification,
// outer elements are interpreted as optional requirements (OR), and
// inner elements are interpreted as jointly required (AND).
type GenSecurityRequirements []GenSecurityRequirement

func (g GenSecurityRequirements) Len() int           { _ = "STUB: not implemented"; return 0 }
func (g GenSecurityRequirements) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (g GenSecurityRequirements) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// GenClientOptions holds extra pieces of information
// to generate a client.
type GenClientOptions struct {
	ProducesMediaTypes []string // filled with all producers if any method as more than 1
	ConsumesMediaTypes []string // filled with all consumers if any method as more than 1
}
