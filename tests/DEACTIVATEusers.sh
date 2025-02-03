curl -X PUT "http://localhost:8888/users/679e49821a094a80b3e9496d/deactivate" \
        -H "Content-Type: application/json" \
        -d '{
      "active": false
    }'
