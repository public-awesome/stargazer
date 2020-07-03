// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignatures)
	t.Run("Blocks", testBlocks)
	t.Run("Posts", testPosts)
	t.Run("Transactions", testTransactions)
	t.Run("Upvotes", testUpvotes)
	t.Run("Validators", testValidators)
}

func TestDelete(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesDelete)
	t.Run("Blocks", testBlocksDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("Transactions", testTransactionsDelete)
	t.Run("Upvotes", testUpvotesDelete)
	t.Run("Validators", testValidatorsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesQueryDeleteAll)
	t.Run("Blocks", testBlocksQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("Transactions", testTransactionsQueryDeleteAll)
	t.Run("Upvotes", testUpvotesQueryDeleteAll)
	t.Run("Validators", testValidatorsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesSliceDeleteAll)
	t.Run("Blocks", testBlocksSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("Transactions", testTransactionsSliceDeleteAll)
	t.Run("Upvotes", testUpvotesSliceDeleteAll)
	t.Run("Validators", testValidatorsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesExists)
	t.Run("Blocks", testBlocksExists)
	t.Run("Posts", testPostsExists)
	t.Run("Transactions", testTransactionsExists)
	t.Run("Upvotes", testUpvotesExists)
	t.Run("Validators", testValidatorsExists)
}

func TestFind(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesFind)
	t.Run("Blocks", testBlocksFind)
	t.Run("Posts", testPostsFind)
	t.Run("Transactions", testTransactionsFind)
	t.Run("Upvotes", testUpvotesFind)
	t.Run("Validators", testValidatorsFind)
}

func TestBind(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesBind)
	t.Run("Blocks", testBlocksBind)
	t.Run("Posts", testPostsBind)
	t.Run("Transactions", testTransactionsBind)
	t.Run("Upvotes", testUpvotesBind)
	t.Run("Validators", testValidatorsBind)
}

func TestOne(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesOne)
	t.Run("Blocks", testBlocksOne)
	t.Run("Posts", testPostsOne)
	t.Run("Transactions", testTransactionsOne)
	t.Run("Upvotes", testUpvotesOne)
	t.Run("Validators", testValidatorsOne)
}

func TestAll(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesAll)
	t.Run("Blocks", testBlocksAll)
	t.Run("Posts", testPostsAll)
	t.Run("Transactions", testTransactionsAll)
	t.Run("Upvotes", testUpvotesAll)
	t.Run("Validators", testValidatorsAll)
}

func TestCount(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesCount)
	t.Run("Blocks", testBlocksCount)
	t.Run("Posts", testPostsCount)
	t.Run("Transactions", testTransactionsCount)
	t.Run("Upvotes", testUpvotesCount)
	t.Run("Validators", testValidatorsCount)
}

func TestHooks(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesHooks)
	t.Run("Blocks", testBlocksHooks)
	t.Run("Posts", testPostsHooks)
	t.Run("Transactions", testTransactionsHooks)
	t.Run("Upvotes", testUpvotesHooks)
	t.Run("Validators", testValidatorsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesInsert)
	t.Run("BlockSignatures", testBlockSignaturesInsertWhitelist)
	t.Run("Blocks", testBlocksInsert)
	t.Run("Blocks", testBlocksInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("Transactions", testTransactionsInsert)
	t.Run("Transactions", testTransactionsInsertWhitelist)
	t.Run("Upvotes", testUpvotesInsert)
	t.Run("Upvotes", testUpvotesInsertWhitelist)
	t.Run("Validators", testValidatorsInsert)
	t.Run("Validators", testValidatorsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BlockSignatureToValidatorUsingValidator", testBlockSignatureToOneValidatorUsingValidator)
	t.Run("BlockToValidatorUsingProposer", testBlockToOneValidatorUsingProposer)
	t.Run("TransactionToBlockUsingBlock", testTransactionToOneBlockUsingBlock)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BlockToHeightTransactions", testBlockToManyHeightTransactions)
	t.Run("ValidatorToValidatorAddressBlockSignatures", testValidatorToManyValidatorAddressBlockSignatures)
	t.Run("ValidatorToProposerAddressBlocks", testValidatorToManyProposerAddressBlocks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BlockSignatureToValidatorUsingValidatorAddressBlockSignatures", testBlockSignatureToOneSetOpValidatorUsingValidator)
	t.Run("BlockToValidatorUsingProposerAddressBlocks", testBlockToOneSetOpValidatorUsingProposer)
	t.Run("TransactionToBlockUsingHeightTransactions", testTransactionToOneSetOpBlockUsingBlock)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BlockToHeightTransactions", testBlockToManyAddOpHeightTransactions)
	t.Run("ValidatorToValidatorAddressBlockSignatures", testValidatorToManyAddOpValidatorAddressBlockSignatures)
	t.Run("ValidatorToProposerAddressBlocks", testValidatorToManyAddOpProposerAddressBlocks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesReload)
	t.Run("Blocks", testBlocksReload)
	t.Run("Posts", testPostsReload)
	t.Run("Transactions", testTransactionsReload)
	t.Run("Upvotes", testUpvotesReload)
	t.Run("Validators", testValidatorsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesReloadAll)
	t.Run("Blocks", testBlocksReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("Transactions", testTransactionsReloadAll)
	t.Run("Upvotes", testUpvotesReloadAll)
	t.Run("Validators", testValidatorsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesSelect)
	t.Run("Blocks", testBlocksSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("Transactions", testTransactionsSelect)
	t.Run("Upvotes", testUpvotesSelect)
	t.Run("Validators", testValidatorsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesUpdate)
	t.Run("Blocks", testBlocksUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("Transactions", testTransactionsUpdate)
	t.Run("Upvotes", testUpvotesUpdate)
	t.Run("Validators", testValidatorsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BlockSignatures", testBlockSignaturesSliceUpdateAll)
	t.Run("Blocks", testBlocksSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("Transactions", testTransactionsSliceUpdateAll)
	t.Run("Upvotes", testUpvotesSliceUpdateAll)
	t.Run("Validators", testValidatorsSliceUpdateAll)
}