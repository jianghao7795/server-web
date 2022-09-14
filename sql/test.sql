SELECT count(*) FROM `comment` WHERE MATCH(comment.content) AGAINST('大') AND `comment`.`deleted_at` IS NULL

-- alter table comment add fulltext index content(`content`);