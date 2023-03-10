import { Children } from "react";

const UIWrapper = ({ children, Class="", ...Rest }) => {
	return (
		<div className={`flex w-[90%] flex-col justify-center items-center ${Class}`} {...Rest} >
			{ Children.map(children, child => child) }
		</div>
	)
}

export { UIWrapper };