// 0607343642 => MohcinSnack.
import {useState, useEffect} from "react";
import { ProductsSection, ProductCard, FilterProdByQuery, FilterProdByPriceVal, OnMaxPrice, OnMinPrice } from "../Components/Product"

let letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const generateLabel = (length) => {
	var r = '';
	
	for (var i = 0; i < length; i++) {
		var index = Math.floor(Math.random() * letters.length);
		r += letters[index]
	}

	return r;
}

const LatestProducts = () => {
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

	return <ProductsSection Products={products} Title="Latest shit" />
}

const Store = () => <LatestProducts />;
export default Store;