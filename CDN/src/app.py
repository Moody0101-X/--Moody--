from socket import gethostbyname, gethostname
from flask import Flask, request, send_file

from database import (
    GetProductImageFromCDN,
	SaveProductImageToCDN,
	makeResponse
)

from json import loads

app  =  Flask(__name__)

PORT =  8500
HOST =  gethostbyname(gethostname())

@app.route("/products/img/add", methods=["POST"])
def PostUserAvatar():
    
    data = request.json
    if "id" in data and "mime" in data:
        result = SaveProductImageToCDN(data)
        return result

    return makeResponse(400, "could not find id or mime in request form data! please recheck")

@app.route("/products/img/<fileName>", methods=["GET"])
def GetProductImage(fileName):
    
    result = GetProductImageFromCDN(fileName)
    if result:
        file, ext = result
        return send_file(file, mimetype=f'image/{ext}')

    return makeResponse(500, "server could not find config file.")

if __name__ == '__main__':
    app.run(host="localhost", port=PORT, debug=True)



















