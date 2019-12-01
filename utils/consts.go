package utils

// MaxUint is the maximum value a Uint can store
const MaxUint = ^uint(0)

// MinUint is the minimum value a Uint can store
const MinUint = 0

// MaxInt is the maximum value an Int can store
const MaxInt = int(MaxUint >> 1)

// MinInt is the minimum value an Int can store
const MinInt = -MaxInt - 1
