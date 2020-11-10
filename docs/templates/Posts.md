# Posts

The post functions play a vital role in obtaining resource and post data from Verbis. 
With the functions outlined below, you are able to retrieve resources and pages to allow you to create
archive pages and display rich data via the `VerbisLoop`. Pagination data can also be retrieved 
allowing the user to filter through pages of your content.

## Post Struct
The post type is what is returned when getting singular or multiple posts from the database/
You are able to access any of the properties below to output your desired content.

```go
type Post struct {
 	Id		        int							
 	UUID                    uuid.UUID				
 	Slug			string 						
 	Title			string 				
 	Status			string						
 	Resource		*string 					
 	PageTemplate            string				
 	Layout			string						
 	UserId 			int 						
 	PublishedAt		*time.Time					
 	CreatedAt		*time.Time					
 	UpdatedAt		*time.Time					
 }
```

## Pagination Struct 

The pagination type is returned when obtaining multiple records of posts from the database.
It's useful for determining what page to go to, the total number of posts, and whether 
to display previous & next buttons making it easy for your users to filter through 
dynamic content.

```go
type Pagination struct {
	Page 		int					
	Pages 		int					
	Limit 		int					
	Total 		int					
	Next 		interface{} 		
	Prev 		interface{} 		
}
```

___

## getPost

Retrieves a post item from the database by ID.

### Accepts: 

`integer` The post ID.

### Returns:

`Post` The post type or nil if no post was found.

### Examples:

**Get the a post by ID**

This example demonstrates how to obtain a post by ID. All the data from the post will be displayed
(see above).

```gotemplate
{{ getPost 123 }}
```

**Assign a post to a variable**

This example demonstrates how to obtain a post by ID and assign the contents to a variable.
The `$post` variable contains all the data you need to get started with displaying useful
information in the theme, such as a slug or title.

```gotemplate
{{ $post := getPost 123 }}
{{ $post.Slug }}
{{ $post.Slug }}
```
___

## getPosts

Retrieves multiple post items from the database by a `dict` query.

### Accepts: 

`Params (dict map[string]interface{})` Query parameters using a `dict` shown below.

```go
type params struct {
    Page 		int
    Limit 		int
    Resource 		string
    OrderBy 		string
    OrderDirection 	string
}
```

### Returns:

```go
map[string]interface{}{
    "Pagination": Pagination
    "Posts": []Posts
}, error
``` 
A map containing the pagnation and an array of posts or an error if there was a problem
unmarshalling to the params type.









    {{ $page := getPagination }}
    <h1>{{ $page }}</h1>
    {{ $query := dict "limit" 100 "resource" "posts" "page" 1 }}
    {{ $posts := getPosts $query }}

    {{ if $posts.Items }}
        <pre>{{ $posts.Pagination }}</pre>
    {{ else }}
        <h1>No posts found</h1>
    {{ end }}
