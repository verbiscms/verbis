# Media

These media functions within templates are used to retrieve items from the media library allowing you 
to output any type of media file into your template.
See below for respective return types & data.

## Media Struct
The media type is what is returned when calling `getMedia`. You are able to access any of the
properties below to output your desired content. 

```go
type Media struct {
	Id				int				
	UUID 			uuid.UUID			
	Url 			string				
	Title			*string 			
	Alt				*string 			
	Description		*string 			
	FileSize		int 			
	FileName		string 				
	Sizes 			MediaSizes		 	
	Type			string 				
	UserID			int					
	CreatedAt		time.Time			
	UpdatedAt		time.Time		
}
```

## Media Size Struct
The media size type contains useful data for displaying different images for different view ports using the
`<picture>` element.
To access the media sizes from a media item, you need to pass the name of the size as the MediaSizes
type is a `map[string]MediaSizes`, for example: `{{ $mymedia.Sizes.thumbnail }}`

```go
type MediaSize struct {
	UUID 			uuid.UUID		
	Url 			string				
	Name			string 				
	SizeName 		string 				
	FileSize		int 				
	Width			int 				
	Height			int 				
	Crop			bool 				
}
```
___

## getMedia

Get the media item from the library.

### Accepts: 

`float64` The media item ID.

### Returns:

`Media` A media struct.

### Examples:

**Get a media item with a given ID**

Get the media item with the ID of 10.

```gotemplate
{{ getMedia 10 }}
```

**Assign to a variable:**

You can also assign the contents of a media item to a variable to be used later on in the template.

```gotemplate
{{ $image := getMedia 10 }}
{{ $image.Url }}
```

**Access from a field**

The `getField` function returns a media library item if it's a type of media.
See `getField` for more information.

```gotemplate
{{ $image := getField "image" }}
{{ if $image }}
    {{ $image.Url }}
{{ end }}
```

**Access a specific image size**

Get the `thumnbail` size of an image.

```gotemplate
{{ $image := getField "image" }}
{{ if $image }}
    {{ if $image.Sizes.thumbnail }}
        {{ $image.Sizes.thumbnail }}
    {{ end }}
{{ end }}
```
