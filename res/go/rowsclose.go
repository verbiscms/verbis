rows := s.db.QueryRowx("SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta' FROM posts LEFT JOIN post_options ON posts.id = post_options.page_id WHERE posts.slug = ? LIMIT 1", slug)
//for rows.Next() {
err := rows.StructScan(&p)
if err != nil {
fmt.Println(err)
}
//}
err := rows.Close()
if err != nil {
fmt.Println(err)
}

