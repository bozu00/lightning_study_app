package models

import (
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"log"
	// "fmt"
	// "strings"
)


// func Delete(model Model) error {
// 	db := DBConnect()
// 	defer db.Close()
// 
// 	sql := `
// 	delete from ? 
// 	where id = ?;
//         `
// 	_, err := db.Exec(sql, model.TableName(), model.GetId())
// 	checkErr(err, "delete Image failed")
// 
// 	return err
// }


type Model interface {
    TableName() string 
	Key()       int64
}

type ModelBase struct {
	// Table  string
	Model	
	Id                int64  `db:"id"`
}

func (self ModelBase) Key() int64 {
	return self.Id
}

type Image struct {
	ModelBase
	Name              string `db:"name"`
	With     		  int    `db:"width"`
	Height            int    `db:"height"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}

func (self *Image) TableName() string {
	return "images"
}


// コンストラクタを追加するだけ
func NewImage() Image {
	// model := Model{Table: "images"}
	// return Image{Model: model}
	return Image{}
}

func Get(model Model, id int) error {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			target.*
        from ` + model.TableName() + ` as target
        where target.id=?
		limit 1
        `

	refv := reflect.ValueOf(model)
	objc := reflect.New(refv.Type()).Interface()

	err := db.Get(&objc, sql, id)
	checkErr(err, "fetch model failed")
	return err
}

func Delete(model Model) error {
	db := DBConnect()
	defer db.Close()

	sql := `
	delete from ` + model.TableName() + `
	where id = ?;
        `
	_, err := db.Exec(sql, model.Key())
	checkErr(err, "delete model failed")
	if err != nil {
		log.Println("delete model filed: " + err.Error())
	}

	return err
}

/*
func (self Model) UpdateModel() (error) {
	db := DBConnect()
	defer db.Close()
	// Value : ValueOf
	// Value : Indirect
    v := reflect.Indirect(reflect.ValueOf(self)) 
	// Type : Type
    t := v.Type()

	var parts []string

    for i := 0; i < t.NumField(); i++ {
        // フィールド名
        println("Field: " + t.Field(i).Name)
		// タグ
        println("Field: " + t.Field(i).Tag.Get("db"))


		// query = t.Field(i).Tag.Get("db") + " = " + v.Field(i).Interface.(string)
		query := fmt.Sprintf("%s = %s", t.Field(i).Tag.Get("db"), v.Field(i))

		parts = append(parts, query)
        // 値
		// Value : Field
        // f := v.Field(i) 
		// Interface : Interface (current value of f)
        // i := f.Interface()

		// 適切な型にキャストする？ しなくていい。プレースホルだに入れるだけだから
        // if value, ok := i.(int); ok {
        //     println("Value: " + strconv.Itoa(value))
        // } else {
        //     println("Value: " + f.String())
        // }
    }

	queryContents := strings.Join(parts, ",\n")
	log.Println(queryContents)

	sql := `
	update` + self.Table + `
	set` + queryContents + ` 
	where id = ?;`

	_,err := db.Exec(sql, self.Id)
	return err 

}
*/


func DeleteImage(id int64) error {
	db := DBConnect()
	defer db.Close()

	sql := `
	delete from images 
	where id = ?;
        `
	_, err := db.Exec(sql, id)
	checkErr(err, "delete Image failed")

	return err
}

func CreateImage(fileName string) (int64, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
	insert into images
	(name, width, height) Values
	(?, ?, ?) 
    `

	res,err := db.Exec(sql, fileName, 0, 0)
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


func GetImage(id int) (Image, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			images.*
        from images
        where images.id=?
		limit 1
        `

	var strct Image 
	err := db.Get(&strct, sql, id)

	checkErr(err, "fetch image failed")
	return strct, err
}

// func GetModel(typ reflect.Type, id int, obj interface{}) (error) {


func GetImages(limit int, offset int) ([]Image, error) {
	db := DBConnect()
	defer db.Close()

	sql := `
        select 
			images.*
        from images
		limit ?
		offset ?
        `

	var strcts []Image 
	err := db.Select(&strcts, sql, limit, offset)
	checkErr(err, "fetch image failed")
	return strcts, err
}

