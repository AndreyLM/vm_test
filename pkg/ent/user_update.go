// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/andreylm/vm_test/pkg/ent/predicate"
	"github.com/andreylm/vm_test/pkg/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	user_name           *string
	full_name           *string
	city                *string
	birth_date          *time.Time
	department          *string
	gender              *string
	experience_years    *int
	addexperience_years *int
	predicates          []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetUserName sets the user_name field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.user_name = &s
	return uu
}

// SetFullName sets the full_name field.
func (uu *UserUpdate) SetFullName(s string) *UserUpdate {
	uu.full_name = &s
	return uu
}

// SetCity sets the city field.
func (uu *UserUpdate) SetCity(s string) *UserUpdate {
	uu.city = &s
	return uu
}

// SetBirthDate sets the birth_date field.
func (uu *UserUpdate) SetBirthDate(t time.Time) *UserUpdate {
	uu.birth_date = &t
	return uu
}

// SetDepartment sets the department field.
func (uu *UserUpdate) SetDepartment(s string) *UserUpdate {
	uu.department = &s
	return uu
}

// SetGender sets the gender field.
func (uu *UserUpdate) SetGender(s string) *UserUpdate {
	uu.gender = &s
	return uu
}

// SetExperienceYears sets the experience_years field.
func (uu *UserUpdate) SetExperienceYears(i int) *UserUpdate {
	uu.experience_years = &i
	uu.addexperience_years = nil
	return uu
}

// AddExperienceYears adds i to experience_years.
func (uu *UserUpdate) AddExperienceYears(i int) *UserUpdate {
	if uu.addexperience_years == nil {
		uu.addexperience_years = &i
	} else {
		*uu.addexperience_years += i
	}
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.experience_years != nil {
		if err := user.ExperienceYearsValidator(*uu.experience_years); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"experience_years\": %v", err)
		}
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.user_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUserName,
		})
	}
	if value := uu.full_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldFullName,
		})
	}
	if value := uu.city; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldCity,
		})
	}
	if value := uu.birth_date; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldBirthDate,
		})
	}
	if value := uu.department; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldDepartment,
		})
	}
	if value := uu.gender; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldGender,
		})
	}
	if value := uu.experience_years; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldExperienceYears,
		})
	}
	if value := uu.addexperience_years; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldExperienceYears,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id                  int64
	user_name           *string
	full_name           *string
	city                *string
	birth_date          *time.Time
	department          *string
	gender              *string
	experience_years    *int
	addexperience_years *int
}

// SetUserName sets the user_name field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.user_name = &s
	return uuo
}

// SetFullName sets the full_name field.
func (uuo *UserUpdateOne) SetFullName(s string) *UserUpdateOne {
	uuo.full_name = &s
	return uuo
}

// SetCity sets the city field.
func (uuo *UserUpdateOne) SetCity(s string) *UserUpdateOne {
	uuo.city = &s
	return uuo
}

// SetBirthDate sets the birth_date field.
func (uuo *UserUpdateOne) SetBirthDate(t time.Time) *UserUpdateOne {
	uuo.birth_date = &t
	return uuo
}

// SetDepartment sets the department field.
func (uuo *UserUpdateOne) SetDepartment(s string) *UserUpdateOne {
	uuo.department = &s
	return uuo
}

// SetGender sets the gender field.
func (uuo *UserUpdateOne) SetGender(s string) *UserUpdateOne {
	uuo.gender = &s
	return uuo
}

// SetExperienceYears sets the experience_years field.
func (uuo *UserUpdateOne) SetExperienceYears(i int) *UserUpdateOne {
	uuo.experience_years = &i
	uuo.addexperience_years = nil
	return uuo
}

// AddExperienceYears adds i to experience_years.
func (uuo *UserUpdateOne) AddExperienceYears(i int) *UserUpdateOne {
	if uuo.addexperience_years == nil {
		uuo.addexperience_years = &i
	} else {
		*uuo.addexperience_years += i
	}
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.experience_years != nil {
		if err := user.ExperienceYearsValidator(*uuo.experience_years); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"experience_years\": %v", err)
		}
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.user_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUserName,
		})
	}
	if value := uuo.full_name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldFullName,
		})
	}
	if value := uuo.city; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldCity,
		})
	}
	if value := uuo.birth_date; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldBirthDate,
		})
	}
	if value := uuo.department; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldDepartment,
		})
	}
	if value := uuo.gender; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldGender,
		})
	}
	if value := uuo.experience_years; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldExperienceYears,
		})
	}
	if value := uuo.addexperience_years; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldExperienceYears,
		})
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
