package optional

import (
	"database/sql"
	"errors"
)

// ErrNoneValueTaken represents the error that is raised when None value is taken.
var ErrNoneValueTaken = errors.New("none value taken")

type Option[T any] struct {
	value T
	valid bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{v, true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func ZeroAsNone[T comparable](v T) (ret Option[T]) {
	if v != ret.value {
		ret.valid = true
		ret.value = v
	}
	return ret
}

func NullString(v sql.NullString) (ret Option[string]) {
	ret.valid = v.Valid
	ret.value = v.String
	return ret
}

func ZeroAsNoneNullString(v sql.NullString) Option[string] {
	if !v.Valid {
		return None[string]()
	}

	return ZeroAsNone(v.String)
}

func Pointer[T comparable](v *T) (ret Option[T]) {
	if v != nil {
		ret.valid = true
		ret.value = *v
	}
	return ret
}

func (v Option[T]) IsSome() bool {
	return v.valid
}

func (v Option[T]) IsNone() bool {
	return !v.valid
}

func (v Option[T]) Unwrap() T {
	return v.value
}

func (v Option[T]) Take() (T, error) {
	if !v.valid {
		var defaultVal T
		return defaultVal, ErrNoneValueTaken
	}

	return v.value, nil
}

func (v Option[T]) UnwrapOr(fallback T) T {
	if v.valid {
		return v.value
	}
	return fallback
}

func (v Option[T]) UnwrapOrElse(f func() T) T {
	if v.valid {
		return v.value
	}
	return f()
}

func (v Option[T]) Or(optb Option[T]) Option[T] {
	if v.valid {
		return v
	}

	return optb
}

func Map[T, U any](option Option[T], mapper func(v T) U) Option[U] {
	if option.IsNone() {
		return None[U]()
	}

	return Some(mapper(option.value))
}

func MapOr[T, U any](option Option[T], fallbackValue U, mapper func(v T) U) U {
	if !option.valid {
		return fallbackValue
	}

	return mapper(option.value)
}

func AndThen[T, U any](option Option[T], then func(v T) Option[U]) Option[U] {
	if !option.valid {
		return None[U]()
	}

	return then(option.value)
}
