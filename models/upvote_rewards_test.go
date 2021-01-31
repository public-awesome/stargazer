// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUpvoteRewards(t *testing.T) {
	t.Parallel()

	query := UpvoteRewards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUpvoteRewardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpvoteRewardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UpvoteRewards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpvoteRewardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UpvoteRewardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpvoteRewardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UpvoteRewardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UpvoteReward exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UpvoteRewardExists to return true, but got false.")
	}
}

func testUpvoteRewardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	upvoteRewardFound, err := FindUpvoteReward(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if upvoteRewardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUpvoteRewardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UpvoteRewards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUpvoteRewardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UpvoteRewards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUpvoteRewardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upvoteRewardOne := &UpvoteReward{}
	upvoteRewardTwo := &UpvoteReward{}
	if err = randomize.Struct(seed, upvoteRewardOne, upvoteRewardDBTypes, false, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}
	if err = randomize.Struct(seed, upvoteRewardTwo, upvoteRewardDBTypes, false, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = upvoteRewardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = upvoteRewardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UpvoteRewards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUpvoteRewardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upvoteRewardOne := &UpvoteReward{}
	upvoteRewardTwo := &UpvoteReward{}
	if err = randomize.Struct(seed, upvoteRewardOne, upvoteRewardDBTypes, false, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}
	if err = randomize.Struct(seed, upvoteRewardTwo, upvoteRewardDBTypes, false, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = upvoteRewardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = upvoteRewardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testUpvoteRewardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpvoteRewardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(upvoteRewardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpvoteRewardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUpvoteRewardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UpvoteRewardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUpvoteRewardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UpvoteRewards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	upvoteRewardDBTypes = map[string]string{`ID`: `integer`, `Height`: `bigint`, `VendorID`: `integer`, `PostID`: `text`, `RewardAddress`: `text`, `RewardAmount`: `bigint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                   = bytes.MinRead
)

func testUpvoteRewardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(upvoteRewardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(upvoteRewardAllColumns) == len(upvoteRewardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUpvoteRewardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(upvoteRewardAllColumns) == len(upvoteRewardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UpvoteReward{}
	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, upvoteRewardDBTypes, true, upvoteRewardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(upvoteRewardAllColumns, upvoteRewardPrimaryKeyColumns) {
		fields = upvoteRewardAllColumns
	} else {
		fields = strmangle.SetComplement(
			upvoteRewardAllColumns,
			upvoteRewardPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UpvoteRewardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUpvoteRewardsUpsert(t *testing.T) {
	t.Parallel()

	if len(upvoteRewardAllColumns) == len(upvoteRewardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UpvoteReward{}
	if err = randomize.Struct(seed, &o, upvoteRewardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UpvoteReward: %s", err)
	}

	count, err := UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, upvoteRewardDBTypes, false, upvoteRewardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpvoteReward struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UpvoteReward: %s", err)
	}

	count, err = UpvoteRewards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}