# github.com/gofrs/uuid
../../go/pkg/mod/github.com/gofrs/uuid@v4.4.0+incompatible/generator.go:335:33: g.epochFunc().UnixMilli undefined (type time.Time has no field or method UnixMilli)




Audit for the main Forum:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/audit

Optionals

Authentication:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/authentication
Security:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/security
Moderation:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/moderation
Image-Upload:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/image-upload
Advanced-Features:
https://learn.01founders.co/git/root/public/src/branch/master/subjects/forum/advanced-features


Objectives
This project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

SQLite
In order to store the data in your forum (like users, posts, comments, etc.) 
you will use the database library SQLite.

To structure your database and to achieve better performance
https://www.smartdraw.com/entity-relationship-diagram/

- You must use at least one SELECT, one CREATE and one INSERT queries.