db = db.getSiblingDB('userdb');

db.users.insert({
    "name": "John Doe",
    "email": "john.doe@example.com",
    "birth_date": "2001-05-03",
    "city": "New York",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "birth_date": "2005-10-23",
    "city": "Chicago",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Alice Smith",
    "email": "alice.smith@example.com",
    "birth_date": "1990-03-15",
    "city": "Los Angeles",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Bob Johnson",
    "email": "bob.johnson@example.com",
    "birth_date": "1985-07-20",
    "city": "San Francisco",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Charlie Brown",
    "email": "charlie.brown@example.com",
    "birth_date": "1998-11-12",
    "city": "Boston",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Diana Green",
    "email": "diana.green@example.com",
    "birth_date": "1993-06-25",
    "city": "Seattle",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Edward White",
    "email": "edward.white@example.com",
    "birth_date": "1982-02-10",
    "city": "Miami",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Fay Lee",
    "email": "fay.lee@example.com",
    "birth_date": "1995-08-17",
    "city": "Houston",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "George King",
    "email": "george.king@example.com",
    "birth_date": "2000-01-05",
    "city": "Dallas",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

db.users.insert({
    "name": "Hannah Scott",
    "email": "hannah.scott@example.com",
    "birth_date": "1996-09-13",
    "city": "Denver",
    "active": true,
    "created_at": new Date(),
    "updated_at": new Date()
});

print("Data has been written to the collection");
