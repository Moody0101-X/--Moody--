import "./Loader.css";
import { useState, useEffect } from "react";
import { faInfoCircle, faWarning, faCircleCheck, faEdit } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon as Fa } from '@fortawesome/react-fontawesome';

const ApplicationMessage = ({ msg, Case, id }) => {
	const [show, setShow] = useState(false);

	const Map_ = {
		info: {
			T: "text-slate-900 bg-white",
			icon: faInfoCircle
		},
		
		error: {
			T: "text-white bg-rose-500",
			icon: faWarning
		},

		success: {
			T: "text-white bg-green-500",
			icon: faCircleCheck
		}
	}

	if(!(Case in Map_)) Case = "error"
	// For developement.
	
	useEffect(() => {
		
		setShow(true); 
		
		const TimeOut = setTimeout(() => setShow(false), 5 * 1000);

		return () => {
			setShow(false)
			clearTimeout(TimeOut);
		};

	}, [id])

	return (
			((Case in Map_) && msg) ? (
			
				<div className={`${show ? "translate-y-2 opacity-100" : "-translate-y-16 opacity-0"} z-50 flex flex-row items-center justify-center shadow visible fixed top-1 p-3 rounded transition-all -translate-x-1/2 left-1/2 ${Map_[Case].T}`}> 
					<span className="mx-2"> { msg } </span>
					<Fa icon={Map_[Case].icon} />
				</div>

			) : ""
	)
}

const Loader = ({
	color="white",
	size=20,
	Class=""
}) => (<div id="circle2" style={{
	width: `${size}px`,
	height: `${size}px`
}} className={`show border-t-${color} ${Class}`}></div>);


export { ApplicationMessage, Loader };