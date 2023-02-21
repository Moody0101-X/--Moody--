package database
import (
	"moody.com/api/models"
	"moody.com/api/auth"
	"fmt"
)

func AuthenticateUser(U *models.User) models.ServerResult {
	Result := Validate(U);
	return Result;
}

func SignUpAUser(U *models.User) models.ServerResult {

	
	Result := Add(U);
	
	if Result.Ok {
		U.Id = get_id("USERS");
	}

	return Result;
}

func GetUserInformationFromDB(Email string) (models.User, bool) {
	
	var U models.User;
	
	rows, err := DATABASE.Query("SELECT ID, EMAIL, PWD, NAME, BALANCE, PHONENUMBER FROM USERS WHERE EMAIL=? ORDER BY ID DESC", Email)

	if err != nil  {
		return U, false;
	}

	defer rows.Close()

	for rows.Next() {
		fmt.Println(U)
		rows.Scan(&U.Id, &U.Email, &U.Pwd, &U.Name, &U.Balance, &U.PhoneNumber)
		return U, true;
	}

	return U, false;

}

func Add(User_ *models.User) models.ServerResult {
	
	if models.IsEmpty(User_.Email) || models.IsEmpty(User_.Pwd) || models.IsEmpty(User_.Name) || models.IsEmpty(User_.PhoneNumber) {
		return models.NewServerResult(false, "(name | pwd | email | phone_number) not specified.", models.ServerCodes.Forbidden)
	}

	_, found := GetUserInformationFromDB(User_.Email);
	
	if found {
		return models.NewServerResult(false, "this email is taken by another person.. please try another.", models.ServerCodes.Forbidden)
	}
		
	stmt, err := DATABASE.Prepare("INSERT INTO USERS(EMAIL, PWD, NAME, BALANCE, PHONENUMBER) VALUES(?, ?, ?, ?, ?)");
	
	if err != nil {
		fmt.Println(err);
		return models.NewServerResult(false, "we could not sign you in, please try later.", models.ServerCodes.InternalServerError);
	}

	_, err = stmt.Exec(User_.Email, auth.HashPwd(User_.Pwd), User_.Name, 100, User_.PhoneNumber);

	if err != nil {
		return models.NewServerResult(false, "we could not sign you in, please try later.", models.ServerCodes.InternalServerError);
	}

	User_.Balance = 100;
	
	Token, err := auth.StoreUserIdentityInJWT(User_.Id, User_.Pwd, User_.Email);

	if err != nil {
		return models.NewServerResult(false, "Failed encoding the JWT.", models.ServerCodes.Forbidden)
	}
	
	User_.JWT = Token

	return models.NewServerResult(true, "success", models.ServerCodes.OK)

}


func Validate(A *models.User) models.ServerResult {
	if models.IsEmpty(A.JWT) {
		if models.IsEmpty(A.Pwd) || models.IsEmpty(A.Email) {
			return models.NewServerResult(false, "(jwt | pwd | email) not specified.", models.ServerCodes.Unauthorized)
		}

		Target, ok := GetUserInformationFromDB(A.Email)

		if !ok {
			return models.NewServerResult(false, "Wrong Email, check your email then retry later.", models.ServerCodes.Unauthorized)
		}

		if !auth.CheckHash(A.Pwd, Target.Pwd) {
			return models.NewServerResult(false, "Wrong Password. try again", models.ServerCodes.Unauthorized)
		}

		A.Id = Target.Id
		A.Name = Target.Name
		A.Balance = Target.Balance
		A.PhoneNumber = Target.PhoneNumber
		
		Token, err := auth.StoreUserIdentityInJWT(A.Id, A.Pwd, A.Email)

		if err != nil {
			return models.NewServerResult(false, "Failed encoding the JWT.", models.ServerCodes.Forbidden)
		}

		A.JWT = Token

		return models.NewServerResult(true, "Logged in Successfully.", models.ServerCodes.OK)

	}

	User, ok := auth.GetUserIdentityFromJWT(A.JWT)
	
	if !ok {
		return models.NewServerResult(false, "Invalid Jwt.", models.ServerCodes.Forbidden)
	}

	Target, _ := GetUserInformationFromDB(User.Email);
	
	A.Id = Target.Id
	A.Email = Target.Email
	A.Name = Target.Name
	A.Balance = Target.Balance
	A.PhoneNumber = Target.PhoneNumber

	return models.NewServerResult(true, "Logged in Successfully.", models.ServerCodes.OK)
}