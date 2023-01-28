// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sheason2019/spoved/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// ProjectName applies equality check predicate on the "project_name" field. It's identical to ProjectNameEQ.
func ProjectName(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProjectName, v))
}

// Describe applies equality check predicate on the "describe" field. It's identical to DescribeEQ.
func Describe(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescribe, v))
}

// GitURL applies equality check predicate on the "git_url" field. It's identical to GitURLEQ.
func GitURL(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldGitURL, v))
}

// DirPath applies equality check predicate on the "dir_path" field. It's identical to DirPathEQ.
func DirPath(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDirPath, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// ProjectNameEQ applies the EQ predicate on the "project_name" field.
func ProjectNameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProjectName, v))
}

// ProjectNameNEQ applies the NEQ predicate on the "project_name" field.
func ProjectNameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldProjectName, v))
}

// ProjectNameIn applies the In predicate on the "project_name" field.
func ProjectNameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldProjectName, vs...))
}

// ProjectNameNotIn applies the NotIn predicate on the "project_name" field.
func ProjectNameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldProjectName, vs...))
}

// ProjectNameGT applies the GT predicate on the "project_name" field.
func ProjectNameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldProjectName, v))
}

// ProjectNameGTE applies the GTE predicate on the "project_name" field.
func ProjectNameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldProjectName, v))
}

// ProjectNameLT applies the LT predicate on the "project_name" field.
func ProjectNameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldProjectName, v))
}

// ProjectNameLTE applies the LTE predicate on the "project_name" field.
func ProjectNameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldProjectName, v))
}

// ProjectNameContains applies the Contains predicate on the "project_name" field.
func ProjectNameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldProjectName, v))
}

// ProjectNameHasPrefix applies the HasPrefix predicate on the "project_name" field.
func ProjectNameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldProjectName, v))
}

// ProjectNameHasSuffix applies the HasSuffix predicate on the "project_name" field.
func ProjectNameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldProjectName, v))
}

// ProjectNameEqualFold applies the EqualFold predicate on the "project_name" field.
func ProjectNameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldProjectName, v))
}

// ProjectNameContainsFold applies the ContainsFold predicate on the "project_name" field.
func ProjectNameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldProjectName, v))
}

// DescribeEQ applies the EQ predicate on the "describe" field.
func DescribeEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescribe, v))
}

// DescribeNEQ applies the NEQ predicate on the "describe" field.
func DescribeNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDescribe, v))
}

// DescribeIn applies the In predicate on the "describe" field.
func DescribeIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDescribe, vs...))
}

// DescribeNotIn applies the NotIn predicate on the "describe" field.
func DescribeNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDescribe, vs...))
}

// DescribeGT applies the GT predicate on the "describe" field.
func DescribeGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDescribe, v))
}

// DescribeGTE applies the GTE predicate on the "describe" field.
func DescribeGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDescribe, v))
}

// DescribeLT applies the LT predicate on the "describe" field.
func DescribeLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDescribe, v))
}

// DescribeLTE applies the LTE predicate on the "describe" field.
func DescribeLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDescribe, v))
}

// DescribeContains applies the Contains predicate on the "describe" field.
func DescribeContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDescribe, v))
}

// DescribeHasPrefix applies the HasPrefix predicate on the "describe" field.
func DescribeHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDescribe, v))
}

// DescribeHasSuffix applies the HasSuffix predicate on the "describe" field.
func DescribeHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDescribe, v))
}

// DescribeEqualFold applies the EqualFold predicate on the "describe" field.
func DescribeEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDescribe, v))
}

// DescribeContainsFold applies the ContainsFold predicate on the "describe" field.
func DescribeContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDescribe, v))
}

// GitURLEQ applies the EQ predicate on the "git_url" field.
func GitURLEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldGitURL, v))
}

// GitURLNEQ applies the NEQ predicate on the "git_url" field.
func GitURLNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldGitURL, v))
}

// GitURLIn applies the In predicate on the "git_url" field.
func GitURLIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldGitURL, vs...))
}

// GitURLNotIn applies the NotIn predicate on the "git_url" field.
func GitURLNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldGitURL, vs...))
}

// GitURLGT applies the GT predicate on the "git_url" field.
func GitURLGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldGitURL, v))
}

// GitURLGTE applies the GTE predicate on the "git_url" field.
func GitURLGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldGitURL, v))
}

// GitURLLT applies the LT predicate on the "git_url" field.
func GitURLLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldGitURL, v))
}

// GitURLLTE applies the LTE predicate on the "git_url" field.
func GitURLLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldGitURL, v))
}

// GitURLContains applies the Contains predicate on the "git_url" field.
func GitURLContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldGitURL, v))
}

// GitURLHasPrefix applies the HasPrefix predicate on the "git_url" field.
func GitURLHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldGitURL, v))
}

// GitURLHasSuffix applies the HasSuffix predicate on the "git_url" field.
func GitURLHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldGitURL, v))
}

// GitURLEqualFold applies the EqualFold predicate on the "git_url" field.
func GitURLEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldGitURL, v))
}

// GitURLContainsFold applies the ContainsFold predicate on the "git_url" field.
func GitURLContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldGitURL, v))
}

// DirPathEQ applies the EQ predicate on the "dir_path" field.
func DirPathEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDirPath, v))
}

// DirPathNEQ applies the NEQ predicate on the "dir_path" field.
func DirPathNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDirPath, v))
}

// DirPathIn applies the In predicate on the "dir_path" field.
func DirPathIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDirPath, vs...))
}

// DirPathNotIn applies the NotIn predicate on the "dir_path" field.
func DirPathNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDirPath, vs...))
}

// DirPathGT applies the GT predicate on the "dir_path" field.
func DirPathGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDirPath, v))
}

// DirPathGTE applies the GTE predicate on the "dir_path" field.
func DirPathGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDirPath, v))
}

// DirPathLT applies the LT predicate on the "dir_path" field.
func DirPathLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDirPath, v))
}

// DirPathLTE applies the LTE predicate on the "dir_path" field.
func DirPathLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDirPath, v))
}

// DirPathContains applies the Contains predicate on the "dir_path" field.
func DirPathContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDirPath, v))
}

// DirPathHasPrefix applies the HasPrefix predicate on the "dir_path" field.
func DirPathHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDirPath, v))
}

// DirPathHasSuffix applies the HasSuffix predicate on the "dir_path" field.
func DirPathHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDirPath, v))
}

// DirPathEqualFold applies the EqualFold predicate on the "dir_path" field.
func DirPathEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDirPath, v))
}

// DirPathContainsFold applies the ContainsFold predicate on the "dir_path" field.
func DirPathContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDirPath, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CreatorTable, CreatorPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CreatorTable, CreatorPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
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
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		p(s.Not())
	})
}
