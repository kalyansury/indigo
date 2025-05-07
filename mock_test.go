package indigo_test

import (
	"fmt"
	"time"

	"github.com/ezachrisen/indigo"
)

// -------------------------------------------------- MOCK EVALUATOR
// mockEvaluator is used for testing
// It provides minimal evaluation of rules and captures
// information about which rules were processed, etc.
type mockEvaluator struct {
	rules []string // a list of rule IDs in the evaluator
	// if set, diagnostic information is only returned if the flag was
	// set during compilation
	diagnosticCompileRequired bool
	// Introduce an artificial delay in evaluating the expression.
	// Used for testing the engine's context cancelation functionality.
	evalDelay time.Duration
}

type program struct {
	compiledDiagnostics bool
}

func newMockEvaluator() *mockEvaluator {
	return &mockEvaluator{}
}

// Compile implements the Evaluator interface.
// It compiles the given expression against the provided schema to produce
// a program that can be evaluated.
//
// Parameters:
//   - expr: The expression to compile
//   - s: Schema to compile against
//   - resultType: Expected type of the result
//   - collectDiagnostics: When true, compilation diagnostics will be included
//   - dryRun: When true, performs compilation checks without creating a full program
//
// Returns:
//   - An executable program (as any)
//   - Error if compilation fails
func (m *mockEvaluator) Compile(_ string, _ indigo.Schema, _ indigo.Type, collectDiagnostics, _ bool) (any, error) {
	p := program{}
	if collectDiagnostics {
		p.compiledDiagnostics = true
	}

	return p, nil
}

// Evaluate executes an expression against the provided data and returns the result.
// It simulates evaluation behavior for testing purposes, with configurable delays and diagnostic behaviors.
// The mockEvaluator only knows how to evaluate 1 string: `true`.
// If the expression is this, the evaluation is true, otherwise false.
//
// Parameters:
//   - data: Map containing variables available during expression evaluation.
//   - expr: The expression string to evaluate.
//   - s: Schema definition used for type validation.
//   - self: Reference to the current object for self-referential expressions.
//   - prog: Optional pre-compiled program that may contain diagnostic information.
//   - resultType: Expected type of the evaluation result.
//   - returnDiagnostics: Whether to include diagnostic information in the result.
//
// Returns:
//   - any: The result of evaluation, with specific handling for "true" and "self" expressions.
//   - *indigo.Diagnostics: Diagnostics information if requested and conditions are met, otherwise nil.
//   - error: Error encountered during evaluation, if any.
func (m *mockEvaluator) Evaluate(_ map[string]any, expr string, _ indigo.Schema, self any, prog any, _ indigo.Type, returnDiagnostics bool) (any, *indigo.Diagnostics, error) {
	// m.rulesTested = append(m.rulesTested, r.ID)
	time.Sleep(m.evalDelay)
	prg := program{}

	p, ok := prog.(program)
	if m.diagnosticCompileRequired {
		if !ok {
			return false, nil, fmt.Errorf("compiled data type assertion failed")
		}
		prg = p
	}

	var diagnostics *indigo.Diagnostics
	if returnDiagnostics && ((m.diagnosticCompileRequired && prg.compiledDiagnostics) || !m.diagnosticCompileRequired) {
		diagnostics = &indigo.Diagnostics{}
	}

	if expr == `true` {
		// return indigo.Value{
		// 	Val:  true,
		// 	Type: indigo.Bool{},
		// }, diagnostics, nil
		return true, diagnostics, nil
	}

	if expr == `self` && self != nil {
		return self, diagnostics, nil
		// return indigo.Value{
		// 	Val:  self,
		// 	Type: indigo.Int{},
		// }, diagnostics, nil
	}

	// return indigo.Value{
	// 	Val:  false,
	// 	Type: indigo.Bool{},
	// }, diagnostics, nil
	return false, diagnostics, nil
}

func (m *mockEvaluator) Reset() {
	m.rules = []string{}
}

func (m *mockEvaluator) PrintInternalStructure() {
	for _, v := range m.rules {
		fmt.Println("Rule id", v)
	}
}
