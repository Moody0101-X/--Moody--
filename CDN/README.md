# DOCS

## Description
- This app is a cdn for my E-commerce app. it support caching and loading image via links.

## Implemented Routes

- `/products/img/add`
	
	Method: POST
	Expected data: { "mime": ImageMime, "id": ProductID}

- `/products/img/<fileName>` 

	Method: GET
	Expected data: fileName, encoded into the url.

HAPPY HACKING!

## None-Implemented Routes

- `/CDNStats?Token={AdminSecret}` an admin route to get information about the cdn.