# What is this

Draft SQL queries that are not used yet

## user

### `/user/{id}/books`

`
SELECT *
FROM books
JOIN favorite_books
ON books.id = favorite_books.book_id
WHERE favorite_books.user_id = ?
`

### `/user/{id}/authors`

`
SELECT *
FROM authors
JOIN favorite_authors
ON authors.id = favorite_authors.author_id
WHERE favorite_authors.user_id = ?
`

### `/user/{id}/reviews`

`
SELECT *
FROM book_reviews
JOIN books
ON book_reviews.book_id = books.id
WHERE book_reviews.user_id = ?
`

## book

### `/book/{id}/users`

`
SELECT *
FROM users
JOIN favorite_books
ON favorite_books.user_id = users.id
WHERE favorite_books.book_id = ?
`

### `/book/{id}/authors`

`
SELECT *
FROM authors
JOIN favorite_authors
ON favorite_authors.author_id = authors.id
WHERE favorite_authors.book_id = ?
`

### `/book/{id}/reviews`

`
SELECT *
FROM book_reviews
JOIN users
ON books_reviews.user_id = users.id
WHERE book_reviews.book_id = ?
`

## author

### `/author/{id}/users`

`
SELECT *
FROM users
JOIN favorite_authors
ON favorite_authors.user_id = users.id
WHERE favorite_authors.authod_id = ?
`

### `/author/{id}/books`

`
SELECT *
FROM books
JOIN authorship
ON authorship.book_id = books.id
WHERE authorship.author_id = ?
`
