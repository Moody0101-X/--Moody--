const SetAuthCookie = (T, ExpirationDays = 365) => {

	if(T) {
		let Date_ = new Date();
		Date_.setTime(Date_.getTime() + ExpirationDays * 24 * 60 * 60 * 1000);
		const expires = `; expires=${Date_.toUTCString()}`;
    	document.cookie = `AccessToken=${ T }${ expires }; path=/;`;
    	console.log(`AccessToken=${ T }${ expires }; path=/;`)
    	return document.cookie;
	} 

	alert("Token was expected from Function `SetAuthCookie`");
};


const getJwtAuthToken = (key="AccessToken") => {
    
    const cookieMap = new Map(document.cookie.split(';').map(v => (v.trim().split("="))));
	const value = cookieMap.get(key);
    
    if(value === undefined) 
    	return "";
    
    return value;
}

function RemoveJWT() {   
    document.cookie = 'AccessToken=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

export { 
	SetAuthCookie, 
	getJwtAuthToken, 
	RemoveJWT
};