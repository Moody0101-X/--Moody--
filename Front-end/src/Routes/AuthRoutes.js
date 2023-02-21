import { AuthUserWithCredentials, SignUpUser } from "../Server/Auth";
import { useState, useEffect } from "react";
import { UIWrapper } from "../Components/Wrapper";
import { Loader } from "../Components/UtilityComponents";
import { useSelector, useDispatch} from "react-redux";
import { login } from "../Store/UserStore";

const isAnyEmpty = (state) => {
	
	var Empty = [];	
	
	Object.keys(state).map((k, index) => {
		if(!state[k]) {
			Empty.push(index);
		}
	});

	return Empty;
}


const GenericInputField = ({ OnTyping = (e) => console.log(e.target.value), ...Rest }) => {
	return (
		<input {...Rest} onChange={OnTyping} className="w-[500px] outline-none transition-all ease-in-out focus:shadow rounded-md border-2 border-green-500 p-4 bg-slate-50 my-4" />
	)
}

const UIButton = ({ OnClick, Class="", Label }) => {
	const [Loading, SetLoading] = useState(false);
	
	const onClickFactory = (e) => {
		e.preventDefault();
		SetLoading(true);
		OnClick(() => SetLoading(false));
	}

	return (
		<button onClick={onClickFactory} className={"flex flex-row justify-center items-center order border-neutral-800 transition-all ease-in-out hover:bg-neutral-900 bg-neutral-800 text-white rounded-md py-2 px-4 my-4 cursor-pointer " + Class}> 
			<p> { Label } </p> 
			{ (Loading) ? <Loader Class="ml-4" size={20} /> : ""}			
		</button>
	)
}

const SignUpRoute = ({ NotificationFunc }) => {

	// SignUpUser(email, pwd, name, phone_number);
	var Validators;
	const dispatch = useDispatch();

	const [FormState, setFormState] = useState({
		email: "", 
		pwd: "", 
		name: "", 
		phone_number: ""
	});

	useEffect(() => {	
		Validators = document.getElementsByClassName("vald");
	}, [])

	const OnTypingFactory = (k) => {
		
		return (e) => {
			var tmp = FormState;
			tmp[k] = e.target.value;
			setFormState(tmp);
		}
	}

	const SendDataToServer = () => {
		var EmptyFieldIndeces = isAnyEmpty(FormState);
		const keys = Object.keys(FormState);
		
		if(EmptyFieldIndeces.length > 0) {
			EmptyFieldIndeces.map(i => {
				if(Validators[i] !== undefined)
					Validators[i].textContent = "This field should not be empty. *";
			})
			
			return
		}


		if(FormState.name.length > 30) {
			Validators[2].textContent = "This name is invalid. please enter a shorter name.";
			return
		}


		if(FormState.pwd.length < 8) {
			Validators[1].textContent = "password Needs to be longer than 8 characters";
			return
		}

		if(FormState.phone_number.length < 10 || FormState.phone_number.length > 25) {
			Validators[3].textContent = "Invalid phone number.";
			return
		}		
		
		SignUpUser(FormState.email, FormState.pwd, FormState.name, FormState.phone_number)
		.then(r => r.json())

		.then(jsonData => {
			if(jsonData.code === 200) {
				const User = jsonData.data;
				dispatch(login(User));
				return
			}

			NotificationFunc(jsonData.data, "error");
		})
		.catch(e => {
			console.log(e);
		})
	}

	const Submit = (CloseLoader) => {	
		SendDataToServer();		
		CloseLoader(p => false);
	}
	
	return (
		<UIWrapper Class="w-[90%] h-[90vh]">
			<h1> Sign up for an account </h1>
			<form className="flex flex-col items-start">			
				
				<GenericInputField OnTyping={OnTypingFactory("email")} placeholder="Email" type="Email"/>
				<span className="vald text-sm text-red-600 px-1"> </span>
				<GenericInputField OnTyping={OnTypingFactory("pwd")} placeholder="password" type="password"/>
				<span className="vald text-sm text-red-600 px-1"> </span>
				<GenericInputField OnTyping={OnTypingFactory("name")} placeholder="Name" type="text"/>
				<span className="vald text-sm text-red-600 px-1"> </span>
				<GenericInputField OnTyping={OnTypingFactory("phone_number")} placeholder="Phone Number" type="text"/>
				<span className="vald text-sm text-red-600 px-1"> </span>	
				
				<UIButton Label="submit" OnClick={ Submit } Class="my-2"/>
			</form>
		</UIWrapper>
	)

}

const LoginRoute = ({ NotificationFunc }) => {
	// AuthUserWithCredentials(email, pwd);
	const [FormState, setFormState] = useState({
		email: "",
		pwd: ""
	});

	const dispatch = useDispatch();

	var Validators;
	
	useEffect(() => {
		Validators = document.getElementsByClassName("vald");

	}, [])

	const OnTypingFactory = (k) => {
		
		return (e) => {
			var tmp = FormState;
			tmp[k] = e.target.value;
			setFormState(tmp);
			console.log(tmp);
		}
	}

	const SendDataToServer = () => {
		var EmptyFieldIndeces = isAnyEmpty(FormState);
		var keys = Object.keys(FormState);

		if(EmptyFieldIndeces.length > 0) {
			EmptyFieldIndeces.map(i => {
				if(Validators[i] !== undefined) {
					Validators[i].textContent = "This field should not be empty. *";
				}
			})

			return
		}

		AuthUserWithCredentials(FormState.email, FormState.pwd)
		.then(response =>  response.json())
		.then(jsonData => {	
			console.log(jsonData);
			if(jsonData.code === 200) {
				const User = jsonData.data;
				dispatch(login(User));
				return
			}

			NotificationFunc(jsonData.data, "error");
		})
		.catch(e => {
			console.log(e);
		})

	}

	const Submit = (CloseLoader) => {
		SendDataToServer();
		CloseLoader();
	}
	
	return (
		<UIWrapper Class="p-4 h-[90vh]">
			
			<h1> Login route. </h1>

			<form className="flex flex-col items-start">
				<GenericInputField OnTyping={OnTypingFactory("email")} placeholder="Email" type="Email"/>
				<span className="vald text-sm text-red-600 px-1"> </span>	
				<GenericInputField OnTyping={OnTypingFactory("pwd")} placeholder="password" type="password"/>
				<span className="vald text-sm text-red-600 px-1"> </span>	
				<UIButton Label="submit" OnClick={Submit} Class="" />			
			</form>

		</UIWrapper>
	)
}

export {
	SignUpRoute,
	LoginRoute
};