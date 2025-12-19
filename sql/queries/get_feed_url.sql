-- name: GetFeedByURL :one
SELECT * FROM Feeds WHERE url = $1;
