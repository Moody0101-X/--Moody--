
import { UIWrapper } from "../Components/Wrapper";
import { useEffect, useState } from "react";
import { ProductCard, FilterProdByQuery, FilterProdByPriceVal, OnMaxPrice, OnMinPrice } from "../Components/Product"
let FuncMap = {
	'Array': (s) => s.map,
	'Map': (s) => s.keys().map,
	'Object': (s) => Object.keys(s).map
}

const Stator = (initial) => {
	
	const [ State, SetState ] = useState(initial);
	const Clear = () => {
		SetState(null);
	}

	const Fill = (FillValue) => {
		
		if((typeof State) === 'object' ) {
			if(type in FuncMap)
				return FuncMap[type](State)(k => FillValue);
		}

		return null;
	}

	const MapOver = (callback, type=null) => {
		if((typeof State) === 'object' ) {
			if(type in FuncMap)
				return FuncMap[type](State)(v => callback(v));
		}

		return null;
	}

	return [State, SetState, Clear]
}


import { Link, useNavigate } from "react-router-dom";

const Home = () => {
	const [query, setQuery] = useState("");
	
	const OnTyping = (e) => {
		setQuery(e.target.value);
		console.log(e.target.value);
	};

	return (
		<UIWrapper Class="h-[90vh]">

			<div className="h-full w-full flex flex-col justify-center items-center">
				<h1 className="my-2 text-[50px] font-semibold"> Find High Quality Product! </h1>
				<p className="text-2xl font-normal text-slate-600"> With the best prices possible. </p>
				<input onChange={OnTyping} type="text" placeholder="Look for Something!" className="w-[500px] outline-none transition-all ease-in-out focus:shadow rounded-md border-2 border-green-500 p-4 bg-slate-50 my-4" />
				<div className="">
					<button className="mx-2 border border-neutral-800 transition-all ease-in-out hover:bg-neutral-900 bg-neutral-800 text-white rounded-md p-4"> Search </button>
					
				<Link className="mx-2 bg-none border text-black hover:text-white hover:bg-green-500 border-green-500 transition-all ease-in-out text-white rounded-md p-4" to="/Store">
					Browse Our Catalogue
				</Link>
				</div>
			</div>

		</UIWrapper>
	)
}

export default Home;