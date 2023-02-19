package models

import (
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

type Product struct {
	Id			  		int           `json:"id"`
	Img			  		string	    `json:"img"`
	Owner_id			int          `json:"owner_id"`
	ProductLabel  		string 	    `json:"label"`
	ProductPrice  		string 	    `json:"price"`
	Desc		  		string 	    `json:"desc"`
	CreatedAt	  		time.Time	    `json:"created_at"`
	Vendor	      		string 	    `json:"vendor"`
}

func MakeProduct(Id int, Img string, Owner_id int, ProductLabel string, ProductPrice string, Desc string, Vendor string ) *Product {
	
	return &Product{
		Id: Id,
		Img: Img,
		Owner_id: Owner_id,
		ProductLabel: ProductLabel,
		ProductPrice: ProductPrice,
		Desc: Desc,
		CreatedAt: time.Now(),
		Vendor: Vendor,
	}

}

func (P *Product) FetchProductById(DBConn *sql.DB) bool {
	var ok = true;
	row, err := DBConn.Query("SELECT ID, IMG, OWNER_ID, PRODUCTLABEL, PRODUCTPRICE, DESC, CREATEDAT, VENDOR FROM PRODUCTS WHERE ID=? ORDER BY ID DESC", P.Id)
	defer row.Close()

	if err != nil {
		return !ok;
	}

	for row.Next() {
		row.Scan(P.Id, P.Img, P.Owner_id, P.ProductLabel, P.ProductPrice, P.Desc, P.CreatedAt, P.Vendor)
	}

	return ok;
}


func (P *Product) InsertToDatabase(DBConn *sql.DB) (error) {
	
	stmt, err := DBConn.Prepare("INSERT INTO PRODUCTS(IMG, OWNER_ID, PRODUCTLABEL, PRODUCTPRICE, DESC, CREATEDAT, VENDOR) VALUES(?, ?, ?, datetime(), ?)");
	
	if err != nil {
		return err;
	}

	_, err = stmt.Exec(P.Img, P.Owner_id, P.ProductLabel, P.ProductPrice, P.Desc, P.CreatedAt, P.Vendor);

	if err != nil {
		return err;
	}

	return nil;
}




/*

	check the balance of the buyyer. (  )
	
	if it is less than the price of the product then return balance_not_enough_for_transaction.
	else buy the product by changing the owner_id in the database. 
	increment the balance of the seller by the product's price then decriment the
	balance of the buyyer.

	this is the whole operation which is a bit clumsy but we can try other ways and locks etcetra....

*/

func (P *Product) BuyProduct(B *Buyyer, S *Seller) bool {  return nil }
