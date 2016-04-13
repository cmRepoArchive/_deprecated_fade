package fade

type NotebookMetadata struct {
	AuthorId     string `json:'authorId'`
	Type         string `json:'type'`
	Time         string `json:'time'`
	ModifiedTime string `json:'modifiedTime'`
}

type Notebook struct {
	Metadata NotebookMetadata `json:'metadata'`
	Notes    []Note           `json:'notes'`
}
