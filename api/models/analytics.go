package models

import (
	"github.com/jmoiron/sqlx"
)

type AnalyticsStore struct {
	db *sqlx.DB
}

//type Category struct {
//	ID			int			`db:"id" json:"id"`
//	Name		string 		`db:"name" json:"name" binding:"required"`
//	CreatedAt	time.Time	`db:"created_at" json:"created_at"`
//	UpdatedAt	time.Time	`db:"updated_at" json:"updated_at"`
//}

type AnalyticsRepository interface {

}

//Construct
func newAnalytics(db *sqlx.DB) *AnalyticsStore {
	return &AnalyticsStore{
		db: db,
	}
}

// Create category
//func (s *AnalyticsStore) UpdatePageView(slug string) error {
//
//
//	q := "UPDATE posts SET page = ?, title = ?, category_id = ?, resourxce_id = ?, content = ?, status = ?, updated_at = NOW() WHERE id = ?"
//	_, err = s.db.Exec(q, p.Slug, p.Title, p.CategoryID, 1, p.Content, 0)
//	if err != nil {
//		return fmt.Errorf("Could not update the post %v - %w", p.Title, err)
//	}
//
//	return nil
//}


/*
	TODO:
		-	Add current people on the website right now (session data).
		-	How many visitors in a month to plot on a graph.
		- 	Look for any other analytics opportunities.
 */