package meta

type PageInfo struct {
	Title       string `meta:"og:title" json:"title"`
	Type        string `meta:"og:type" json:"type"`
	Url         string `meta:"og:url" json:"url"`
	Site        string `meta:"og:site" json:"site"`
	SiteName    string `meta:"og:site_name" json:"site_name"`
	Description string `meta:"og:description" json:"description"`
	Locale      string `meta:"og:locale" json:"locale"`
	// TODO: Implement images, videos and audios
	// Images      []OGImage
	// Videos      []OGVideo
	// Audios      []OGAudio
	Twitter TwitterCard `json:"twitter"`
}

type OGImage struct {
	Url       string `meta:"og:image,og:image:url" json:"url"`
	SecureUrl string `meta:"og:image:secure_url" json:"secure_url"`
	Width     int    `meta:"og:image:width" json:"width"`
	Height    int    `meta:"og:image:height" json:"height"`
	Type      string `meta:"og:image:type" json:"type"`
}

type OGVideo struct {
	Url       string `meta:"og:video,og:video:url" json:"url"`
	SecureUrl string `meta:"og:video:secure_url" json:"secure_url"`
	Width     int    `meta:"og:video:width" json:"width"`
	Height    int    `meta:"og:video:height" json:"height"`
	Type      string `meta:"og:video:type" json:"type"`
}

type OGAudio struct {
	Url       string `meta:"og:audio,og:audio:url" json:"url"`
	SecureUrl string `meta:"og:audio:secure_url" json:"secure_url"`
	Type      string `meta:"og:audio:type" json:"type"`
}

type TwitterCard struct {
	Card        string `meta:"twitter:card" json:"card"`
	Site        string `meta:"twitter:site" json:"site"`
	SiteId      string `meta:"twitter:site:id" json:"site_id"`
	Creator     string `meta:"twitter:creator" json:"creator"`
	CreatorId   string `meta:"twitter:creator:id" json:"creator_id"`
	Description string `meta:"twitter:description" json:"description"`
	Title       string `meta:"twitter:title" json:"title"`
	Image       string `meta:"twitter:image,twitter:image:src" json:"image"`
	ImageAlt    string `meta:"twitter:image:alt" json:"image_alt"`
	Url         string `meta:"twitter:url" json:"url"`
	Player      struct {
		Url    string `meta:"twitter:player" json:"url"`
		Width  int    `meta:"twitter:width" json:"width"`
		Height int    `meta:"twitter:height" json:"height"`
		Stream string `meta:"twitter:stream" json:"stream"`
	}
	IPhone struct {
		Name string `meta:"twitter:app:name:iphone" json:"name"`
		Id   string `meta:"twitter:app:id:iphone" json:"id"`
		Url  string `meta:"twitter:app:url:iphone" json:"url"`
	}
	IPad struct {
		Name string `meta:"twitter:app:name:ipad" json:"name"`
		Id   string `meta:"twitter:app:id:ipad" json:"id"`
		Url  string `meta:"twitter:app:url:ipad" json:"url"`
	}
	GooglePlay struct {
		Name string `meta:"twitter:app:name:googleplay" json:"name"`
		Id   string `meta:"twitter:app:id:googleplay" json:"id"`
		Url  string `meta:"twitter:app:url:googleplay" json:"url"`
	}
}
