import { createSlice, configureStore } from '@reduxjs/toolkit';

import { 
	SetAuthCookie, 
	RemoveJWT
} from '../clientOperations/cookies';

const UserSlice = createSlice({
	
	name: 'Auth',
	initialState: {
		User: false
	},
	
	reducers: {

		login: (state, action) => {
			// Gets the token using the payload.
			// state.User = action.payload;

			if(action.payload.jwt) {
				console.log("Tokenization: ")
				console.log(action.payload.jwt);
				SetAuthCookie(action.payload.jwt);
			}

			state.User = action.payload;
		},
		
		updateUser: (state, action) => {
			// doc: update a field.
			Object.keys(state.User).map(k => {
				if(k in action.payload) {
					state.User[k] = action.payload[k];
				}
			})
		},

		logout: state => {
			state.User = null;
			RemoveJWT();
		},
	}
});

export const { 
	logout, 
	login,
	updateUser
} = UserSlice.actions;

export default configureStore({
	reducer: UserSlice.reducer
});
