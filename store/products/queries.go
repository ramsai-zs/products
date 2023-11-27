package products

const (
	create  = "INSERT into products(name,brand_name,details,image_url) VALUES(?,?,?,?)"
	getById = "SELECT id,name,brand_name,details,image_url FROM products WHERE id = ?"
	get     = "SELECT id,name,brand_name,details,image_url FROM products"
)
