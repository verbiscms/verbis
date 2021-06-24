# Frontend Routes

Front end routes are dedicated endpoints to serve dynamic content through Verbis such as assets, upload and sitemaps.

## Verbis

🔗 `/verbis`

Any route that is prefixed with `/verbis` is reserved by the system and is used for serving emails, content and
debugging assets. This cannot be changed.

## Assets

🔗 `/assets`

The assets' directory is dynamic and can be changed through the `config.yml` file. By default, Verbis looks for
the `assets` directory. Any route that contains `/assets` will be automatically resolved to that directory.

### Example:

If you have set the assets' path to `public` within the `config.yml` file, and visit `/assets/images/cows.png` the file
will be resolved. Similarly, if you visit `/assets/css/styles.css` the `styles.css` will be resolved.

```
mytheme
│   config.yml
│
└───public
│   │
│   └───images
│   │   │   cow.png
│   │   │   ...
|   └───css
│       │   styles.css
│       │   ...
│   
```

## Uploads

The upload path can be changed in the `config.yml` file and defines where the media will be requested from the front-end
of the website. For example, a file requested with the  `/uploads` would have the url of to `/uploads/2020/09/photo.jpg`

## Favicon

🔗 `/favicon.ico`

The `/favicon.ico` url will look for the `favicon.ico` file in the active Theme's directory. If no favicon is found, a
`404` will be returned.

## Robots

🔗 `/robots.txt`

The `robots.txt` file can be directly edited in the Verbis admin interface and as such the `/robots.txt` route is
reserved by the system to serve the file.

## Sitemaps

🔗 `/{sitemaps}`

Sitemaps in Verbis are automatically generated, you can exclude individual posts by clicking the `SEO` tab in the
editor. The main sitemap is an index of individual smaller sitemaps which are categorised by resources. Below is a list
of routes and files that are used to serve sitemaps.

- `/sitemap.xml`: The main index sitemap.
- `/sitemaps/:resource/:map`: Individual sitemaps by resource.
- `/main-sitemap.xsl`: Styling `.xsl` file for the index.
- `/resource-sitemap.xsl`: Styling `.xsl` file for the resource sitemaps.



