package meta

type PageInfo struct {
	Title       string `meta:"og:title"`
	Type        string `meta:"og:type"`
	Url         string `meta:"og:url"`
	Site        string `meta:"og:site"`
	SiteName    string `meta:"og:site_name"`
	Description string `meta:"og:description"`
	Locale      string `meta:"og:locale"`
	Images      []*OGImage
	Videos      []*OGVideo
	Audios      []*OGAudio
	Twitter     *TwitterCard
}

type OGImage struct {
	Url       string `meta:"og:image,og:image:url"`
	SecureUrl string `meta:"og:image:secure_url"`
	Width     int    `meta:"og:image:width"`
	Height    int    `meta:"og:image:height"`
	Type      string `meta:"og:image:type"`
}

type OGVideo struct {
	Url       string `meta:"og:video,og:video:url"`
	SecureUrl string `meta:"og:video:secure_url"`
	Width     int    `meta:"og:video:width"`
	Height    int    `meta:"og:video:height"`
	Type      string `meta:"og:video:type"`
}

type OGAudio struct {
	Url       string `meta:"og:audio,og:audio:url"`
	SecureUrl string `meta:"og:audio:secure_url"`
	Type      string `meta:"og:audio:type"`
}

type TwitterCard struct {
	Card        string `meta:"twitter:card"`
	Site        string `meta:"twitter:site"`
	SiteId      string `meta:"twitter:site:id"`
	Creator     string `meta:"twitter:creator"`
	CreatorId   string `meta:"twitter:creator:id"`
	Description string `meta:"twitter:description"`
	Title       string `meta:"twitter:title"`
	Image       string `meta:"twitter:image,twitter:image:src"`
	ImageAlt    string `meta:"twitter:image:alt"`
	Url         string `meta:"twitter:url"`
	Player      struct {
		Url    string `meta:"twitter:player"`
		Width  int    `meta:"twitter:width"`
		Height int    `meta:"twitter:height"`
		Stream string `meta:"twitter:stream"`
	}
	IPhone struct {
		Name string `meta:"twitter:app:name:iphone"`
		Id   string `meta:"twitter:app:id:iphone"`
		Url  string `meta:"twitter:app:url:iphone"`
	}
	IPad struct {
		Name string `meta:"twitter:app:name:ipad"`
		Id   string `meta:"twitter:app:id:ipad"`
		Url  string `meta:"twitter:app:url:ipad"`
	}
	GooglePlay struct {
		Name string `meta:"twitter:app:name:googleplay"`
		Id   string `meta:"twitter:app:id:googleplay"`
		Url  string `meta:"twitter:app:url:googleplay"`
	}
}
