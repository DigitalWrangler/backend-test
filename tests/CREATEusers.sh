❯ curl -X POST "http://localhost:8888/users" \
        -H "Content-Type: application/json" \
        -d '{
      "name": "Vincent Wigardt",
      "email": "vincent.wigardt@example.com",
      "birth_date": "1992-12-06",
      "city": "New York"
    }'
