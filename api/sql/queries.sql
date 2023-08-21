-- name: InsertNewProductMeta :exec
INSERT INTO product_meta (id, "name", symbol, "description", "type", exchange, "location") 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: InsertNewProductPlatformMeta :exec
INSERT INTO product_platform (product_id, platform_name, identifier)
VALUES ($1, $2, $3);

-- name: GetProductMeta :one
SELECT id, "name", symbol, "description", "type", exchange,  "location"  FROM product_meta WHERE id = ($1);

-- name: GetProductMetaWithPlatform :one
SELECT product_meta.id, "name", symbol, "description", "type", exchange,  "location" , platform_name, identifier 
FROM product_meta 
JOIN product_platform 
ON product_meta.id = product_platform.product_id 
WHERE product_meta.id = ($1);

-- name: UpdatePlatformIdentifier :exec
UPDATE product_platform SET identifier = ($1) WHERE product_id = ($2) AND platform_name = ($3);

-- name: DeletePlatformInfo :exec
DELETE FROM product_platform WHERE product_id = ($1) AND platform_name = ($2);

-- name: InsertPlatformInfo :exec
INSERT INTO product_platform (product_id, platform_name, identifier) VALUES ($1, $2, $3);

-- name: GetProductIdBySymbol :one
SELECT id FROM product_meta WHERE symbol = ($1);

-- name: GetStoredProductList :many
SELECT id FROM product_meta,store_log WHERE product_meta.id = store_log.product_id;