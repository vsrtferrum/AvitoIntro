package database

var (
	listOfBuyedItems string = `SELECT shop.name, COUNT(sales.item_id) AS item_count
		FROM sales
		JOIN shop ON sales.item_id = shop.id
		WHERE sales.customer_id = $1
		GROUP BY shop.name;`
	getUserBalanceById string = `SELECT balance
		FROM users
		WHERE id = $1;`
	getUserBalanceByName string = `SELECT balance
		FROM users
		WHERE name = $1;`
	getItemCost string = `SELECT cost
		FROM shop
		WHERE name = $1;`
	getIdByName string = `SELECT id
		FROM users
		WHERE name = $1;`
	sendedMoneyStat string = `
		SELECT 
			sender.name AS sender_name,  
			recipient.name AS recipient_name, 
			transfers.cost AS cost
		FROM transfers
		JOIN users AS sender ON transfers.sender_id = sender.id
		JOIN users AS recipient ON transfers.recipient_id = recipient.id
		WHERE sender.id = $1;`
	recievedMoneyStat string = `
		SELECT 
			sender.name AS sender_name,  
			recipient.name AS recipient_name, 
			transfers.cost AS cost
		FROM transfers
		JOIN users AS sender ON transfers.sender_id = sender.id
		JOIN users AS recipient ON transfers.recipient_id = recipient.id
		WHERE recipient.id = $1;`
	authUser string = `SELECT id, password
		FROM users
		WHERE name = $1 AND password = $2;`
	updateBalanceById string = `UPDATE users
		SET balance = $1 
		WHERE id = $2;`
	insertSale string = `INSERT INTO sales
		(customer_id, item_id, cost)
		VALUES ($1, $2, $3);`
	insertTransfer string = `INSERT INTO transfers
		(sender_id, recipient_id, cost) 
		VALUES ($1, $2, $3);`
	insertUser string = `INSERT INTO users
		(name, password, balance)
		VALUES 
		($1 , $2, 1000);`
)
