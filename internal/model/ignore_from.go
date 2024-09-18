package model

type IgnoreFrom []int64

func (i IgnoreFrom) Contains(userID int64) bool {
	for _, v := range i {
		if v == userID {
			return true
		}
	}

	return false
}
