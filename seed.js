db = db.getSiblingDB('userdb');

db.users.insert({
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30,
    "city": "New York"
});

db.users.insert({
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "age": 25,
    "city": "Chicago"
});

print("Data has been written to the collection");
