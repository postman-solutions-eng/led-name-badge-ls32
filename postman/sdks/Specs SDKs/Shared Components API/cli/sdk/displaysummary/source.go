package displaysummary

// Present when a Postman API key was supplied. Indicates the summary was built from the Postman API catalog.
type Source string

const (
	SourcePostmanApiCatalog Source = "postman-api-catalog"
)
