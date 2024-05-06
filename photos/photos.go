package photos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OrlovDiga/flickr"
)

type PhotoInfo struct {
	Id           string `xml:"id,attr"`
	Secret       string `xml:"secret,attr"`
	Server       string `xml:"server,attr"`
	Farm         string `xml:"farm,attr"`
	DateUploaded string `xml:"dateuploaded,attr"`
	IsFavorite   bool   `xml:"isfavorite,attr"`
	License      string `xml:"license,attr"`
	// NOTE: one less than safety level set on upload (ie, here 0 = safe, 1 = moderate, 2 = restricted)
	//       while on upload, 1 = safe, 2 = moderate, 3 = restricted
	SafetyLevel    int    `xml:"safety_level,attr"`
	Rotation       int    `xml:"rotation,attr"`
	OriginalSecret string `xml:"originalsecret,attr"`
	OriginalFormat string `xml:"originalformat,attr"`
	Views          int    `xml:"views,attr"`
	Media          string `xml:"media,attr"`
	Title          string `xml:"title"`
	Description    string `xml:"description"`
	Visibility     struct {
		IsPublic bool `xml:"ispublic,attr"`
		IsFriend bool `xml:"isfriend,attr"`
		IsFamily bool `xml:"isfamily,attr"`
	} `xml:"visibility"`
	Dates struct {
		Posted           string `xml:"posted,attr"`
		Taken            string `xml:"taken,attr"`
		TakenGranularity string `xml:"takengranularity,attr"`
		TakenUnknown     string `xml:"takenunknown,attr"`
		LastUpdate       string `xml:"lastupdate,attr"`
	} `xml:"dates"`
	Permissions struct {
		PermComment string `xml:"permcomment,attr"`
		PermAdMeta  string `xml:"permadmeta,attr"`
	} `xml:"permissions"`
	Editability struct {
		CanComment string `xml:"cancomment,attr"`
		CanAddMeta string `xml:"canaddmeta,attr"`
	} `xml:"editability"`
	PublicEditability struct {
		CanComment string `xml:"cancomment,attr"`
		CanAddMeta string `xml:"canaddmeta,attr"`
	} `xml:"publiceditability"`
	Usage struct {
		CanDownload string `xml:"candownload,attr"`
		CanBlog     string `xml:"canblog,attr"`
		CanPrint    string `xml:"canprint,attr"`
		CanShare    string `xml:"canshare,attr"`
	} `xml:"usage"`
	Comments int   `xml:"comments"`
	Tags     []Tag `xml:"tags>tag"`
	// Notes XXX: not handled yet
	// People XXX: not handled yet
	// Urls XXX: not handled yet
}
type Tag struct {
	ID    string `xml:"id,attr"`
	Raw   string `xml:"raw,attr"`
	Value string `xml:",chardata"`
}

type PhotoInfoResponse struct {
	flickr.BasicResponse
	Photo PhotoInfo `xml:"photo"`
}
type PrivacyType int64

const (
	yes PrivacyType = 1
	no  PrivacyType = 0
)

type PhotoDownloadInfo struct {
	Label  string `xml:"label,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Source string `xml:"source,attr"`
	Url    string `xml:"url,attr"`
	Media  string `xml:"media,attr"`
}
type PhotoAccessInfo struct {
	flickr.BasicResponse
	Sizes []PhotoDownloadInfo `xml:"sizes>size"`
}

type PhotoSearchRequest struct {
	Text         string
	Sort         string
	License      []string
	ContentTypes []string
	Media        string
	Lat          float64
	Lon          float64
	Extras       []string
	Radius       float64
	Page         int // first page is 1
	PerPage      int // max 500
}

type PhotosSearchResponse struct {
	Photos Photos `xml:"photos"`
	flickr.BasicResponse
}

type Photos struct {
	Page    string            `xml:"page,attr"`
	Pages   string            `xml:"pages,attr"`
	Perpage string            `xml:"perpage,attr"`
	Total   string            `xml:"total,attr"`
	Photo   []PhotoSearchItem `xml:"photo"`
}

type PhotoSearchItem struct {
	ID                   string `xml:"id,attr"`
	Owner                string `xml:"owner,attr"`
	Secret               string `xml:"secret,attr"`
	Server               string `xml:"server,attr"`
	Farm                 string `xml:"farm,attr"`
	Title                string `xml:"title,attr"`
	IsPublic             string `xml:"ispublic,attr"`
	IsFriend             string `xml:"isfriend,attr"`
	IsFamily             string `xml:"isfamily,attr"`
	License              string `xml:"license,attr"`
	OWidth               string `xml:"o_width,attr"`
	OHeight              string `xml:"o_height,attr"`
	DateUpload           string `xml:"dateupload,attr"`
	LastUpdate           string `xml:"lastupdate,attr"`
	DateTaken            string `xml:"datetaken,attr"`
	DateTakenGranularity string `xml:"datetakengranularity,attr"`
	DateTakenUnknown     string `xml:"datetakenunknown,attr"`
	OwnerName            string `xml:"ownername,attr"`
	IconServer           string `xml:"iconserver,attr"`
	IconFarm             string `xml:"iconfarm,attr"`
	Views                string `xml:"views,attr"`
	Tags                 string `xml:"tags,attr"`
	MachineTags          string `xml:"machine_tags,attr"`
	OriginalSecret       string `xml:"originalsecret,attr"`
	OriginalFormat       string `xml:"originalformat,attr"`
	Latitude             string `xml:"latitude,attr"`
	Longitude            string `xml:"longitude,attr"`
	Accuracy             string `xml:"accuracy,attr"`
	Context              string `xml:"context,attr"`
	Media                string `xml:"media,attr"`
	MediaStatus          string `xml:"media_status,attr"`
	PathAlias            string `xml:"pathalias,attr"`
	URLSq                string `xml:"url_sq,attr"`
	HeightSq             string `xml:"height_sq,attr"`
	WidthSq              string `xml:"width_sq,attr"`
	URLT                 string `xml:"url_t,attr"`
	HeightT              string `xml:"height_t,attr"`
	WidthT               string `xml:"width_t,attr"`
	URLS                 string `xml:"url_s,attr"`
	HeightS              string `xml:"height_s,attr"`
	WidthS               string `xml:"width_s,attr"`
	URLQ                 string `xml:"url_q,attr"`
	HeightQ              string `xml:"height_q,attr"`
	WidthQ               string `xml:"width_q,attr"`
	URLM                 string `xml:"url_m,attr"`
	HeightM              string `xml:"height_m,attr"`
	WidthM               string `xml:"width_m,attr"`
	URLN                 string `xml:"url_n,attr"`
	HeightN              string `xml:"height_n,attr"`
	WidthN               string `xml:"width_n,attr"`
	URLZ                 string `xml:"url_z,attr"`
	HeightZ              string `xml:"height_z,attr"`
	WidthZ               string `xml:"width_z,attr"`
	URLC                 string `xml:"url_c,attr"`
	HeightC              string `xml:"height_c,attr"`
	WidthC               string `xml:"width_c,attr"`
	URLL                 string `xml:"url_l,attr"`
	HeightL              string `xml:"height_l,attr"`
	WidthL               string `xml:"width_l,attr"`
	URLO                 string `xml:"url_o,attr"`
	HeightO              string `xml:"height_o,attr"`
	WidthO               string `xml:"width_o,attr"`
	Description          string `xml:"description"`
}

// description,license,date_upload,date_taken,owner_name,icon_server,original_format,last_update,geo,tags,machine_tags,o_dims,views,media,path_alias,url_sq,url_t,url_s,url_q,url_m,url_n,url_z,url_c,url_l,url_o
// GetSizes get all the downloadable link as
func GetSizes(client *flickr.FlickrClient, photoId string) (*PhotoAccessInfo, error) {

	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"

	client.Args.Set("method", "flickr.photos.getSizes")
	client.Args.Set("photo_id", photoId)
	client.OAuthSign()
	response := &PhotoAccessInfo{}
	err := flickr.DoPost(client, response)
	return response, err

}

// Set permission of a photo from flickr
// this method requires authentica with 'write' permission
func SetPerms(client *flickr.FlickrClient, id string, isPublic PrivacyType, IsFriend PrivacyType, isFamily PrivacyType) (*flickr.BasicResponse, error) {

	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.setPerms")
	client.Args.Set("photo_id", id)
	client.Args.Set("is_public", strconv.Itoa(int(isPublic)))
	client.Args.Set("is_friend", strconv.Itoa(int(IsFriend)))
	client.Args.Set("is_family", strconv.Itoa(int(isFamily)))
	client.OAuthSign()
	response := &flickr.BasicResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// Delete a photo from Flickr
// This method requires authentication with 'delete' permission.
func Delete(client *flickr.FlickrClient, id string) (*flickr.BasicResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.delete")
	client.Args.Set("photo_id", id)
	client.OAuthSign()

	response := &flickr.BasicResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// Get information about a Flickr photo
func GetInfo(client *flickr.FlickrClient, id string, secret string) (*PhotoInfoResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.getInfo")
	client.Args.Set("photo_id", id)
	if secret != "" {
		client.Args.Set("secret", secret)
	}
	client.OAuthSign()

	response := &PhotoInfoResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// Set date posted and date taken on a Flickr photo
// datePosted and dateTaken are optional and may be set to ""
func SetDates(client *flickr.FlickrClient, id string, datePosted string, dateTaken string) (*flickr.BasicResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.setDates")
	client.Args.Set("photo_id", id)
	if datePosted != "" {
		client.Args.Set("date_posted", datePosted)
	}
	if dateTaken != "" {
		client.Args.Set("date_taken", dateTaken)
	}
	client.OAuthSign()

	response := &flickr.BasicResponse{}
	err := flickr.DoPost(client, response)
	return response, err
}

// AddTags add tags to an existing photo
func AddTags(client *flickr.FlickrClient, photoId string, tags []string) error {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.addTags")
	client.Args.Set("photo_id", photoId)
	client.Args.Set("tags", strings.Join(tags, ","))
	client.OAuthSign()
	response := &flickr.BasicResponse{}
	return flickr.DoPost(client, response)
}

func Search(client *flickr.FlickrClient, req PhotoSearchRequest) (*PhotosSearchResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.HTTPVerb = "POST"
	client.Args.Set("method", "flickr.photos.search")
	if req.Text != "" {
		client.Args.Set("text", req.Text)
	}
	if req.Sort != "" {
		client.Args.Set("sort", req.Sort)
	}
	if len(req.License) > 0 {
		client.Args.Set("license", strings.Join(req.License, ","))
	}
	if len(req.ContentTypes) > 0 {
		client.Args.Set("content_types", strings.Join(req.ContentTypes, ","))
	}
	if req.Media != "" {
		client.Args.Set("media", req.Media)
	}
	if req.Lat != 0 {
		client.Args.Set("lat", fmt.Sprintf("%v", req.Lat))
	}
	if req.Lon != 0 {
		client.Args.Set("lon", fmt.Sprintf("%v", req.Lon))
	}
	if len(req.Extras) > 0 {
		client.Args.Set("extras", strings.Join(req.Extras, ","))
	}
	if req.Radius != 0 {
		client.Args.Set("radius", fmt.Sprintf("%v", req.Radius))
	}
	if req.Page != 0 {
		client.Args.Set("page", fmt.Sprintf("%d", req.Page))
	}
	if req.PerPage != 0 {
		client.Args.Set("per_page", fmt.Sprintf("%d", req.PerPage))
	}
	client.OAuthSign()

	response := &PhotosSearchResponse{}
	err := flickr.DoPost(client, response)

	return response, err
}
