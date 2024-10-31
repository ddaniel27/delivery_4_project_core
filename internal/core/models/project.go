package models

type ProjectDocument struct {
	ID          int64        `json:"id" bson:"id"`
	Name        string       `json:"name" bson:"name"`
	Owner       Owner        `json:"owner" bson:"owner"`
	Groups      []Group      `json:"groups" bson:"groups"`
	Nominations []Nomination `json:"nominations" bson:"nominations"`
	Metadata    Metadata     `json:"metadata" bson:"metadata"`
}

type Owner struct {
	ID          int64  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Institution string `json:"institution" bson:"institution"`
}

type Group struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Nomination struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

type Metadata struct {
	IsUploaded bool   `json:"isUploaded" bson:"isUploaded"`
	IsPublic   bool   `json:"isPublic" bson:"isPublic"`
	Filename   string `json:"filename" bson:"filename"`
	Filetype   string `json:"filetype" bson:"filetype"`
	Filesize   int64  `json:"filesize" bson:"filesize"`
	ObjectKey  string `json:"objectKey" bson:"objectKey"`
	Duration   int64  `json:"duration" bson:"duration"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at"`
}
