// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPayments(t *testing.T) {
	t.Parallel()

	query := Payments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPaymentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
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

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Payments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PaymentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPaymentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PaymentExists(ctx, tx, o.PaymentID)
	if err != nil {
		t.Errorf("Unable to check if Payment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PaymentExists to return true, but got false.")
	}
}

func testPaymentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	paymentFound, err := FindPayment(ctx, tx, o.PaymentID)
	if err != nil {
		t.Error(err)
	}

	if paymentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPaymentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Payments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPaymentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Payments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPaymentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	paymentOne := &Payment{}
	paymentTwo := &Payment{}
	if err = randomize.Struct(seed, paymentOne, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}
	if err = randomize.Struct(seed, paymentTwo, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = paymentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = paymentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Payments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPaymentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	paymentOne := &Payment{}
	paymentTwo := &Payment{}
	if err = randomize.Struct(seed, paymentOne, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}
	if err = randomize.Struct(seed, paymentTwo, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = paymentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = paymentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func paymentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func paymentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Payment) error {
	*o = Payment{}
	return nil
}

func testPaymentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Payment{}
	o := &Payment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, paymentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Payment object: %s", err)
	}

	AddPaymentHook(boil.BeforeInsertHook, paymentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	paymentBeforeInsertHooks = []PaymentHook{}

	AddPaymentHook(boil.AfterInsertHook, paymentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	paymentAfterInsertHooks = []PaymentHook{}

	AddPaymentHook(boil.AfterSelectHook, paymentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	paymentAfterSelectHooks = []PaymentHook{}

	AddPaymentHook(boil.BeforeUpdateHook, paymentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	paymentBeforeUpdateHooks = []PaymentHook{}

	AddPaymentHook(boil.AfterUpdateHook, paymentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	paymentAfterUpdateHooks = []PaymentHook{}

	AddPaymentHook(boil.BeforeDeleteHook, paymentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	paymentBeforeDeleteHooks = []PaymentHook{}

	AddPaymentHook(boil.AfterDeleteHook, paymentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	paymentAfterDeleteHooks = []PaymentHook{}

	AddPaymentHook(boil.BeforeUpsertHook, paymentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	paymentBeforeUpsertHooks = []PaymentHook{}

	AddPaymentHook(boil.AfterUpsertHook, paymentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	paymentAfterUpsertHooks = []PaymentHook{}
}

func testPaymentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPaymentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(paymentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPaymentToOneGroundBookingUsingBooking(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Payment
	var foreign GroundBooking

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, groundBookingDBTypes, false, groundBookingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GroundBooking struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BookingID, foreign.BookingID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Booking().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.BookingID, foreign.BookingID) {
		t.Errorf("want: %v, got %v", foreign.BookingID, check.BookingID)
	}

	ranAfterSelectHook := false
	AddGroundBookingHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *GroundBooking) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := PaymentSlice{&local}
	if err = local.L.LoadBooking(ctx, tx, false, (*[]*Payment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Booking == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Booking = nil
	if err = local.L.LoadBooking(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Booking == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testPaymentToOneSetOpGroundBookingUsingBooking(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Payment
	var b, c GroundBooking

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentDBTypes, false, strmangle.SetComplement(paymentPrimaryKeyColumns, paymentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, groundBookingDBTypes, false, strmangle.SetComplement(groundBookingPrimaryKeyColumns, groundBookingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, groundBookingDBTypes, false, strmangle.SetComplement(groundBookingPrimaryKeyColumns, groundBookingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*GroundBooking{&b, &c} {
		err = a.SetBooking(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Booking != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BookingPayments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BookingID, x.BookingID) {
			t.Error("foreign key was wrong value", a.BookingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookingID))
		reflect.Indirect(reflect.ValueOf(&a.BookingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BookingID, x.BookingID) {
			t.Error("foreign key was wrong value", a.BookingID, x.BookingID)
		}
	}
}

func testPaymentToOneRemoveOpGroundBookingUsingBooking(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Payment
	var b GroundBooking

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, paymentDBTypes, false, strmangle.SetComplement(paymentPrimaryKeyColumns, paymentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, groundBookingDBTypes, false, strmangle.SetComplement(groundBookingPrimaryKeyColumns, groundBookingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBooking(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBooking(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Booking().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Booking != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BookingID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BookingPayments) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPaymentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
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

func testPaymentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PaymentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPaymentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Payments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	paymentDBTypes = map[string]string{`PaymentID`: `integer`, `BookingID`: `integer`, `Amount`: `numeric`, `PaymentDate`: `date`, `PaymentStatus`: `character varying`}
	_              = bytes.MinRead
)

func testPaymentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(paymentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(paymentAllColumns) == len(paymentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPaymentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(paymentAllColumns) == len(paymentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Payment{}
	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, paymentDBTypes, true, paymentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(paymentAllColumns, paymentPrimaryKeyColumns) {
		fields = paymentAllColumns
	} else {
		fields = strmangle.SetComplement(
			paymentAllColumns,
			paymentPrimaryKeyColumns,
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

	slice := PaymentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPaymentsUpsert(t *testing.T) {
	t.Parallel()

	if len(paymentAllColumns) == len(paymentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Payment{}
	if err = randomize.Struct(seed, &o, paymentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Payment: %s", err)
	}

	count, err := Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, paymentDBTypes, false, paymentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Payment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Payment: %s", err)
	}

	count, err = Payments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}