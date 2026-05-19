// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

const (
	iface   = "any"
	array   = "array"
	file    = "file"
	number  = "number"
	integer = "integer"
	boolean = "boolean"
	str     = "string"
	object  = "object"
	binary  = "binary"
	body    = "body"
	b64     = "byte"
)

// Extensions supported by go-swagger.
const (
	xClass        = "x-class"         // class name used by discriminator
	xGoCustomTag  = "x-go-custom-tag" // additional tag for serializers on struct fields
	xGoName       = "x-go-name"       // name of the generated go variable
	xGoType       = "x-go-type"       // reuse existing type (do not generate)
	xIsNullable   = "x-isnullable"
	xNullable     = "x-nullable" // turns the schema into a pointer
	xOmitEmpty    = "x-omitempty"
	xSchemes      = "x-schemes" // additional schemes supported for operations (server generation)
	xOrder        = "x-order"   // sort order for properties (or any schema)
	xGoJSONString = "x-go-json-string"
	xGoEnumCI     = "x-go-enum-ci" // make string enumeration case-insensitive

	xGoOperationTag = "x-go-operation-tag" // additional tag to override generation in operation groups
)

// swaggerTypeName contains a mapping from go type to swagger type or format.
var swaggerTypeName map[string]string

func initTypes() { _ = "STUB: not implemented"; return }

type typeResolver struct {
	Doc           *loads.Document
	ModelsPackage string // package alias (e.g. "models")
	ModelsFullPkg string // fully qualified package (e.g. "github.com/example/models")
	ModelName     string
	KnownDefs     map[string]struct{}
	// unexported fields
	keepDefinitionsPkg string
	knownDefsKept      map[string]struct{}
	definitionPkg      string // pkg alias to fill in GenSchema.Pkg
}

func newTypeResolver(pkg, _ string, doc *loads.Document) *typeResolver {
	_ = "STUB: not implemented"
	return nil
}

// NewWithModelName clones a type resolver and specifies a new model name.
func (t *typeResolver) NewWithModelName(name string) *typeResolver {
	_ = "STUB: not implemented"
	return nil
}

// propagates kept definitions

//nolint:gocognit // TODO(fredbi): refactor
func (t *typeResolver) ResolveSchema(schema *spec.Schema, isAnonymous, isRequired bool) (result resolvedType, err error) {
	_ = "STUB: not implemented"
	return *new(resolvedType), nil
}

// use hint to qualify type

// mark anonymous external types only, not definitions

// use spec to qualify type

// enforce bubbling up decisions taken about being an external type
// mark this type as an embedded external definition if requested

// for non-embedded, mark anonymous external types only, not definitions

// mark anonymous external types only, not definitions

// embedded external: by default consider validation is skipped for the external type
//
// NOTE: at this moment the template generates a type assertion, so this setting does not really matter
// for embedded types.

// non-embedded external type: by default consider that validation is enabled (SkipExternalValidation: false)

// special case of swagger type "file", rendered as io.ReadCloser interface

// no explicit object type, but inferred from object validations:
// this makes the type a map[string]any instead of any

func (t typeResolver) resolveExternalType(ext spec.Extensions) (*externalTypeDefinition, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// NOTE:
// * basic deconfliction of the default alias
// * if no package is specified, defaults to models (as provided from CLI or defaut generation location for models)

// in this case, the external type is assumed to be present in the current package.
// For completion, whenever this type is used in anonymous types declared by operations,
// we assume this is the package where models are expected to be found.

// knownDefGoType returns go type, package and package alias for definition.
func (t typeResolver) knownDefGoType(def string, schema spec.Schema, clearFunc func(string) string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// external type definition trumps regular type resolution

// withKeepDefinitionsPackage instructs the type resolver to keep previously resolved package name for
// definitions known at the moment it is first called.
func (t *typeResolver) withKeepDefinitionsPackage(definitionsPackage string) *typeResolver {
	_ = "STUB: not implemented"
	return nil
}

// withDefinitionPackage sets the definition pkg that object/struct types to be generated
// in GenSchema.Pkg field.
// ModelsPackage field can not replace definitionPkg since ModelsPackage will be prepend to .GoType,
// while definitionPkg is just used to fill the .Pkg in GenSchema.
func (t *typeResolver) withDefinitionPackage(pkg string) *typeResolver {
	_ = "STUB: not implemented"
	return nil
}

func (t *typeResolver) resolveSchemaRef(schema *spec.Schema, isRequired bool) (returns bool, result resolvedType, err error) {
	_ = "STUB: not implemented"
	return false, *new(resolvedType), nil
}

// deal with validations for an aliased external type

// this has to be overridden for slices and maps

func (t *typeResolver) inferAliasing(result *resolvedType, _ *spec.Schema, isAnonymous bool, _ bool) {
	_ = "STUB: not implemented"
	return
}

func (t *typeResolver) resolveFormat(schema *spec.Schema, isAnonymous bool, isRequired bool) (returns bool, result resolvedType) {
	_ = "STUB: not implemented"
	return false,

		// defaults to string
		*new(resolvedType)
}

// special case of swagger format "binary", rendered as io.ReadCloser interface and is therefore not a primitive type
// TODO: should set IsCustomFormatter=false in this case.

// propagate extensions in resolvedType

// isNullable hints the generator as to render the type with a pointer or not.
//
// A schema is deemed nullable (i.e. rendered by a pointer) when:
// - a custom extension says it has to be so
// - it is an object with properties
// - it is a composed object (allOf)
//
// The interpretation of Required as a mean to make a type nullable is carried out elsewhere.
func (t *typeResolver) isNullable(schema *spec.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

// isNullableOverride determines a nullable flag forced by an extension.
func (t *typeResolver) isNullableOverride(schema *spec.Schema) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (t *typeResolver) firstType(schema *spec.Schema) string { _ = "STUB: not implemented"; return "" }

// JSON-Schema multiple types, e.g. {"type": [ "object", "array" ]} are not supported.
// TODO: should keep the first _supported_ type, e.g. skip null

func (t *typeResolver) resolveArray(schema *spec.Schema, isAnonymous, isRequired bool) (result resolvedType, err error) {
	_ = "STUB: not implemented"
	return *new(resolvedType), nil
}

// resolve anonymous items

// Override the general nullability rule from ResolveSchema() in array elements:
// - only complex items are nullable (when not discriminated, not forced by x-nullable)
// - arrays of allOf have non nullable elements when not forced by x-nullable

// drill into $ref to figure out whether we want the element type to nullable or not

// this differs from isNullable for elements with AllOf

func (t *typeResolver) goTypeName(nm string) string { _ = "STUB: not implemented"; return "" }

// if a definitions package has been defined, already resolved definitions are
// always resolved against their original package (e.g. "models"), and not the
// current package.
// This allows complex anonymous extra schemas to reuse known definitions generated in another package.

//nolint:gocognit // TODO(fredbi): refactor
func (t *typeResolver) resolveObject(schema *spec.Schema, isAnonymous bool) (result resolvedType, err error) {
	_ = "STUB: not implemented"
	return *new(resolvedType), nil
}

// prioritize x-nullable extensions

// if this schema has properties, build a map of property name to
// resolved type, this should also flag the object as anonymous,
// when a ref is found, the anonymous flag will be reset

// no return here, still need to check for additional properties

// account for additional properties

// external AdditionalProperties are a special case because we look ahead into schemas

// only complex map elements are nullable (when not forced by x-nullable)
// TODO: figure out if required to check when not discriminated like arrays?

// Resolving nullability conflicts for:
// - map[][]...[]{items}
// - map[]{aliased type}
//
// when IsMap is true and the type is a distinct definition,
// aliased type or anonymous construct generated independently.
//
// IsMapNullOverride is to be handled by the generator for special cases
// where the map element is considered non nullable and the element itself is.
//
// This allows to appreciate nullability according to the context

// resolve the last items after nested arrays

// mark an override when nullable status conflicts, i.e. when the original type is not already nullable

// this locks the generator on the local nullability status

// an object without property and without AdditionalProperties schema is rendered as any

// an object without properties but with MinProperties or MaxProperties is rendered as map[string]any

// nullableBool makes a boolean a pointer when we want to distinguish the zero value from no value set.
// This is the case when:
// - a x-nullable extension says so in the spec
// - it is **not** a read-only property
// - it is a required property
// - it has a default value.
func nullableBool(schema *spec.Schema, isRequired bool) bool {
	_ = "STUB: not implemented"
	return false
}

// nullableNumber makes a number a pointer when we want to distinguish the zero value from no value set.
// This is the case when:
// - a x-nullable extension says so in the spec
// - it is **not** a read-only property
// - it is a required property
// - boundaries defines the zero value as a valid value:
//   - there is a non-exclusive boundary set at the zero value of the type
//   - the [min,max] range crosses the zero value of the type
func nullableNumber(schema *spec.Schema, isRequired bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *typeResolver) shortCircuitResolveExternal(tpe, pkg, alias string, extType *externalTypeDefinition, schema *spec.Schema, isRequired bool) resolvedType {
	_ = "STUB: not implemented"
	// short circuit type resolution for external types
	return *new(resolvedType)
}

// by default consider that we have a type with validations. Use hint "interface" or "noValidation" to disable validations

// x-nullable directive rules them all

// other extensions

// nullableString makes a string nullable when we want to distinguish the zero value from no value set.
// This is the case when:
// - a x-nullable extension says so in the spec
// - it is **not** a read-only property
// - it is a required property
// - it has a MinLength property set to 0
// - it has a default other than "" (the zero for strings) and no MinLength or zero MinLength.
func nullableString(schema *spec.Schema, isRequired bool) bool {
	_ = "STUB: not implemented"
	return false
}

func nullableStrfmt(schema *spec.Schema, isRequired bool) bool {
	_ = "STUB: not implemented"
	return false
}

func nullableExtension(ext spec.Extensions) *bool { _ = "STUB: not implemented"; return nil }

func boolExtension(ext spec.Extensions, key string) *bool { _ = "STUB: not implemented"; return nil }

func hasEnumCI(ve spec.Extensions) bool { _ = "STUB: not implemented"; return false }

// All enumeration types are case-sensitive by default

func warnSkipValidation(types any) func(string, any) { _ = "STUB: not implemented"; return nil }

// guardValidations removes (with a warning) validations that don't fit with the schema type.
//
// Notice that the "enum" validation is allowed on any type but file.
func guardValidations(tpe string, schema interface {
	Validations() spec.SchemaValidations
	SetValidations(validations spec.SchemaValidations)
}, types ...string,
) {
	_ = "STUB: not implemented"
	return
}

// keep MinLength/MaxLength on file

// other cases:  mapped as any: no validations allowed but Enum

// guardFormatConflicts handles all conflicting properties
// (for schema model or simple schema) when a format is set.
//
// At this moment, validation guards already handle all known conflicts, but for the
// special case of binary (i.e. io.Reader).
func guardFormatConflicts(format string, schema interface {
	Validations() spec.SchemaValidations
	SetValidations(validations spec.SchemaValidations)
},
) {
	_ = "STUB: not implemented"
	return
}

// for this format, no additional validations are supported

// no validations supported on binary fields at this moment (io.Reader)

// more cases should be inserted here if they arise

// resolvedType is a swagger type that has been resolved and analyzed for usage
// in a template.
type resolvedType struct {
	IsAnonymous       bool
	IsArray           bool
	IsMap             bool
	IsInterface       bool
	IsPrimitive       bool
	IsCustomFormatter bool
	IsAliased         bool
	IsNullable        bool
	IsStream          bool
	IsEmptyOmitted    bool
	IsJSONString      bool
	IsEnumCI          bool
	IsBase64          bool
	IsExternal        bool

	// A tuple gets rendered as an anonymous struct with P{index} as property name
	IsTuple            bool
	HasAdditionalItems bool

	// A complex object gets rendered as a struct
	IsComplexObject bool

	// A polymorphic type
	IsBaseType       bool
	HasDiscriminator bool

	GoType        string
	Pkg           string
	PkgAlias      string
	AliasedType   string
	SwaggerType   string
	SwaggerFormat string
	Extensions    spec.Extensions

	// The type of the element in a slice or map
	ElemType *resolvedType

	// IsMapNullOverride indicates that a nullable object is used within an
	// aliased map. In this case, the reference is not rendered with a pointer
	IsMapNullOverride bool

	// IsSuperAlias indicates that the aliased type is really the same type,
	// e.g. in golang, this translates to: type A = B
	IsSuperAlias bool

	// IsEmbedded applies to externally defined types. When embedded, a type
	// is generated in models that embeds the external type, with the Validate
	// method.
	IsEmbedded bool

	SkipExternalValidation bool
}

func simpleResolvedType(tn, fmt string, items *spec.Items, v *spec.CommonValidations) (result resolvedType) {
	_ = "STUB: not implemented"
	return *new(resolvedType)
}

// special case of swagger type "file", rendered as io.ReadCloser interface

// special case of swagger format "binary", rendered as io.ReadCloser interface
// TODO(fredbi): should set IsCustomFormatter=false when binary

// special case of swagger format "byte", rendered as a strfmt.Base64 type: no validation

// Zero returns an initializer for the type.
func (rt resolvedType) Zero() string {
	_ = "STUB: not implemented"
	// if type is aliased, provide zero from the aliased type
	return ""
}

// zero function provided as native or by strfmt function

// map and slice initializer

// object initializer

// interface initializer

// ToString returns a string conversion for a type akin to a string.
func (rt resolvedType) ToString(value string) string { _ = "STUB: not implemented"; return "" }

func (rt *resolvedType) setExtensions(schema *spec.Schema, origType string) {
	_ = "STUB: not implemented"
	return
}

func (rt *resolvedType) setIsEmptyOmitted(schema *spec.Schema, tpe string) {
	_ = "STUB: not implemented"
	return
}

// array of primitives are by default not empty-omitted, but arrays of aliased type are

func (rt *resolvedType) setIsJSONString(schema *spec.Schema, _ string) {
	_ = "STUB: not implemented"
	return
}

func (rt *resolvedType) setKind(kind string) { _ = "STUB: not implemented"; return }

// x-go-type:
//
//	 type: mytype
//	 import:
//	   package:
//	   alias:
//	 hints:
//	   kind: map|object|array|interface|primitive|stream|tuple
//	   nullable: true|false
//	embedded: true
type externalTypeDefinition struct {
	Type   string
	Import struct {
		Package string
		Alias   string
	}
	Hints struct {
		Kind         string
		Nullable     *bool
		NoValidation *bool
	}
	Embedded bool
}

func hasExternalType(ext spec.Extensions) (*externalTypeDefinition, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
