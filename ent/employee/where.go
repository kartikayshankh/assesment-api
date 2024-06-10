// Code generated by ent, DO NOT EDIT.

package employee

import (
	"api/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPosition, v))
}

// Salary applies equality check predicate on the "salary" field. It's identical to SalaryEQ.
func Salary(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldSalary, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldName, v))
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldPosition, v))
}

// SalaryEQ applies the EQ predicate on the "salary" field.
func SalaryEQ(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldSalary, v))
}

// SalaryNEQ applies the NEQ predicate on the "salary" field.
func SalaryNEQ(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldSalary, v))
}

// SalaryIn applies the In predicate on the "salary" field.
func SalaryIn(vs ...float64) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldSalary, vs...))
}

// SalaryNotIn applies the NotIn predicate on the "salary" field.
func SalaryNotIn(vs ...float64) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldSalary, vs...))
}

// SalaryGT applies the GT predicate on the "salary" field.
func SalaryGT(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldSalary, v))
}

// SalaryGTE applies the GTE predicate on the "salary" field.
func SalaryGTE(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldSalary, v))
}

// SalaryLT applies the LT predicate on the "salary" field.
func SalaryLT(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldSalary, v))
}

// SalaryLTE applies the LTE predicate on the "salary" field.
func SalaryLTE(v float64) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldSalary, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(sql.NotPredicates(p))
}