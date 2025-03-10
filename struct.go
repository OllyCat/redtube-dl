package main

// JSONScript struct
type JSONScript struct {
	DisableSharebar   int    `json:"disable_sharebar"`
	HTMLPauseRoll     string `json:"htmlPauseRoll"`
	HTMLPostRoll      string `json:"htmlPostRoll"`
	EmbedCode         string `json:"embedCode"`
	Autoplay          int    `json:"autoplay"`
	Autoreplay        string `json:"autoreplay"`
	HidePostPauseRoll string `json:"hidePostPauseRoll"`
	VideoUnavailable  string `json:"video_unavailable"`
	PauserollURL      string `json:"pauseroll_url"`
	PostrollURL       string `json:"postroll_url"`
	VideoDuration     string `json:"video_duration"`
	ActionTags        string `json:"actionTags"`
	LinkURL           string `json:"link_url"`
	RelatedURL        string `json:"related_url"`
	ImageURL          string `json:"image_url"`
	VideoTitle        string `json:"video_title"`
	DefaultQuality    []int  `json:"defaultQuality"`
	VcServerURL       string `json:"vcServerUrl"`
	Quality480P       string `json:"quality_480p"`
	MediaDefinitions  []struct {
		DefaultQuality bool   `json:"defaultQuality"`
		Format         string `json:"format"`
		Quality        string `json:"quality"`
		VideoURL       string `json:"videoUrl"`
	} `json:"mediaDefinitions"`
	VideoUnavailableCountry string      `json:"video_unavailable_country"`
	TopratedURL             string      `json:"toprated_url"`
	MostviewedURL           string      `json:"mostviewed_url"`
	BrowserURL              interface{} `json:"browser_url"`
	MorefromthisuserURL     string      `json:"morefromthisuser_url"`
	Options                 string      `json:"options"`
	Cdn                     string      `json:"cdn"`
	StartLagThreshold       int         `json:"startLagThreshold"`
	OutBufferLagThreshold   int         `json:"outBufferLagThreshold"`
	AppID                   string      `json:"appId"`
	Service                 string      `json:"service"`
	CdnProvider             string      `json:"cdnProvider"`
	PrerollGlobalConfig     struct {
		ID                     int  `json:"id"`
		TimesToOmitBeforeFirst int  `json:"timesToOmitBeforeFirst"`
		ForgetUserAfter        int  `json:"forgetUserAfter"`
		Skippable              bool `json:"skippable"`
		SkipSeconds            int  `json:"skipSeconds"`
		ConfigByHostType       struct {
			Native struct {
				TimesToOmitBeforeFirst int  `json:"timesToOmitBeforeFirst"`
				ForgetUserAfter        int  `json:"forgetUserAfter"`
				Skippable              bool `json:"skippable"`
				SkipSeconds            int  `json:"skipSeconds"`
			} `json:"native"`
			Embed struct {
				TimesToOmitBeforeFirst int  `json:"timesToOmitBeforeFirst"`
				ForgetUserAfter        int  `json:"forgetUserAfter"`
				Skippable              bool `json:"skippable"`
				SkipSeconds            int  `json:"skipSeconds"`
			} `json:"embed"`
		} `json:"configByHostType"`
		Targeting []struct {
			PlatformsByHost struct {
				Native []string `json:"native"`
			} `json:"platformsByHost"`
		} `json:"targeting"`
	} `json:"prerollGlobalConfig"`
	Mp4Seek  string   `json:"mp4_seek"`
	Hotspots []string `json:"hotspots"`
	Thumbs   struct {
		SamplingFrequency int    `json:"samplingFrequency"`
		Type              string `json:"type"`
		CdnType           string `json:"cdnType"`
		URLPattern        string `json:"urlPattern"`
		ThumbHeight       string `json:"thumbHeight"`
		ThumbWidth        string `json:"thumbWidth"`
	} `json:"thumbs"`
	NextVideo struct {
		Thumb    string `json:"thumb"`
		Video    string `json:"video"`
		Duration string `json:"duration"`
		Title    string `json:"title"`
		IsHD     string `json:"isHD"`
		NextURL  string `json:"nextUrl"`
		Timeout  int    `json:"timeout"`
		Desktop  bool   `json:"desktop"`
		Mobile   bool   `json:"mobile"`
		Chanel   string `json:"chanel"`
	} `json:"nextVideo"`
	Language string `json:"language"`
}
