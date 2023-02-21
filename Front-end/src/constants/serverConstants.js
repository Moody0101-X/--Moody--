const SPORT = '8888';
var API;

if(location.port === SPORT) {
	API = `/v2` // Production
} else {
	API = `http://${location.hostname}:${SPORT}`; // DEV
}

export {
	API
};