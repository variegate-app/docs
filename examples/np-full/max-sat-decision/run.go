package main

import "fmt"

// SolveMaxSATDecision solves decision version of MAX-SAT:
// given nVars and a CNF (clauses contain arbitrary number of literals),
// determine whether there exists an assignment that satisfies at least k clauses.
//
// This is an exact backtracking solver (exponential time).
// It is "optimal" in the sense that it returns SAT iff such assignment exists.
func SolveMaxSATDecision(nVars int, cnf [][]int, k int) (bool, []bool) {
	asn := make([]int8, nVars) // -1 unassigned, 0 false, 1 true
	for i := range asn {
		asn[i] = -1
	}

	clauses := make([][]int, len(cnf))
	copy(clauses, cnf)

	var bestModel []bool
	ok := backtrack(nVars, clauses, asn, 0, 0, k, &bestModel)
	if !ok {
		return false, nil
	}

	// backtrack stores a satisfying model when it finds one.
	return true, bestModel
}

// backtrack tries assignments variable by variable.
// It tracks currently satisfied clauses count, and uses an upper bound:
// maximum additional satisfied clauses is limited by remaining clauses.
func backtrack(
	nVars int,
	cnf [][]int,
	asn []int8,
	varIdx int,
	satisfiedCount int,
	k int,
	bestModel *[]bool,
) bool {
	// If we already satisfy enough clauses -> found a witness (optimal decision).
	if satisfiedCount >= k {
		model := make([]bool, nVars)
		for i := 0; i < nVars; i++ {
			model[i] = asn[i] == 1
		}
		*bestModel = model
		return true
	}

	if varIdx == nVars {
		return false
	}

	// Compute a quick upper bound for remaining clauses.
	// We can cheaply bound by "at most totalClauses - satisfiedCount".
	// This is weak but safe.
	totalClauses := len(cnf)
	if satisfiedCount+(totalClauses-satisfiedCount) < k {
		return false
	}

	// Branch var=true then var=false.
	{
		asn[varIdx] = 1
		newSat := countSatisfiedClauses(cnf, asn)
		if backtrack(nVars, cnf, asn, varIdx+1, newSat, k, bestModel) {
			return true
		}
		asn[varIdx] = -1
	}
	{
		asn[varIdx] = 0
		newSat := countSatisfiedClauses(cnf, asn)
		if backtrack(nVars, cnf, asn, varIdx+1, newSat, k, bestModel) {
			return true
		}
		asn[varIdx] = -1
	}

	return false
}

func countSatisfiedClauses(cnf [][]int, asn []int8) int {
	sat := 0
	for _, clause := range cnf {
		if isClauseSatisfied(clause, asn) {
			sat++
		}
	}
	return sat
}

func isClauseSatisfied(clause []int, asn []int8) bool {
	for _, lit := range clause {
		v := abs(lit) - 1
		val := asn[v]
		if val == -1 {
			continue
		}
		litIsTrue := (lit > 0 && val == 1) || (lit < 0 && val == 0)
		if litIsTrue {
			return true
		}
	}
	// Unassigned literals cannot make a clause satisfied yet.
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Real-life example:
	// In cybersecurity policy we configure authentication/audit mechanisms.
	//
	// Business/technical constraints mean that we may not satisfy all rules at once,
	// so we ask whether it's possible to satisfy at least k of them.
	//
	// Variables:
	//  A: enable strong authentication (e.g., require cryptographic approval)
	//  B: require one-time codes (OTP)
	//  C: enable extended audit logs (forensics/audit retention)
	//
	// Rules in CNF (each clause is an OR of literals). We try to satisfy at least k rules.
	//
	// k = 3 of 4 clauses:
	// 1) A OR B OR C
	// 2) ¬A OR ¬B OR C        (if strong auth and OTP both enabled -> audit logs required)
	// 3) ¬A OR B OR ¬C       (if strong auth and audit logs enabled -> OTP required)
	// 4)  A OR ¬B OR ¬C     (if OTP and audit logs enabled -> strong auth required)
	//
	// For this example there exists an assignment satisfying at least k clauses.

	labels := []string{"A(strong auth)", "B(OTP)", "C(audit logs)"}
	nVars := 3
	cnf := [][]int{
		{1, 2, 3},    // A ∨ B ∨ C
		{-1, -2, 3},  // ¬A ∨ ¬B ∨ C
		{-1, 2, -3},  // ¬A ∨ B ∨ ¬C
		{1, -2, -3},  // A ∨ ¬B ∨ ¬C
	}

	k := 3

	ok, model := SolveMaxSATDecision(nVars, cnf, k)
	fmt.Println("MAX-SAT decision (k=", k, "):", ok)
	if !ok {
		return
	}
	for i := 0; i < nVars; i++ {
		fmt.Printf("%s = %v\n", labels[i], model[i])
	}
}

