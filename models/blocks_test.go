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

func testBlocks(t *testing.T) {
	t.Parallel()

	query := Blocks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBlocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
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

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Blocks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BlockExists(ctx, tx, o.Height)
	if err != nil {
		t.Errorf("Unable to check if Block exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlockExists to return true, but got false.")
	}
}

func testBlocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	blockFound, err := FindBlock(ctx, tx, o.Height)
	if err != nil {
		t.Error(err)
	}

	if blockFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBlocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Blocks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBlocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Blocks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBlocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(blockColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlockToManyHeightTransactions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Block
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Height = a.Height
	c.Height = a.Height

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.HeightTransactions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Height == b.Height {
			bFound = true
		}
		if v.Height == c.Height {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlockSlice{&a}
	if err = a.L.LoadHeightTransactions(ctx, tx, false, (*[]*Block)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HeightTransactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.HeightTransactions = nil
	if err = a.L.LoadHeightTransactions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HeightTransactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBlockToManyAddOpHeightTransactions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Block
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Transaction{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddHeightTransactions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Height != first.Height {
			t.Error("foreign key was wrong value", a.Height, first.Height)
		}
		if a.Height != second.Height {
			t.Error("foreign key was wrong value", a.Height, second.Height)
		}

		if first.R.Block != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Block != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.HeightTransactions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.HeightTransactions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.HeightTransactions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBlockToOneValidatorUsingProposer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Block
	var foreign Validator

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProposerAddress = foreign.Address
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Proposer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Address != foreign.Address {
		t.Errorf("want: %v, got %v", foreign.Address, check.Address)
	}

	slice := BlockSlice{&local}
	if err = local.L.LoadProposer(ctx, tx, false, (*[]*Block)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Proposer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Proposer = nil
	if err = local.L.LoadProposer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Proposer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlockToOneSetOpValidatorUsingProposer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Block
	var b, c Validator

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Validator{&b, &c} {
		err = a.SetProposer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Proposer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ProposerAddressBlocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProposerAddress != x.Address {
			t.Error("foreign key was wrong value", a.ProposerAddress)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProposerAddress))
		reflect.Indirect(reflect.ValueOf(&a.ProposerAddress)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ProposerAddress != x.Address {
			t.Error("foreign key was wrong value", a.ProposerAddress, x.Address)
		}
	}
}

func testBlocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
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

func testBlocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBlocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blockDBTypes = map[string]string{`Height`: `bigint`, `Hash`: `text`, `NumTXS`: `integer`, `TotalGas`: `bigint`, `ProposerAddress`: `text`, `Signatures`: `integer`, `BlockTimestamp`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_            = bytes.MinRead
)

func testBlocksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBlocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blockAllColumns, blockPrimaryKeyColumns) {
		fields = blockAllColumns
	} else {
		fields = strmangle.SetComplement(
			blockAllColumns,
			blockPrimaryKeyColumns,
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

	slice := BlockSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBlocksUpsert(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Block{}
	if err = randomize.Struct(seed, &o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, blockDBTypes, false, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err = Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
