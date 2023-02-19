import { Link, useNavigate } from "react-router-dom";
import Logo from "../Assets/Logo.svg";
const NavBar = () => {
	
	return (
		<nav className="border-b-2 border-b-yellow-500 w-full flex flex-row justify-around bg-black text-white items-center py-2 transition-all z-10 shadow-b">
		
			<Link className="NavLink flex flex-col p-2 shadow font-semibold rounded" to="/">      
				<img className="w-10 h-10" src={Logo} alt="Logo." />
			</Link>
			
			<div className="flex justify-between items-center">
				<Link className="mx-4 NavLink flex flex-col p-2 shadow font-semibold rounded" to="/Home">  
					Home  
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				<Link className="mx-4 NavLink flex flex-col p-2 shadow font-semibold rounded" to="/">
					profile
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				<Link className="mx-4 NavLink flex flex-col p-2 shadow font-semibold rounded" to="/Store">
					Store 
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				<Link className="mx-4 NavLink flex flex-col p-2 shadow font-semibold rounded" to="/Store">
					About 
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				<Link className="mx-4 NavLink flex flex-col p-2 shadow font-semibold rounded" to="/Store">
					FAQ
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>
				
			</div>

		</nav>
	)
};

export default NavBar;
