package metadata

type PageInfo struct {
	Title       string `metadata:"og:title" json:"title"`
	Type        string `metadata:"og:type" json:"type"`
	URL         string `metadata:"og:url" json:"url"`
	Site        string `metadata:"og:site" json:"site"`
	SiteName    string `metadata:"og:site_name" json:"site_name"`
	Description string `metadata:"description,og:description" json:"description"`
	Locale      string `metadata:"og:locale" json:"locale"`
	// TODO: Implement images, videos and audios
	// Images  []OGImage
	// Videos  []OGVideo
	// Audios  []OGAudio
	Twitter TwitterCard `json:"twitter"`
}

type OGImage struct {
	Url       string `metadata:"og:image,og:image:url" json:"url"`
	SecureUrl string `metadata:"og:image:secure_url" json:"secure_url"`
	Width     int64  `metadata:"og:image:width" json:"width"`
	Height    int64  `metadata:"og:image:height" json:"height"`
	Type      string `metadata:"og:image:type" json:"type"`
}

type OGVideo struct {
	Url       string `metadata:"og:video,og:video:url" json:"url"`
	SecureUrl string `metadata:"og:video:secure_url" json:"secure_url"`
	Width     int64  `metadata:"og:video:width" json:"width"`
	Height    int64  `metadata:"og:video:height" json:"height"`
	Type      string `metadata:"og:video:type" json:"type"`
}

type OGAudio struct {
	Url       string `metadata:"og:audio,og:audio:url" json:"url"`
	SecureUrl string `metadata:"og:audio:secure_url" json:"secure_url"`
	Type      string `metadata:"og:audio:type" json:"type"`
}

type TwitterCard struct {
	Card        string `metadata:"twitter:card" json:"card"`
	Site        string `metadata:"twitter:site" json:"site"`
	SiteId      string `metadata:"twitter:site:id" json:"site_id"`
	Creator     string `metadata:"twitter:creator" json:"creator"`
	CreatorId   string `metadata:"twitter:creator:id" json:"creator_id"`
	Description string `metadata:"twitter:description" json:"description"`
	Title       string `metadata:"twitter:title" json:"title"`
	Image       string `metadata:"twitter:image,twitter:image:src" json:"image"`
	ImageAlt    string `metadata:"twitter:image:alt" json:"image_alt"`
	Url         string `metadata:"twitter:url" json:"url"`
	Player      struct {
		Url    string `metadata:"twitter:player" json:"url"`
		Width  int64  `metadata:"twitter:width" json:"width"`
		Height int64  `metadata:"twitter:height" json:"height"`
		Stream string `metadata:"twitter:stream" json:"stream"`
	}
	IPhone struct {
		Name string `metadata:"twitter:app:name:iphone" json:"name"`
		Id   string `metadata:"twitter:app:id:iphone" json:"id"`
		Url  string `metadata:"twitter:app:url:iphone" json:"url"`
	}
	IPad struct {
		Name string `metadata:"twitter:app:name:ipad" json:"name"`
		Id   string `metadata:"twitter:app:id:ipad" json:"id"`
		Url  string `metadata:"twitter:app:url:ipad" json:"url"`
	}
	GooglePlay struct {
		Name string `metadata:"twitter:app:name:googleplay" json:"name"`
		Id   string `metadata:"twitter:app:id:googleplay" json:"id"`
		Url  string `metadata:"twitter:app:url:googleplay" json:"url"`
	}
}
