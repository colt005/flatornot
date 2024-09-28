package models

import (
	"database/sql/driver"
	"errors"
)

// VoteType is a custom type to represent the enum
type VoteType int

const (
	Flat VoteType = iota
	Round
)

// String method to convert VoteType to its string representation
func (v VoteType) String() string {
	switch v {
	case Flat:
		return "Flat"
	case Round:
		return "Round"
	default:
		return "Unknown"
	}
}

// ParseVoteType converts an integer to the corresponding VoteType
func ParseVoteType(value int) (VoteType, error) {
	switch value {
	case 0:
		return Flat, nil
	case 1:
		return Round, nil
	default:
		return -1, errors.New("invalid value for VoteType")
	}
}

// ParseStringToVoteType converts a string to the corresponding VoteType
func ParseStringToVoteType(value string) (VoteType, error) {
	switch value {
	case "Flat":
		return Flat, nil
	case "Round":
		return Round, nil
	default:
		return -1, errors.New("invalid string for VoteType")
	}
}

// ToInt converts the VoteType back to an integer representation
func (v VoteType) ToInt() int {
	return int(v)
}

func (v *VoteType) Scan(value interface{}) error {
	if value == nil {
		return errors.New("vote type is null")
	}

	switch t := value.(type) {
	case int64:
		switch t {
		case 0:
			*v = Flat
		case 1:
			*v = Round
		default:
			return errors.New("invalid value for VoteType")
		}
	default:
		return errors.New("invalid type for VoteType")
	}

	return nil
}

func (v VoteType) Value() (driver.Value, error) {
	switch v {
	case Flat, Round:
		return int(v), nil
	default:
		return nil, errors.New("invalid VoteType value")
	}
}
