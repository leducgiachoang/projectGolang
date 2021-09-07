package models

type File struct {
	ID       string  `json:"id"`
	Source   string  `json:"source"`
	Key      string  `json:"key"`
	FileName string  `json:"file_name"`
	FileMime *string `json:"file_mime"`
	FileSize *int    `json:"file_size"`
	URL      string  `json:"url"`
	Sku      *string `json:"sku"`
}

type FileInput struct {
	// id use in create clipart
	ID       *string `json:"id"`
	Source   *string `json:"source"`
	Key      string  `json:"key"`
	FileName string  `json:"fileName"`
	FileMime string  `json:"fileMime"`
	FileSize int     `json:"fileSize"`
}

type FileSigned struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

type NewFile struct {
	// id use in create clipart
	ID       *string `json:"id"`
	Source   string  `json:"source"`
	Key      string  `json:"key"`
	FileName string  `json:"file_name"`
	FileMime string  `json:"file_mime"`
	FileSize int     `json:"file_size"`
}