
import { useSelector, useDispatch } from "react-redux";

const Profile = () => {
	const User = useSelector(state => state.User);
	const Dispatch = useDispatch();

	return (
		<h1> Profile Page. </h1>
	)

}


export { Profile };
