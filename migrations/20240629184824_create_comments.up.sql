CREATE TABLE IF NOT EXISTS Comments(
    id serial PRIMARY KEY,
    username varchar(64),
    text varchar(2000),
    postID int,
    commentID int,
    FOREIGN KEY (postID) REFERENCES Posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (commentID) REFERENCES Comments(id) ON DELETE SET NULL ON UPDATE CASCADE
)