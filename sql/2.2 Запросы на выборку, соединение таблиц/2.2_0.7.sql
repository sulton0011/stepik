SELECT book.title as Название, name_author as Автор, sum(book.amount + supply.amount) as Количество
FROM 
    author 
    INNER JOIN book USING (author_id)   
    INNER JOIN supply ON book.title = supply.title 
                         and author.name_author = supply.author
    where book.price = supply.price
    group by book.title, name_author;