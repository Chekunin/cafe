package memory

import (
	"cafe/pkg/common"
	errs "cafe/pkg/db_manager/errors"
	"errors"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/go-pg/pg/v10"
	"reflect"
	"strings"
)

const modelTag = "json"

func handleSqlError(err error, t reflect.Type) error {
	var pgErr pg.Error
	if ok := errors.As(err, &pgErr); ok {
		if pgErr.IntegrityViolation() {
			err = wrapErr.NewWrapErr(getCorrectSqlError(pgErr, t), err)
		} else if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityNotFound, err)
		}
	}
	return err
}

func getCorrectSqlError(pgErr pg.Error, t reflect.Type) error {
	var err = errs.ErrIntegrityConstraintViolation(nil)
	if tempErrFunc, has := sqlCodeMapping[pgErr.Field('C')]; has {
		// достаём ошибку по коду ошибки в sql
		if q, has := getFieldWithTagValue(t, "pg", pgErr.Field('c')); has {
			fieldName, has2 := q.Tag.Lookup(modelTag)
			if !has2 {
				fieldName = q.Tag.Get("json")
			}
			err = tempErrFunc(fieldName)
		} else {
			err = tempErrFunc(nil)
		}
	}
	return err
}

func getFieldWithTagValue(t reflect.Type, tagName, tagValue string) (reflect.StructField, bool) {
	switch t.Kind() {
	case reflect.Array, reflect.Ptr, reflect.Slice:
		return getFieldWithTagValue(t.Elem(), tagName, tagValue)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if val, ok := f.Tag.Lookup(tagName); ok && strings.Split(val, ",")[0] == tagValue {
				return f, true
			}
		}
	}
	return reflect.StructField{}, false
}

var sqlCodeMapping = map[string](func(meta interface{}) common.Err){
	"23000": errs.ErrIntegrityConstraintViolation,
	"23001": errs.ErrRestrictViolation,
	"23502": errs.ErrNotNullViolation,
	"23503": errs.ErrForeignKeyViolation,
	"23505": errs.ErrUniqueViolation,
	"23514": errs.ErrCheckViolation,
	"23P01": errs.ErrExclusionViolation,
}
