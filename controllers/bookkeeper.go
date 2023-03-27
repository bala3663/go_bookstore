package controllers

import (
	"Final-Project-gin/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" //high performance HTTP(hypertext transfer protocol) web framework
	_ "github.com/go-sql-driver/mysql"
)

//get method

func Get_books() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		results, err := db.Query("SELECT * FROM bookmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookId int
			var BookName string
			var AuthorName string
			var Publication string
			var Year int
			err = results.Scan(&BookId, &BookName, &AuthorName, &Publication, &Year)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d %s %s %s %d ", BookId, BookName, AuthorName, Publication, Year)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

// post method

func Post_book() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var new_book models.Info
		err = c.BindJSON(&new_book)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf(`INSERT INTO bookmanagement (BookId,BookName,AuthorName,Publication,Year) VALUES(%d,"%s","%s","%s",%d)`, new_book.BookId, new_book.BookName, new_book.AuthorName, new_book.Publication, new_book.Year)

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		c.IndentedJSON(201, "Yes, values added!")
	}

}

// Put Method

func Update_book() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("database connected")
		defer db.Close()
		var edit_book models.Info
		err = c.BindJSON(&edit_book)
		if err != nil {
			return
		}
		query := fmt.Sprintf("UPDATE bookmanagement SET BookName = '%s',AuthorName = '%s',Publication ='%s',year = '%d' WHERE BookId = '%d' ", edit_book.BookName, edit_book.AuthorName, edit_book.Publication, edit_book.Year, edit_book.BookId)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(201, "Yes, Successfully Updated")
	}
}

func Delete_book() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("database connected")
		defer db.Close()
		var delete_book models.Info
		err = c.BindJSON(&delete_book)
		if err != nil {
			return
		}
		query := fmt.Sprintf(" DELETE FROM bookmanagement WHERE BookId=%d ", delete_book.BookId)
		_, err = db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, "Successfully Deleted")
	}
}

func Search_book_id() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var search_byid models.Info
		err = c.BindJSON(&search_byid)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookmanagement WHERE BookId='%d'", search_byid.BookId)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookId int
			var BookName string
			var AuthorName string
			var Publication string
			var Year int
			err = results.Scan(&BookId, &BookName, &AuthorName, &Publication, &Year)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' %d ", BookId, BookName, AuthorName, Publication, Year)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

func Search_book_name() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var search_by_name models.Info
		err = c.BindJSON(&search_by_name)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookmanagement WHERE BookName='%s'", search_by_name.BookName)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookId int
			var BookName string
			var AuthorName string
			var Publication string
			var Year int
			err = results.Scan(&BookId, &BookName, &AuthorName, &Publication, &Year)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' %d ", BookId, BookName, AuthorName, Publication, Year)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

func Search_book_publication() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var search_book_publication models.Info
		err = c.BindJSON(&search_book_publication)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookmanagement WHERE Publication ='%s'", search_book_publication.Publication)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookId int
			var BookName string
			var AuthorName string
			var Publication string
			var Year int
			err = results.Scan(&BookId, &BookName, &AuthorName, &Publication, &Year)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' %d ", BookId, BookName, AuthorName, Publication, Year)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

func Publication_Year() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var publication_year models.Info
		err = c.BindJSON(&publication_year)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookmanagement WHERE Year  BETWEEN '%d' AND '%d'", publication_year.Year, publication_year.Year+5)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookId int
			var BookName string
			var AuthorName string
			var Publication string
			var Year int
			err = results.Scan(&BookId, &BookName, &AuthorName, &Publication, &Year)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' %d ", BookId, BookName, AuthorName, Publication, Year)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

func Book_price() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var Book_price models.Info1
		err = c.BindJSON(&Book_price)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookprice WHERE BookPrice ='%d'", Book_price.BookPrice)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookNo int
			var BookName string
			var BookPrice int
			var BookId int
			err = results.Scan(&BookNo, &BookName, &BookPrice, &BookId)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' %d %d ", BookNo, BookName, BookPrice, BookId)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

func Update_book_Price() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var Update_book_Price models.Info1
		err = c.BindJSON(&Update_book_Price)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf(`INSERT INTO BookPrice (BookNo,BookName,BookPrice,BookId) VALUES(%d,"%s", %d, %d)`, Update_book_Price.BookId, Update_book_Price.BookName, Update_book_Price.BookPrice, Update_book_Price.BookId)

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		c.IndentedJSON(201, "Yes, values added!")
	}

}
