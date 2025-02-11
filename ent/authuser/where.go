// Code generated by ent, DO NOT EDIT.

package authuser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Company applies equality check predicate on the "company" field. It's identical to CompanyEQ.
func Company(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompany), v))
	})
}

// Occupation applies equality check predicate on the "occupation" field. It's identical to OccupationEQ.
func Occupation(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupation), v))
	})
}

// PrivateKeyPath applies equality check predicate on the "private_key_path" field. It's identical to PrivateKeyPathEQ.
func PrivateKeyPath(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrivateKeyPath), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// CompanyEQ applies the EQ predicate on the "company" field.
func CompanyEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompany), v))
	})
}

// CompanyNEQ applies the NEQ predicate on the "company" field.
func CompanyNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompany), v))
	})
}

// CompanyIn applies the In predicate on the "company" field.
func CompanyIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompany), v...))
	})
}

// CompanyNotIn applies the NotIn predicate on the "company" field.
func CompanyNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompany), v...))
	})
}

// CompanyGT applies the GT predicate on the "company" field.
func CompanyGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompany), v))
	})
}

// CompanyGTE applies the GTE predicate on the "company" field.
func CompanyGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompany), v))
	})
}

// CompanyLT applies the LT predicate on the "company" field.
func CompanyLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompany), v))
	})
}

// CompanyLTE applies the LTE predicate on the "company" field.
func CompanyLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompany), v))
	})
}

// CompanyContains applies the Contains predicate on the "company" field.
func CompanyContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCompany), v))
	})
}

// CompanyHasPrefix applies the HasPrefix predicate on the "company" field.
func CompanyHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCompany), v))
	})
}

// CompanyHasSuffix applies the HasSuffix predicate on the "company" field.
func CompanyHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCompany), v))
	})
}

// CompanyEqualFold applies the EqualFold predicate on the "company" field.
func CompanyEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCompany), v))
	})
}

// CompanyContainsFold applies the ContainsFold predicate on the "company" field.
func CompanyContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCompany), v))
	})
}

// OccupationEQ applies the EQ predicate on the "occupation" field.
func OccupationEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupation), v))
	})
}

// OccupationNEQ applies the NEQ predicate on the "occupation" field.
func OccupationNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOccupation), v))
	})
}

// OccupationIn applies the In predicate on the "occupation" field.
func OccupationIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOccupation), v...))
	})
}

// OccupationNotIn applies the NotIn predicate on the "occupation" field.
func OccupationNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOccupation), v...))
	})
}

// OccupationGT applies the GT predicate on the "occupation" field.
func OccupationGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOccupation), v))
	})
}

// OccupationGTE applies the GTE predicate on the "occupation" field.
func OccupationGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOccupation), v))
	})
}

// OccupationLT applies the LT predicate on the "occupation" field.
func OccupationLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOccupation), v))
	})
}

// OccupationLTE applies the LTE predicate on the "occupation" field.
func OccupationLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOccupation), v))
	})
}

// OccupationContains applies the Contains predicate on the "occupation" field.
func OccupationContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOccupation), v))
	})
}

// OccupationHasPrefix applies the HasPrefix predicate on the "occupation" field.
func OccupationHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOccupation), v))
	})
}

// OccupationHasSuffix applies the HasSuffix predicate on the "occupation" field.
func OccupationHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOccupation), v))
	})
}

// OccupationEqualFold applies the EqualFold predicate on the "occupation" field.
func OccupationEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOccupation), v))
	})
}

// OccupationContainsFold applies the ContainsFold predicate on the "occupation" field.
func OccupationContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOccupation), v))
	})
}

// PrivateKeyPathEQ applies the EQ predicate on the "private_key_path" field.
func PrivateKeyPathEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathNEQ applies the NEQ predicate on the "private_key_path" field.
func PrivateKeyPathNEQ(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathIn applies the In predicate on the "private_key_path" field.
func PrivateKeyPathIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrivateKeyPath), v...))
	})
}

// PrivateKeyPathNotIn applies the NotIn predicate on the "private_key_path" field.
func PrivateKeyPathNotIn(vs ...string) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrivateKeyPath), v...))
	})
}

// PrivateKeyPathGT applies the GT predicate on the "private_key_path" field.
func PrivateKeyPathGT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathGTE applies the GTE predicate on the "private_key_path" field.
func PrivateKeyPathGTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathLT applies the LT predicate on the "private_key_path" field.
func PrivateKeyPathLT(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathLTE applies the LTE predicate on the "private_key_path" field.
func PrivateKeyPathLTE(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathContains applies the Contains predicate on the "private_key_path" field.
func PrivateKeyPathContains(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathHasPrefix applies the HasPrefix predicate on the "private_key_path" field.
func PrivateKeyPathHasPrefix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathHasSuffix applies the HasSuffix predicate on the "private_key_path" field.
func PrivateKeyPathHasSuffix(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathEqualFold applies the EqualFold predicate on the "private_key_path" field.
func PrivateKeyPathEqualFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrivateKeyPath), v))
	})
}

// PrivateKeyPathContainsFold applies the ContainsFold predicate on the "private_key_path" field.
func PrivateKeyPathContainsFold(v string) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrivateKeyPath), v))
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRole), v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRole), v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRole), v...))
	})
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v Provider) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v Provider) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvider), v))
	})
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...Provider) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProvider), v...))
	})
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...Provider) predicate.AuthUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProvider), v...))
	})
}

// HasAuthUserToToken applies the HasEdge predicate on the "AuthUserToToken" edge.
func HasAuthUserToToken() predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthUserToTokenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthUserToTokenTable, AuthUserToTokenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthUserToTokenWith applies the HasEdge predicate on the "AuthUserToToken" edge with a given conditions (other predicates).
func HasAuthUserToTokenWith(preds ...predicate.Token) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthUserToTokenInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthUserToTokenTable, AuthUserToTokenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthUserToServerTasks applies the HasEdge predicate on the "AuthUserToServerTasks" edge.
func HasAuthUserToServerTasks() predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthUserToServerTasksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AuthUserToServerTasksTable, AuthUserToServerTasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthUserToServerTasksWith applies the HasEdge predicate on the "AuthUserToServerTasks" edge with a given conditions (other predicates).
func HasAuthUserToServerTasksWith(preds ...predicate.ServerTask) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthUserToServerTasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AuthUserToServerTasksTable, AuthUserToServerTasksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AuthUser) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AuthUser) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AuthUser) predicate.AuthUser {
	return predicate.AuthUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
