package ce_ps

type CdcChange struct {
	After struct {
		Attributes      string `json:"attributes"`
		ID              string `json:"id"`
		IsOriginalTitle bool   `json:"is_original_title"`
		Language        string `json:"language"`
		Ordering        int    `json:"ordering"`
		Region          string `json:"region"`
		Title           string `json:"title"`
		TitleID         string `json:"title_id"`
		Types           string `json:"types"`
	} `json:"after"`
	Key     []string `json:"key"`
	Updated string   `json:"updated"`
}