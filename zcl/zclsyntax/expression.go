package zclsyntax

import (
	"github.com/apparentlymart/go-cty/cty"
	"github.com/apparentlymart/go-zcl/zcl"
)

// Expression is the abstract type for nodes that behave as zcl expressions.
type Expression interface {
	Node

	// The zcl.Expression methods are duplicated here, rather than simply
	// embedded, because both Node and zcl.Expression have a Range method
	// and so they conflict.

	Value(ctx *zcl.EvalContext) (cty.Value, zcl.Diagnostics)
	Variables() []zcl.Traversal
	StartRange() zcl.Range
}

// Assert that Expression implements zcl.Expression
var assertExprImplExpr zcl.Expression = Expression(nil)

// LiteralValueExpr is an expression that just always returns a given value.
type LiteralValueExpr struct {
	Val      cty.Value
	SrcRange zcl.Range
}

func (e *LiteralValueExpr) walkChildNodes(w internalWalkFunc) {
	// Literal values have no child nodes
}

func (e *LiteralValueExpr) Value(ctx *zcl.EvalContext) (cty.Value, zcl.Diagnostics) {
	return e.Val, nil
}

func (e *LiteralValueExpr) Range() zcl.Range {
	return e.SrcRange
}

func (e *LiteralValueExpr) StartRange() zcl.Range {
	return e.SrcRange
}

// ScopeTraversalExpr is an Expression that retrieves a value from the scope
// using a traversal.
type ScopeTraversalExpr struct {
	Traversal zcl.Traversal
	SrcRange  zcl.Range
}

func (e *ScopeTraversalExpr) walkChildNodes(w internalWalkFunc) {
	// Scope traversals have no child nodes
}

func (e *ScopeTraversalExpr) Value(ctx *zcl.EvalContext) (cty.Value, zcl.Diagnostics) {
	panic("ScopeTraversalExpr.Value not yet implemented")
}

func (e *ScopeTraversalExpr) Range() zcl.Range {
	return e.SrcRange
}

func (e *ScopeTraversalExpr) StartRange() zcl.Range {
	return e.SrcRange
}

// FunctionCallExpr is an Expression that calls a function from the EvalContext
// and returns its result.
type FunctionCallExpr struct {
	Name string
	Args []Expression

	NameRange       zcl.Range
	OpenParenRange  zcl.Range
	CloseParenRange zcl.Range
}

func (e *FunctionCallExpr) walkChildNodes(w internalWalkFunc) {
	for i, arg := range e.Args {
		e.Args[i] = w(arg).(Expression)
	}
}

func (e *FunctionCallExpr) Value(ctx *zcl.EvalContext) (cty.Value, zcl.Diagnostics) {
	panic("FunctionCallExpr.Value not yet implemented")
}

func (e *FunctionCallExpr) Range() zcl.Range {
	return zcl.RangeBetween(e.NameRange, e.CloseParenRange)
}

func (e *FunctionCallExpr) StartRange() zcl.Range {
	return zcl.RangeBetween(e.NameRange, e.OpenParenRange)
}