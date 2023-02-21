
package database

import (
	"moody.com/api/models"
	"moody.com/api/images"
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
	ok := Prod.FetchProductById(GetDb());
	return ok;
}


func AddNewProduct(P *models.Product) models.ServerResult {
	// Add img to cdn.
	var NewId int = get_id("PRODUCTS");
	
	P.Id = NewId;
	ok, ImgPath := images.SaveImage(P.Id, P.Img);

	// Reset img property to url path pushed from cdn.
	if ok {
		err := P.InsertToDatabase(GetDb());

		if err != nil {
			return models.NewServerResult(false, "product Could not be added.", models.ServerCodes.InternalServerError);
		}

		P.Img = ImgPath;
		return models.NewServerResult(true, "success.", models.ServerCodes.OK);
	}	

	return models.NewServerResult(false, "You product image could not be added to the cdn, try again later..", models.ServerCodes.InternalServerError);
}