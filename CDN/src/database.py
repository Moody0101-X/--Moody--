from base64 import b64decode
from pathlib import Path
from json import load, dump 
from random import randint
from hashlib import sha256
from dataclasses import dataclass
from os import environ, path

if "CDN" in environ:
    CDN = environ["CDN"]

    if path.exists(CDN): print(' * DATA CLUSTER -> ', CDN)
    assert CDN, "Could not find the cdn, check if it is assigned in the env vars."

@dataclass
class Response:
    code: int
    data: any

    def __dict__(self) -> dict:
        return {
            "code": self.code,
            "data": self.data
        }
    def make(self) -> dict:
        return self.__dict__()

def makeResponse(code: int = 200, data: any = "No data") -> None: return Response(code, data).make()

NOT_EMP = makeResponse(404, "NotImplemented")

def Unpack(IMime, id_) -> tuple:
    """ Unpacking the mime image. """
    Extention = IMime.split(";")[0].split(":")[1].split("/")[1]
    Bytes = b64decode(IMime.split(";")[1].split(",")[1].encode())
    FileName = f"{id_}___{generateRandomName()}"
    
    return Bytes, f"{FileName}.{Extention}"


def SaveProductImageToCDN(data: dict) -> dict:
    
    MIME, ID = data["mime"], data["id"]
    if isinstance(ID, int): ID = str(ID)
    Upath = Path(CDN) / ID
    
    if not Upath.exists(): Upath.mkdir()
    
    Bytes, FName = Unpack(MIME, ID)
    ImagePath = Upath / FName

    with open(ImagePath, "wb") as fp:
        fp.write(Bytes)

    return makeResponse(200, {
        "url": f"/products/img/{FName}"
    })

def GetProductImageFromCDN(FileName: str) -> tuple[str, str] | bool:    
    imgPath = Path(CDN) / FileName
    
    if imgPath.exists():
        Extention = FileName.split(".")[1] # get img ext.
        return imgPath, Extention

    return False

def generateRandomName():
    RandomBytes = [chr(i) for i in [randint(0, 100) for i in range(32)]]
    return sha256("".join(RandomBytes).encode()).hexdigest()