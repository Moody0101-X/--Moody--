import { Link, useNavigate } from "react-router-dom";
import Logo from "../Assets/Logo.svg";
import { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';

const NavBar = () => {
	const User = useSelector(state => state.User);
	const dispatch = useDispatch();
	
	const [show, setShow] = useState(false);
	const showNavbarMobile = () => {
		setShow(p => !p);
	}

	const OnLink = (e) => {
		if(show) setShow(false);	
	}

	return (
		<nav className="border-b-2 border-b-yellow-500 h-[70px] w-full flex flex-row justify-around bg-black text-white items-center py-2 transition-all z-10 shadow-b">
			
			<Link className="NavLink bg-black fixed left-2 sm:static flex flex-col p-2 shadow font-semibold rounded" to="/">      
				<img className="w-10 h-10" src={Logo} alt="Logo." />
			</Link>
			
			<div className={`transition-all ease-in-out flex flex-col items-center justify-center fixed z-10 top-0 right-0 w-screen h-screen bg-black ${(show) ? "translate-x-0" : "sm:translate-x-0 translate-x-[100%]"} sm:h-full sm:w-1/2 sm:bg-black sm:static sm:flex sm:flex-row sm:justify-between sm:items-center`}>
				
				<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 NavLink flex flex-col p-2 font-semibold rounded" to="/Home">  
					Home  
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				{
					(User) ? (
						
						<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 NavLink flex flex-col p-2 font-semibold rounded" to="/profile">
							profile
							<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
						</Link>

					) : ""
				}

				<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 NavLink flex flex-col p-2 font-semibold rounded" to="/Store">
					Store 
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>

				<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 NavLink flex flex-col p-2 font-semibold rounded" to="/Store">
					About 
					<span className="bg-white h-[1px] transition-all ease-in-out w-[0]">  </span>
				</Link>
				{
					(User) ? (
						<>
							<div className="sm:my-0 my-3 mx-4 NavLink flex flex-row p-1 font-semibold rounded bg-rose-600 items-center">				
								{ User.name }
							</div>
							
							<p className="text-sm bg-blue-500 p-1 rounded"> uid: #{ User.id } </p>
							
							<button className="sm:my-0 my-3 mx-4 NavLink flex flex-col p-2 font-semibold rounded bg-rose-600">
								Logout
							</button>
						</>
					) : (
						<>
							<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 transition-all ease-in-out flex flex-col p-2 font-semibold rounded border-b border-b-blue-600 hover:bg-slate-800 bg-slate-900" to="/Login">
								Login
							</Link>

							<Link onClick={OnLink} className="sm:my-0 my-3 mx-4 transition-all ease-in-out flex flex-col p-2 font-semibold rounded border-b border-b-blue-600 hover:bg-slate-800 bg-slate-900" to="/Signup">
								Sign up
							</Link>
						</>
					)
				}

			</div>

			<div onClick={showNavbarMobile} className="cursor-pointer z-20 fixed right-2 p-2 w-[50px] flex flex-col sm:hidden bg-black rounded">
			
				<span className={`transition-all ease-in-out ${(show) ? "rotate-45 relative top-2" : ""} w-full h-[2px] bg-white my-1`}></span>
				<span className={`transition-all ease-in-out ${(show) ? "-translate-x-[100px] opacity-0" : ""} w-full h-[2px] bg-white my-1`}></span>
				<span className={`transition-all ease-in-out ${(show) ? "-rotate-45 relative bottom-3" : ""} w-full h-[2px] bg-white my-1`}></span>
			
			</div>

		</nav>
	)
};

export default NavBar;
