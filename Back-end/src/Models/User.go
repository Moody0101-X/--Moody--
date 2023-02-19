package models

import (
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User_ struct {
	Id			int 	`json:"id"`
	Email		string 	`json:"email"`
	Name		string 	`json:"name"`
	Balance 	int 	`json:"balance"`
	PhoneNumber string 	`json:"phone_number"`
}

type User struct {
	Id			int 	`json:"id"`
	Email		string 	`json:"email"`
	Pwd     	string 	`json:"pwd"`	
	Name		string 	`json:"name"`
	Balance 	int 	`json:"balance"`
	PhoneNumber string 	`json:"phone_number"`
	JWT   string  `json:"jwt"`
}

// Seller and buyer are two struct that are used in handling sales operations between users......

type Seller struct {
	Identity User_ `json:"user"`
}

type Buyer struct {
	Identity     User_ 		  `json:"user"`
	JWT          string 	  `json:"jwt"`
	TimeStamp    time.Time    `json:"ts"`
	Target		 Product      `json:"product"`
}

func (U *User) Add(DBConn *sql.DB, StoreUserIdentityInJWT func(int, string, string) (string, error)) ServerResult {
	
	if IsEmpty(U.Email) || IsEmpty(U.Pwd) || IsEmpty(U.Name) || IsEmpty(U.PhoneNumber) {
		return NewServerResult(false, "(name | pwd | email | phone_number) not specified.", ServerCodes.Forbidden)
	}
	
	_, found := U.GetUserInformationFromDB(DBConn)
	
	if found {
		return NewServerResult(false, "this email is taken by another person.. please try another.", ServerCodes.Forbidden)
	} else {
		
		stmt, err := DBConn.Prepare("INSERT INTO USERS(EMAIL, PASSWORD, NAME, BALANCE, PHONENUMBER) VALUES(?, ?, ?, ?, ?)");
	
		if err != nil {
			return NewServerResult(false, "we could not sign you in, please try later.", ServerCodes.InternalServerError);
		}

		_, err = stmt.Exec(U.Email, U.Pwd, U.Name, 100, U.PhoneNumber);

		if err != nil {
			return NewServerResult(false, "we could not sign you in, please try later.", ServerCodes.InternalServerError);
		}

		U.Balance = 100;
		
		Token, err = StoreUserIdentityInJWT(A.Id, A.Pwd, A.Email);

		if err != nil {
			return NewServerResult(false, "Failed encoding the JWT.", ServerCodes.Forbidden)
		}
		
		U.JWT = Token

		return NewServerResult(true, "success", ServerCodes.OK)
	}
}

func (U User_) GetUserById(DBConn *sql.DB) ServerResult {
	return NewServerResult(false, "", ServerCodes.NotImplemented);
}

func (A *User) GetUserInformationFromDB(DBConn *sql.DB) (User, bool) {
	
	var U User;
	
	row, err := DBConn.Query("SELECT ID, EMAIL, PASSWORD, NAME, BALANCE, PHONENUMBER FROM USERS WHERE EMAIL=? ORDER BY ID DESC", A.Email)
	
	defer row.Close()

	if err != nil {
		return U, false;
	}

	for row.Next() {
		row.Scan(&U.Id, &U.Email, &U.Pwd, &U.Name, &U.Balance, &U.PhoneNumber)
		return U, true;
	}

	return U, false;
}

func (A *User) Validate(DBConn *sql.DB, CheckHash func(string, string) bool, GetUserIdentityFromJWT func(string) (User, bool), StoreUserIdentityInJWT func(int, string, string) (string, error)) ServerResult {
	/*
		Needs those two funcs as args.
		func CheckHash(Pwd string, Hash_ string) bool
		func GetUserIdentityFromJWT(TokenStr string) (models.Admin, bool)
		func StoreUserIdentityInJWT(Admin_id int, Admin_pwd string) (string, error)
		var ServerCodes CodeTable = CodeTable{
			OK: 200,
			Created: 201,
			Accepted: 202,
			NoContent: 204,
			MovedPermanently: 301,
			MovedTemporarily: 302,
			NotModified: 304,
			BadRequest: 400,
			Unauthorized: 401,
			Forbidden: 403,
			NotFound: 404,
			InternalServerError: 500,
			NotImplemented: 501,
			BadGateway: 502,
			ServiceUnavailable: 503,
		}
	*/

	if IsEmpty(A.JWT) {		
		if IsEmpty(A.Pwd) || IsEmpty(A.Email) {
			return NewServerResult(false, "(jwt | pwd | email) not specified.", ServerCodes.Unauthorized)
		}

		Target, ok := A.GetUserInformationFromDB(DBConn)

		if !ok {
			return NewServerResult(false, "Wrong Email, check your email then retry later.", ServerCodes.Unauthorized)
		}

		if !CheckHash(A.Pwd, Target.Pwd) {
			return NewServerResult(false, "Wrong Password. try again", ServerCodes.Unauthorized)
		}

		A.Id = Target.Id
		A.Name = Target.Name
		A.Balance = Target.Balance
		A.PhoneNumber = Target.PhoneNumber
		
		Token, err := StoreUserIdentityInJWT(A.Id, A.Pwd, A.Email)

		if err != nil {
			return NewServerResult(false, "Failed encoding the JWT.", ServerCodes.Forbidden)
		}

		A.JWT = Token

		return NewServerResult(true, "Logged in Successfully.", ServerCodes.OK)

	}

	_, ok := GetUserIdentityFromJWT(A.JWT)
	
	if !ok {
		return NewServerResult(false, "Invalid Jwt.", ServerCodes.Forbidden)
	}

	Target, _ := A.GetUserInformationFromDB(DBConn)

	if !CheckHash(A.Pwd, Target.Pwd) {
		return NewServerResult(false, "Wrong Password. try again", ServerCodes.Unauthorized)
	}

	A.Id = Target.Id
	A.Email = Target.Email
	A.Name = Target.Name
	A.Balance = Target.Balance
	A.PhoneNumber = Target.PhoneNumber
	
	return NewServerResult(true, "Logged in Successfully.", ServerCodes.OK)
}