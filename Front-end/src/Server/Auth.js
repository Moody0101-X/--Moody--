import { Post } from "./Requests";
import { API } from "../constants/serverConstants";
import { getJwtAuthToken } from "../clientOperations/cookies"

var jwt = getJwtAuthToken();

const AuthUserWithJWT = async () => await Post(`${API}/User/Auth/Login`, { jwt });
const AuthUserWithCredentials = async (email, pwd) => await Post(`${API}/User/Auth/Login`, { email, pwd });
const SignUpUser = async (email, pwd, name, phone_number) => await Post(`${API}/User/Auth/SignUp`, { email, pwd, name, phone_number });

export {
	AuthUserWithJWT,
	AuthUserWithCredentials,
	SignUpUser,
	getJwtAuthToken
}