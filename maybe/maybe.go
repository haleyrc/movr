package maybe

import "time"

type Float64 struct {
	Provided bool
	Value    float64
}

type String struct {
	Provided bool
	Value    string
}

type Time struct {
	Provided bool
	Value    time.Time
}
