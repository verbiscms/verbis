# Media

The media functions used within templates are used to retrieve items from the media library.

## Media Type
```

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

## Media Size Type
The m

```
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

Gets the media item from the library.

### Accepts: 

`float64` The media item ID.

### Returns:

`Media` A media struct.

### Examples:

**Get a media item with a given ID**

Get the media item with the ID of 10.

```
{{ getMedia 10 }}
```

**Assign to a variable:**

You can also assign the contents of a media item to a variable to be used later on in the template.

```
{{ $image := getMedia 10" }}
{{ $image.Url }}
```