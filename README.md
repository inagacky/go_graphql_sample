# GO×graphQLのサンプルコード


## DB
Dockerコンテナで起動します。
/bin/sh run.sh

##API実行
#### ユーザー作成
```
% curl -X POST -d 'mutation { createUser(firstName: "test", lastName: "hoge" email: "test@test.co.jp" ){ id, firstName, lastName, email } }' http://localhost:8080/graphql | jq .
{
  "data": {
    "createUser": {
      "email": "test@test.co.jp",
      "firstName": "test",
      "id": "5",
      "lastName": "hoge"
    }
  }
}
```

#### ユーザー取得
```
% curl -X POST -d 'query { user(id:1) {id, firstName, lastName, email} }' http://localhost:8080/graphql | jq .
{
  "data": {
    "user": {
      "email": "test@test.co.jp",
      "firstName": "test",
      "id": "1",
      "lastName": "hoge"
    }
  }
}
```


#### ユーザー一覧取得
```
% curl -X POST -d 'query { userList{id, firstName, lastName, email} }' http://localhost:8080/graphql | jq .
{
  "data": {
    "userList": [
      {
        "email": "test@test.co.jp",
        "firstName": "test",
        "id": "1",
        "lastName": "hoge"
      },
      {
        "email": "test@test.co.jp",
        "firstName": "test",
        "id": "2",
        "lastName": "hoge"
      },
      {
        "email": "test@test.co.jp",
        "firstName": "test",
        "id": "3",
        "lastName": "hoge"
      },
      {
        "email": "test@test.co.jp",
        "firstName": "test",
        "id": "4",
        "lastName": "hoge"
      },
      {
        "email": "test@test.co.jp",
        "firstName": "test",
        "id": "5",
        "lastName": "hoge"
      }
    ]
  }
}
```

#### タスク作成

```
% curl -X POST -d 'mutation { createTask(userId: 1, title: "testTitle" content: "testContent" ){ id, userId, title, content } }' http://localhost:8080/graphql | jq .
{
  "data": {
    "createTask": {
      "content": "testContent",
      "id": 7,
      "title": "testTitle",
      "userId": 1
    }
  }
}

```

#### タスク取得
```
% curl -X POST -d 'query { task(id:7) {id, userId, title, content} }' http://localhost:8080/graphql | jq .
{
  "data": {
    "task": {
      "content": "testContent",
      "id": 7,
      "title": "testTitle",
      "userId": 1
    }
  }
}
```
#### タスク一覧取得
```
% curl -X POST -d 'query { taskList{ id, userId, title, content } }' http://localhost:8080/graphql | jq .
{
  "data": {
    "taskList": [
      {
        "content": "testContent",
        "id": 2,
        "title": "testTitle",
        "userId": 1
      },
      ...
      {
        "content": "testContent",
        "id": 7,
        "title": "testTitle",
        "userId": 1
      }
    ]
  }
}

```

#### タスクとユーザーを一緒に取得
```
% curl -X POST -d 'query { user(id:1) {id, firstName, lastName, email}, task(id:2) {id, userId, title, content} }' http://localhost:8080/graphql | jq .
{
  "data": {
    "task": {
      "content": "testContent",
      "id": 2,
      "title": "testTitle",
      "userId": 1
    },
    "user": {
      "email": "test@test.co.jp",
      "firstName": "test",
      "id": "1",
      "lastName": "hoge"
    }
  }
}

```
