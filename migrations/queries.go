package migrations

const (
	createTableProducts = "CREATE TABLE IF NOT EXISTS `products` (" +
		"  `id` int(11) NOT NULL AUTO_INCREMENT," +
		"  `name` varchar(255) NOT NULL," +
		"  `brand_name` varchar(255) NOT NULL," +
		"  `details` varchar(255) NOT NULL," +
		"  `image_url` varchar(255) NOT NULL," +
		"  PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8"

	dropTableProducts = "DROP TABLE products"

	createTableVariants = "CREATE TABLE IF NOT EXISTS `varaints` (" +
		"  `id` int(11) NOT NULL AUTO_INCREMENT," +
		"  `product_id` int(11) NOT NULL," +
		"  `variant_name` varchar(255) NOT NULL," +
		"  `variant_details` varchar(255) NOT NULL," +
		"  PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8"

	dropTableVariants = "DROP TABLE products"
)
