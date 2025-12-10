#!/bin/bash

BASE_URL="http://localhost:8080"
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzY1MzM3NzU3LCJpYXQiOjE3NjUyNTEzNTd9.NO80Mwdv1OvrCw0Z7jiSTx6xJvhnKBhCioGwirOGOeM"

echo "=== 1. Login ==="
RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice"}')

echo $RESPONSE | jq '.'

# Extract token
TOKEN=$(echo $RESPONSE | jq -r '.token')
echo "Token: $TOKEN"

echo -e "\n=== 2. Get All Topics ==="
curl -s $BASE_URL/api/topics | jq '.'

echo -e "\n=== 3. Create Topic ==="
curl -s -X POST $BASE_URL/api/topics \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Test Topic",
    "description": "This is a test"
  }' | jq '.'

echo -e "\n=== 4. Create Post ==="
curl -s -X POST $BASE_URL/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "topic_id": 1,
    "title": "Test Post",
    "content": "This is test content"
  }' | jq '.'

echo -e "\n=== 5. Create Comment ==="
curl -s -X POST $BASE_URL/api/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "post_id": 1,
    "content": "Great post!"
  }' | jq '.'

echo -e "\n=== 6. Get Comments ==="
curl -s $BASE_URL/api/posts/1/comments | jq '.'