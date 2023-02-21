import { Post } from "./Requests";
import { getJwtAuthToken } from "../clientOperations/cookies";
import { API } from "../Constants/serverConstants";
var jwt = getJwtAuthToken();

const AddProduct = async (img, owner_id, label, price, vendor) => await Post(`${API}/Product/Add`, { img, owner_id, label, price, vendor, jwt});
const GetProducts = async () => await Get(`${API}/Products`);
