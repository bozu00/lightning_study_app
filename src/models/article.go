package models

import (
	_ "github.com/go-sql-driver/mysql"
	// "../formModel"
    "virtualhost.local/kirakira/lightning_study_app/src/formModel"
	//"strconv"
	//"log"
)

type Article struct {
	Id                int64  `db:"id"`
	Title             string `db:"title"`
	Description		  string `db:"description"`
	TableOfContents   string `db:"table_of_contents"`
	EyeCatchImageID   int    `db:"eyecatch_image_id"`
	Content           string `db:"content"`
	AuthorId          int64  `db:"author_id"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}


type ArticleWithAuthor struct {
	Id                int64  `db:"id"`
	Title             string `db:"title"`
	Content           string `db:"content"`
	ArticleCategoryId int64  `db:"article_category_id"`
	MediaId           int64  `db:"media_id"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
	AuthorId          int64  `db:"author_id"`
	Author            `db:"author"`
}

type Author struct {
	Id        int64  `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func GetArticles(limit int, offset int) ([]Article, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			article.*
        from articles as article
        order by article.updated_at   
		limit ?
		offset ?
        `
	var articles []Article
	err := db.Select(&articles, sql, limit, offset)
	checkErr(err, "fetch Articles failed")

	return articles, err
}

func GetArticle(article_id int) (Article, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			article.*
        from articles as article
        where article.id=?
		limit 1
        `

	var article Article 
	err := db.Get(&article, sql, article_id)
	checkErr(err, "fetch Articles failed")

	return article, err
}

func CreateArticle(form formModel.NewArticle) (int64, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
	insert into articles 
	(title, description, table_of_contents, content, author_id, eyecatch_image_id) Values 
	(?, ?, ?, ?, ?, ?); 
	`

	res,err := db.Exec(sql, form.Title, form.Description, form.TableOfContents, form.Content, 1, form.EyecatchImageId)
	logger(res)
	if !checkErr(err, "create user failed") {
		return -1, err
	}

	id, err := res.LastInsertId()
	if !checkErr(err, "get last insert id failed") {
		return -1, err
	}
	logger(id)

	return id, err
}

func UpdateArticle(form formModel.UpdateArticle) (int64, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
	update articles 
	set title = ?, 
		description = ?, 
		table_of_contents = ?, 
		content = ?, 
		author_id = ?,
		eyecatch_image_id = ?
	where id = ?;
	`

	res,err := db.Exec(sql, form.Title, form.Description, form.TableOfContents, form.Content, 1, form.EyecatchImageId, form.Id)
	logger(res)
	if !checkErr(err, "update user failed") {
		return -1, err
	}

	return form.Id, err

}
