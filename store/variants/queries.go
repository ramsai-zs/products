package variants

const (
	create              = "INSERT INTO varaints(product_id,variant_name,variant_details) VALUES(?,?,?)"
	getById             = "SELECT id,product_id,variant_name,variant_details FROM varaints where id=?"
	getByProductId      = "SELECT id,variant_name,variant_details FROM varaints where product_id=?"
	getByIdAndProductId = "SELECT id,product_id,variant_name,variant_details FROM varaints where id=? AND product_id=?"
)
