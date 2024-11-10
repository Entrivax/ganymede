// Code generated by ent, DO NOT EDIT.

package multistreaminfo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/zibbp/ganymede/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldLTE(FieldID, id))
}

// DelayMs applies equality check predicate on the "delay_ms" field. It's identical to DelayMsEQ.
func DelayMs(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldEQ(FieldDelayMs, v))
}

// DelayMsEQ applies the EQ predicate on the "delay_ms" field.
func DelayMsEQ(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldEQ(FieldDelayMs, v))
}

// DelayMsNEQ applies the NEQ predicate on the "delay_ms" field.
func DelayMsNEQ(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldNEQ(FieldDelayMs, v))
}

// DelayMsIn applies the In predicate on the "delay_ms" field.
func DelayMsIn(vs ...int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldIn(FieldDelayMs, vs...))
}

// DelayMsNotIn applies the NotIn predicate on the "delay_ms" field.
func DelayMsNotIn(vs ...int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldNotIn(FieldDelayMs, vs...))
}

// DelayMsGT applies the GT predicate on the "delay_ms" field.
func DelayMsGT(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldGT(FieldDelayMs, v))
}

// DelayMsGTE applies the GTE predicate on the "delay_ms" field.
func DelayMsGTE(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldGTE(FieldDelayMs, v))
}

// DelayMsLT applies the LT predicate on the "delay_ms" field.
func DelayMsLT(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldLT(FieldDelayMs, v))
}

// DelayMsLTE applies the LTE predicate on the "delay_ms" field.
func DelayMsLTE(v int) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldLTE(FieldDelayMs, v))
}

// DelayMsIsNil applies the IsNil predicate on the "delay_ms" field.
func DelayMsIsNil() predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldIsNull(FieldDelayMs))
}

// DelayMsNotNil applies the NotNil predicate on the "delay_ms" field.
func DelayMsNotNil() predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.FieldNotNull(FieldDelayMs))
}

// HasVod applies the HasEdge predicate on the "vod" edge.
func HasVod() predicate.MultistreamInfo {
	return predicate.MultistreamInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VodTable, VodColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVodWith applies the HasEdge predicate on the "vod" edge with a given conditions (other predicates).
func HasVodWith(preds ...predicate.Vod) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(func(s *sql.Selector) {
		step := newVodStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlaylist applies the HasEdge predicate on the "playlist" edge.
func HasPlaylist() predicate.MultistreamInfo {
	return predicate.MultistreamInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaylistTable, PlaylistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistWith applies the HasEdge predicate on the "playlist" edge with a given conditions (other predicates).
func HasPlaylistWith(preds ...predicate.Playlist) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(func(s *sql.Selector) {
		step := newPlaylistStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MultistreamInfo) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MultistreamInfo) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MultistreamInfo) predicate.MultistreamInfo {
	return predicate.MultistreamInfo(sql.NotPredicates(p))
}
