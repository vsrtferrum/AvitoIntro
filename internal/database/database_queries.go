package database

var (
	listOfBuyedItems string = `SELECT shop.name
		FROM sales
		JOIN shop ON sales.item_id = shop.id
		WHERE id = $1`
	getUserBalanceById string = `SELECT balance
		FROM users
		WHERE id = $1`
	getUserBalanceByName string = `SELECT balance
		FROM users
		WHERE name = $1`
	getItemCost string = `SELECT cost
		FROM shop
		WHERE name = $1`
	sendedMoneyStat string = `SELECT users.name, users.id, transations.cost
		FROM users
		JOIN transations ON transations.sender_id = users.id
		WHERE id = $1`
	recievedMoneyStat string = `SELECT users.name, users.id, transations.cost
		FROM users+
		JOIN transations ON transations.recipient_id = users.id
		WHERE id = $1`
	authUser string = `SELECT id, password
		FROM users
		WHERE name = $1 AND password = $2`
	updateBalanceById string = `UPDATE users
		SET balance = $1 
		WHERE id = $2`
	updateBalanceByName string = `UPDATE users
		SET balance = $1 
		WHERE name = $2`
	insertSale string = `INSERT INTO sales
		(customer_id, item_id, cost) 
		$1, $2, $3`
	insertTransfer string = `INSERT INTO transfers
		(sender_id, recipient_id, cost) 
		$1, $2, $3`
)
