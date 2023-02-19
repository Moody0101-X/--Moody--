const SPORT = '8888';
var API;

if(location.port === SPORT) {
	API = `/v2` // Production
} else {
	API = `http://${location.host}:${SPORT}/v2`; // DEV
}

export {
	API
};