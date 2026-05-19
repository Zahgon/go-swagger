// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/go-openapi/spec"
)

type mapStack struct {
	Type     *spec.Schema
	Next     *mapStack
	Previous *mapStack
	ValueRef *schemaGenContext
	Context  *schemaGenContext
	NewObj   *schemaGenContext
}

func newMapStack(context *schemaGenContext) (first, last *mapStack, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// reached the end of the rabbit hole

// found an anonymous object: create the struct from a newly created definition

// other cases where to stop are: a $ref or a simple object

// continue digging for maps

// return top and bottom entries of this stack of AdditionalProperties

// Build rewinds the stack of additional properties, building schemas from bottom to top.
//
//nolint:gocognit,gocyclo,cyclop // TODO(fredbi): refactor
func (mt *mapStack) Build() error { _ = "STUB: not implemented"; return nil }

// when the schema is an array or an alias, this may result in inconsistent
// nullable status between the map element and the array element (resp. the aliased type).
//
// Example: when an object has no property and only additionalProperties,
// which turn out to be arrays of some other object.

// save the initial override

// if we have an override at the top of stack, propagates it down nested arrays

// do it for nested arrays: override is also about map[string][][]... constructs

// cover other cases than arrays (aliased types)

// lift validations

// - we stopped on a ref, or anything else that require we call its Validate() method
// - if the alias / ref is on an interface (or stream) type: no validation

// a new model has been created during the stack construction (new ref on anonymous object)

// newly created model from anonymous object is declared as extra schema

// propagates extra schemas

// this is the genSchema for this new anonymous AdditionalProperty

// if there is a ValueRef, we must have a NewObj (from newMapStack() construction)

// we have a parent schema: build a schema for current AdditionalProperties

// we previously made a child schema: lifts things from that one
// - Required is not lifted (in a cascade of maps, only the last element is actually checked for Required)

// lift validations

// - we stopped on a ref, or anything else that require we call its Validate()
// - if the alias / ref is on an interface (or stream) type: no validation

// propagate overrides up the resolved schemas, but leaves any ExtraSchema untouched

func (mt *mapStack) HasMore() bool { _ = "STUB: not implemented"; return false }

/* currently unused:
func (mt *mapStack) Dict() map[string]any {
	res := make(map[string]any)
	res["context"] = mt.Context.Schema
	if mt.Next != nil {
		res["next"] = mt.Next.Dict()
	}
	if mt.NewObj != nil {
		res["obj"] = mt.NewObj.Schema
	}
	if mt.ValueRef != nil {
		res["value"] = mt.ValueRef.Schema
	}
	return res
}
*/
