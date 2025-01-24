db = db.getSiblingDB('userdb');

db.users.insert({
    "name": "John Doe",
    "email": "john.doe@example.com",
    "birth": "2001-05-03",
    "city": "New York"
});

db.users.insert({
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "birth": "2005-10-23",
    "city": "Chicago"
});

print("Data has been written to the collection");
