package optioner

// Some[T] creates an Option[T] with the given value.
func Some[T any](v T) Optioner[T] {
	var o Optioner[T]
	o.v = &v
	return o
}

// None[T] creates an Option[T] with no value.
func None[T any]() Optioner[T] {
	var o Optioner[T]
	return o
}

// Option[T] creates a Option[T] that may or may not have a value.
// If *T is nil the return Option[T] is empty (equivalent to None[T]),
// If *T is non-nil return Option[T] will contain value (equivalent to Some[T]).
func Option[T any](v *T) Optioner[T] {
	return Optioner[T]{
		v: v,
	}
}

// Of[T] creates a Option[T] that may or may not have a value.
// If *T is nil the return Option[T] is empty (equivalent to None[T]),
// If *T is non-nil return Option[T] will contain value (equivalent to Some[T]).
//
// This is an alias for Option[T]
func Of[T any](v *T) Optioner[T] {
	return Option(v)
}
