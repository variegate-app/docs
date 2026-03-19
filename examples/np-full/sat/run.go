package main

import (
	"fmt"
)

// SAT (satisfiability) decision solver for a CNF formula using exact backtracking:
// DPLL with unit propagation.
//
// - CNF is represented as a slice of clauses
// - each clause is a slice of literals
// - literals are ints: 1..n for x1..xn, and -1..-n for ¬x1..¬xn
//
// Returns a satisfying assignment if it exists; otherwise (nil, false).
func SolveSAT(nVars int, cnf [][]int) ([]bool, bool) {
	// -1 = unassigned, 0 = false, 1 = true
	asn := make([]int8, nVars)
	for i := range asn {
		asn[i] = -1
	}

	if ok := dpll(nVars, cnf, asn); !ok {
		return nil, false
	}

	model := make([]bool, nVars)
	for i := 0; i < nVars; i++ {
		model[i] = asn[i] == 1
	}
	return model, true
}

func dpll(nVars int, cnf [][]int, asn []int8) bool {
	// Propagate all implied assignments from unit clauses.
	if !propagateUnitClauses(cnf, asn) {
		return false
	}

	// If everything is assigned, the formula is SAT (propagation would
	// have found a conflict otherwise).
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

	// Pick the first unassigned variable and branch.
	var varIdx int
	for i := 0; i < nVars; i++ {
		if asn[i] == -1 {
			varIdx = i
			break
		}
	}

	// Try var = true first.
	{
		asn2 := append([]int8(nil), asn...)
		asn2[varIdx] = 1
		if dpll(nVars, cnf, asn2) {
			// Copy the found assignment back into the caller slice.
			copy(asn, asn2)
			return true
		}
	}

	// Try var = false.
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
			// Track whether the clause is satisfied,
			// and whether it contains exactly one unassigned literal.
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

			// Clause not satisfied:
			if unassignedCount == 0 {
				// All literals are false => conflict.
				return false
			}

			// Unit clause: only one possible literal can make it true.
			if unassignedCount == 1 {
				v := abs(unitLit) - 1
				must := int8(0)
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
	// Real-life example:
	// Features (логистика и контроль документов):
	// A: "Выбран способ доставки курьер"
	// B: "Выбран способ доставки пункт выдачи"
	// C: "Включена обязательная проверка документов"
	//
	// Constraints:
	// 1) Для заказа выбран хотя бы один способ доставки:
	//      (A ∨ B)
	// 2) Если выбран курьер (A), то обязателен C:
	//      (¬A ∨ C)
	// 3) Если выбран пункт выдачи (B), то обязателен C:
	//      (¬B ∨ C)
	// 4) Но C выключен по условию реальной ситуации:
	//      (¬C)
	//
	// Вопрос SAT: существуют ли значения A,B,C такие что все 4 ограничения выполняются?

	// Variables: 1->A, 2->B, 3->C
	nVars := 3

	// CNF:
	// (A ∨ B)
	// (¬A ∨ C)
	// (¬B ∨ C)
	// (¬C)
	cnf := [][]int{
		{1, 2},
		{-1, 3},
		{-2, 3},
		{-3},
	}

	model, ok := SolveSAT(nVars, cnf)
	if !ok {
		fmt.Println("SAT:", false)
		fmt.Println("Причина (интуитивно): если выбран любой способ доставки (A или B), то C обязан быть включен, но по условию C выключен.")
		return
	}

	fmt.Println("SAT:", true)
	fmt.Printf("A=%v, B=%v, C=%v\n", model[0], model[1], model[2])
}

