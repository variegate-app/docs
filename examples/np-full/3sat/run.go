package main

import "fmt"

// Solve3SAT solves a 3-CNF formula exactly via DPLL with unit propagation.
// CNF is represented as a list of clauses, each clause must contain exactly 3 literals.
// Literals are ints: 1..n for x1..xn, and -1..-n for ¬x1..¬xn.
//
// Returns a satisfying assignment if it exists; otherwise (nil, false).
func Solve3SAT(nVars int, cnf [][]int) ([]bool, bool) {
	asn := make([]int8, nVars) // -1 unassigned, 0 false, 1 true
	for i := range asn {
		asn[i] = -1
	}
	if !dpll(nVars, cnf, asn) {
		return nil, false
	}
	model := make([]bool, nVars)
	for i := 0; i < nVars; i++ {
		model[i] = asn[i] == 1
	}
	return model, true
}

func dpll(nVars int, cnf [][]int, asn []int8) bool {
	if !propagateUnitClauses(cnf, asn) {
		return false
	}

	allAssigned := true
	for i := 0; i < nVars; i++ {
		if asn[i] == -1 {
			allAssigned = false
			break
		}
	}
	if allAssigned {
		return true
	}

	var varIdx int
	for i := 0; i < nVars; i++ {
		if asn[i] == -1 {
			varIdx = i
			break
		}
	}

	// Branch: x=true.
	{
		asn2 := append([]int8(nil), asn...)
		asn2[varIdx] = 1
		if dpll(nVars, cnf, asn2) {
			copy(asn, asn2)
			return true
		}
	}
	// Branch: x=false.
	{
		asn2 := append([]int8(nil), asn...)
		asn2[varIdx] = 0
		if dpll(nVars, cnf, asn2) {
			copy(asn, asn2)
			return true
		}
	}

	return false
}

func propagateUnitClauses(cnf [][]int, asn []int8) bool {
	changed := true
	for changed {
		changed = false

		for _, clause := range cnf {
			satisfied := false
			unassignedCount := 0
			var unitLit int

			for _, lit := range clause {
				v := abs(lit) - 1
				val := asn[v]

				if val == -1 {
					unassignedCount++
					unitLit = lit
					continue
				}

				litIsTrue := (lit > 0 && val == 1) || (lit < 0 && val == 0)
				if litIsTrue {
					satisfied = true
					break
				}
			}

			if satisfied {
				continue
			}

			if unassignedCount == 0 {
				return false // all literals are false => conflict
			}

			// Unit clause: all but one literal are false; the remaining literal must be true.
			if unassignedCount == 1 {
				v := abs(unitLit) - 1
				var must int8
				if unitLit > 0 {
					must = 1
				}
				if asn[v] != -1 && asn[v] != must {
					return false
				}
				if asn[v] == -1 {
					asn[v] = must
					changed = true
				}
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Real-life example (fraud / security checks for a delivery):
	//
	// x1 (A): require signature
	// x2 (B): require one-time code
	// x3 (C): enable photo evidence
	//
	// Constraints (each clause has exactly 3 literals):
	// 1) At least one check must be enabled: (A ∨ B ∨ C)
	// 2) If A and B are enabled together, then C must be enabled: (¬A ∨ ¬B ∨ C)
	// 3) If A and C are enabled together, then B must be enabled: (¬A ∨ B ∨ ¬C)
	// 4) If B and C are enabled together, then A must be enabled: (A ∨ ¬B ∨ ¬C)

	labels := []string{"A(signature)", "B(OTP code)", "C(photo evidence)"}
	nVars := 3
	cnf := [][]int{
		{1, 2, 3},    // (A ∨ B ∨ C)
		{-1, -2, 3},  // (¬A ∨ ¬B ∨ C)
		{-1, 2, -3},  // (¬A ∨ B ∨ ¬C)
		{1, -2, -3},  // (A ∨ ¬B ∨ ¬C)
	}

	model, ok := Solve3SAT(nVars, cnf)
	if !ok {
		fmt.Println("3SAT: UNSAT")
		return
	}

	fmt.Println("3SAT: SAT")
	for i := 0; i < nVars; i++ {
		fmt.Printf("%s = %v\n", labels[i], model[i])
	}
}

