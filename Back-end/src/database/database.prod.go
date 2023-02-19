
package database

import (
	"moody.com/api/models"
)

// Seller, Buyer, Owner


func GetAllProducts() (error, []models.Product) {

	var all []models.Product;
	var tmp models.Product;
	
	rows, err := DATABASE.Query("SELECT ID, IMG, OWNER_ID, PRODUCTLABEL, PRODUCTPRICE, DESC, CREATEDAT, VENDOR FROM PRODUCTS ORDER BY ID DESC");

	defer rows.Close()

	if err != nil {
		return err, all;
	}

	for rows.Next() {
		rows.Scan(&tmp.Id, &tmp.Img, &tmp.Owner_id, &tmp.ProductLabel, &tmp.ProductPrice, &tmp.Desc, &tmp.CreatedAt, &tmp.Vendor)
		all = append(all, tmp);
	}

	return nil, all;
	
}

func GetProductByID(Prod models.Product) bool {
	ok := Prod.FetchProductById(DATABASE);
	return ok;
}