const FilterProduct = (setCallBack, predicate, all) => {
	var tmp = all;
	
	if(tmp.length > 0) {
		tmp = tmp.filter(v => (predicate(v)))
	}

	setCallBack(tmp);
}

const FilterProdByQuery 	= 	(set, Q, a)      =>  FilterProduct(set, v => (v.get("label").includes(Q)), a);
const FilterProdByPriceVal 	= 	(set, p, a)      =>  FilterProduct(set, v => (v.get("price") === p), a);
const OnMaxPrice 			= 	(set, Max, a)	 =>  FilterProduct(set, v => (v.get("price") < Max), a);
const OnMinPrice 			=   (set, Min, a)    =>  FilterProduct(set, v => (v.get("price") > Min), a);

const ProductCard = ({ label, price, desc, ...RestProps }) => {

	return (
		<li { ...RestProps } className="m-2 flex flex-col justify-center items-start bg-slate-100 transition-all ease-in-out rounded hover:shadow-xl">										
			
			<div className="w-full bg-blue-400 rounded-t h-32" style={{
				backgroundImage: `url(https://picsum.photos/id/${price}/480/300)`,
				backgroundPosition: "center",
				backgroundSize: "100%",
				backgroundRepeat: "no-repeat"
			}}>

			</div>

			<h1 className="hover:underline text-xl mx-2 py-2 px-2 text-neutral-900 font-semibold">
				{label}
			</h1>

			<p className="mx-2 py-2 px-2 text-neutral-600 text-lg">
				{price}$ <span className="mx-1 bg-green-500 rounded-md text-white text-sm px-1"> USD </span>
			</p>

			<p className="mx-2 py-2 px-2 text-neutral-500 text-lg">
				{desc}
			</p>

			<div className="p-2">
				<button className="rounded-md m-2 py-1 px-4 text-lg text-white hover:bg-blue-900 transition-all ease-in-out bg-blue-800">
					Buy
				</button>

				<button className="rounded-md m-2 py-1 px-5 text-lg text-neutral-800 hover:bg-slate-300 transition-all ease-in-out bg-slate-200">
					Learn More!
				</button>
			</div>
		</li>
	)
}

export { ProductCard, FilterProdByQuery, FilterProdByPriceVal, OnMaxPrice, OnMinPrice };