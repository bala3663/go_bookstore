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
		results, err := db.Query("SELECT * FROM bookstable")
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
		query_data := fmt.Sprintf(`INSERT INTO bookstable (BookId,BookName,AuthorName,Publication,Year) VALUES(%d,"%s","%s","%s",%d)`, new_book.BookId, new_book.BookName, new_book.AuthorName, new_book.Publication, new_book.Year)

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
		query := fmt.Sprintf("UPDATE bookstable SET BookName = '%s',AuthorName = '%s',Publication ='%s',year = '%d' WHERE BookId = '%d' ", edit_book.BookName, edit_book.AuthorName, edit_book.Publication, edit_book.Year, edit_book.BookId)
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
		query := fmt.Sprintf(" DELETE FROM bookstable WHERE BookId=%d ", delete_book.BookId)
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
		query_data := fmt.Sprintf("SELECT * FROM bookstable WHERE BookId='%d'", search_byid.BookId)
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
		query_data := fmt.Sprintf("SELECT * FROM bookstable WHERE BookName='%s'", search_by_name.BookName)
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
		query_data := fmt.Sprintf("SELECT * FROM bookstable WHERE Publication ='%s'", search_book_publication.Publication)
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
		query_data := fmt.Sprintf("SELECT * FROM bookstable WHERE Year  BETWEEN '%d' AND '%d'", publication_year.Year, publication_year.Year+5)
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
			var Discount int
			var BookId int

			err = results.Scan(&BookNo, &BookName, &BookPrice, &Discount, &BookId)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' %d %d %d ", BookNo, BookName, BookPrice, Discount, BookId)
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
		query_data := fmt.Sprintf(`INSERT INTO bookprice (BookNo,BookName,BookPrice,Discount,BookId) VALUES(%d,"%s", %d, %d,%d)`, Update_book_Price.BookId, Update_book_Price.BookName, Update_book_Price.BookPrice, Update_book_Price.Discount, Update_book_Price.BookId)

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		c.IndentedJSON(201, "Yes, values added!")
	}

}

func Book_Awards() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var Book_Awards models.Info2
		err = c.BindJSON(&Book_Awards)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf(`INSERT INTO bookawards (BookNumber,BookName,Year,Award,book_category,summary,BookId) VALUES(%d,"%s", %d, "%s","%s","%s",%d)`, Book_Awards.BookNunber, Book_Awards.BookName, Book_Awards.Year, Book_Awards.Award, Book_Awards.Book_category, Book_Awards.Summary, Book_Awards.BookId)

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		c.IndentedJSON(201, "Yes, values added!")
	}

}

func Book_Infor() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/studentinfor")
		if err != nil {
			panic(err.Error())
		}
		var Book_Infor models.Info2
		err = c.BindJSON(&Book_Infor)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM bookawards WHERE Award ='%s'", Book_Infor.Award)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var BookNumber int
			var BookName string
			var Year int
			var Award string
			var Book_category string
			var Summary string
			var BookId int

			err = results.Scan(&BookNumber, &BookName, &Year, &Award, &Book_category, &Summary, &BookId)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' %d '%s' '%s' '%s'  %d ", BookNumber, BookName, Year, Award, Book_category, Summary, BookId)
			c.IndentedJSON(200, "Book")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}
