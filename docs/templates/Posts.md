    {{ include "partials/header" }}
    {{ template "content" . }}
    {{ $page := getPagination }}
    <h1>{{ $page }}</h1>
    {{ $query := dict "limit" 100 "resource" "posts" "page" 1 }}
    {{ $posts := getPosts $query }}

    {{ if $posts }}
        <pre>{{ $posts.Pagination }}</pre>
    {{ else }}
        <h1>No posts found</h1>
    {{ end }}
