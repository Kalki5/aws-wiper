package types

type Region string

var regions = [4]Region{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

func IsValidRegion(region Region) bool {
	for _, r := range regions {
		if region == r {
			return true
		}
	}
	return false
}

func (r *Region) ToString() string {
	return string(*r)
}
