/* Write your PL/SQL query statement below */
SELECT name as Customers
FROM Customers
WHERE id NOT IN (
    SELECT DISTINCT customerId
    FROM Orders
)
;
