package types

type Resource struct {
	Id string
}

func (r *Resource) ToMap() Map {
	return Map{
		"id": r.Id,
	}
}
