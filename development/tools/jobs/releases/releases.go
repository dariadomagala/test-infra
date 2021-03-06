package releases

// DO NOT EDIT. THIS FILE IS GENERATED

// List of currently supported releases
var (
	Release16 = mustParse("1.6")
	Release15 = mustParse("1.5")
	Release14 = mustParse("1.4")
	Release13 = mustParse("1.3")
)

// GetAllKymaReleaseBranches returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release15,
		Release14,
		Release13,
	}
}


