// 0607343642 => MohcinSnack.
import {useState, useEffect} from "react";
import { ProductCard, FilterProdByQuery, FilterProdByPriceVal, OnMaxPrice, OnMinPrice } from "../Components/Product"

let letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const generateLabel = (length) => {
	var r = '';
	
	for (var i = 0; i < length; i++) {
		var index = Math.floor(Math.random() * letters.length);
		r += letters[index]
	}

	return r;
}

const Latest = () => {
	const [products, setProducts] = useState([]);

	const GetProducts = (n) => {
		
		let tmp = [];
		

		for (var i = 0; i < n; i++) {
			tmp.push(
				new Map([
					["price", Math.floor(Math.random() * 1000)],
					["label", generateLabel(7)],
					["desc", "Something to be sold!"]
				])
			)
		}

		setProducts(tmp);
		tmp = [];
	}

	useEffect(() => {
		
		GetProducts(20)
		// const t = setTimeout(() => GetProducts(10), 1000 * 5);
		
		return () => {
			setProducts([]);
			// clearTimeout(t);
		}

	}, []);

	return (
		<>
			<div className="flex items-center justify-between w-[90%] text-lg font-semibold my-2 p-2 border-b border-b-slate-100 text-slate-800"> 
				<h1>
					Latest products
				</h1>
			</div>
			<ul className="w-full flex items-center justify-center flex-row flex-wrap">
				{ (products.length > 0) ? (products.map(v => <ProductCard label={v.get("label")} price={v.get("price")} desc={v.get("desc")}/>)) : (<h1> Loading... </h1>) }
			</ul>
		</>
	)
}

const Store = () => <Latest />;
export default Store;