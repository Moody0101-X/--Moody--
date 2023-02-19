package database


import (
	"moody.com/api/models"
	"moody.com/api/auth"
)

func AuthenticateUser(U *models.User) models.ServerResult {
	Result := U.Validate(DATABASE, auth.CheckHash, auth.GetUserIdentityFromJWT, auth.StoreUserIdentityInJWT);
	return Result;
}

func SignUpAUser(U *models.User) models.ServerResult {
	Result := U.Add(DATABASE, auth.StoreUserIdentityInJWT)
	
	if Result.Ok {
		U.Id = get_id("USERS");
	}

	return Result;
}
















