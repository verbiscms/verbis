# Posts

The post functions play a vital role in obtaining resource and post data from Verbis. 
With the functions outlined below, you are able to retrieve resources and pages to allow you to create
archive pages and display rich data via the `VerbisLoop`. Pagination data can also be retrieved 
allowing the user to filter through pages of your content.

## Post Struct
The post type is what is returned when getting singular or multiple posts from the database.
You are able to access any of the properties below to output your desired content.

```go
type Post struct {
      Id                int           
      UUID              uuid.UUID  
      Slug              string        
      Permalink         string   
      Title             string     
      Status            string        
      Resource          string        
      PageTemplate      string        
      PageLayout        string      
      CodeInjectionHead string        
      CodeInjectionFoot string      
      UserId            int         
      IsArchive         bool
      PublishedAt       *time.Time    
      CreatedAt         time.Time     
      UpdatedAt         time.Time    
      SeoMeta           PostOptions  
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

## post

Retrieves a post item from the database by ID.

### Accepts: 

`id integer` The post ID.

### Returns:

`Post` The post type or nil if no post was found.

### Examples:

**Get the a post by ID**

This example demonstrates how to obtain a post by ID. All the data from the post will be displayed
(see above).

```gotemplate
{{ post 123 }}
```

**Assign a post to a variable**

This example demonstrates how to obtain a post by ID and assign the contents to a variable.
The `$post` variable contains all the data you need to get started with displaying useful
information in the theme, such as a slug or title.

```gotemplate
{{ $post := post 123 }}
{{ $post.Slug }}
{{ $post.Slug }}
```
___

## posts

Retrieves multiple post items from the database by a `dict` query. Default values from the return
are shown below.

### Accepts: 

`Params (dict map[string]interface{})` Query parameters using a `dict`with default
properties shown below.

```go
type Params struct {
    Page 		int         `default: 1`
    Limit 		int         `default: 15`
    Resource 		string      `default: all`
    OrderBy 		string      `default: published_at`
    OrderDirection 	string      `default: desc`
    Category     	string      `default: nil`
}
```

### Returns:

```go
map[string]interface{}{
    "Pagination": Pagination
    "Posts": []Posts
}, error
``` 
A map containing the pagination and an array of posts, or an error if there was a problem
unmarshalling to the params type. By default, only published pages will be returned.

### Examples:

**Get a limited number of posts by a resource**

This example demonstrates how to obtain 10 posts with the resource of type posts.

First a query is assigned to type `dict`; here is where we define the type of data we
expect to receive from the database. A limit of 10 means we will only retrieve 10 items.

We then assign the result to a variable and pass the query into the `posts` function.
The `$result` variable will now contain an `map[string]interface{}` of the Pagination 
& Post types (as shown above).

A good idea is to check if any posts exists by using an `with` statement, so a notification
can be displayed if none were found. A range loop can be used to loop over the posts and 
output any data.

```gotemplate
{{ $query := dict "limit" 10 "resource" "posts" }}
{{ $result := post (dict "limit" 10 "resource" "posts") }}
{{ with $result.Posts }}
    {{ range $post := . }}
        <h2>{{ $post.Title }}</h2>
        <a href="{{ $post.Slug }}">Read more</a>
    {{ end }}
{{ else }}
    <h4>No posts found</h4>
{{ end }}
```

**Get posts with a specific order direction**

This example demonstrates how to obtain posts with no particular resource, but order them by
`title` in `ascending` direction.

```gotemplate
{{ $query := dict "order_by" "title" "order_direction" "asc" }}
{{ $result := posts $query }}
{{ with $result.Posts }}
    {{ range $post := . }}
        <h2>{{ $post.Title }}</h2>
        <a href="{{ $post.Slug }}">Read more</a>
    {{ end }}
{{ else }}
    <h4>No posts found</h4>
{{ end }}
```

**Get posts with a specific category**

This example demonstrates how to obtain posts with the resource of type posts that have a category
of sports.

```gotemplate
{{ $query := dict "resource" "posts" "category" "sports" }}
{{ $result := posts $query }}
{{ with $result.Posts }}
    {{ range $post := . }}
        <h2>{{ $post.Title }}</h2>
        <a href="{{ $post.Slug }}">Read more</a>
    {{ end }}
{{ else }}
    <h4>No sports found</h4>
{{ end }}
```

___

## paginationPage

Retrieves the page query parameter, useful for working with the `posts` function.

### Returns:

`int` The page query parameter if found or 1 if it wasn't or there was a problem 
casting the page to an `integer`.

### Examples:

**Get pagination page**

This example demonstrates how to obtain the pagination `page` query parameter assuming
the url looks similar to `/posts?page=2`, the function will return two.

```gotemplate
{{ paginationPage }}
```

**Display pagination**

This example demonstrates how to display pagination so users can filter through posts in your theme.

First the `$page` variable is assigned to the `paginationPage` call which simply returns the page query
parameter. It's important to pass this variable into your query, this will ensure the correct content is
returned from the database.

After the loop we can check we can use the Pagination object to determine if there are previous and next 
pages available. The `Prev`and `Next` variables are an `interface{}` so they can be an integer (the corresponding)
 page or a `false` if there is no page available in the given direction.
 
We can check this using a simple`if statement` and output a link with the query parameter, for example:
`<a href="/posts?page={{ $result.Pagination.Prev }}">Previous Post</a>`

```gotemplate
{{ $page := paginationPage }}
{{ $query := dict "page" $page "resource" "posts" limit 10 }}
{{ $result := posts $query }}

{{ with $result.Posts }}
    {{ range $post := . }}
        <h2>{{ $post.Title }}</h2>
        <a href="{{ $post.Slug }}">Read more</a>
    {{ end }}
{{ else }}
    <h4>No posts found</h4>
{{ end }}

<div class="pagination">
   {{ if $result.Pagination.Prev }}
       <a href="/posts?page={{ $result.Pagination.Prev }}">Previous Post</a>
   {{ end }}
   {{ if $result.Pagination.Next }}
       <a href="/posts?page={{ $result.Pagination.Next }}">Next Post</a>
   {{ end }}
</div>
```