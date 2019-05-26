# Supertest

`npm install -D supertest @types/supertest`

## 예제

```ts
import * as request from "supertest";

describe("GET /users", function() {
  it("responds with json", function() {
    return request(app)
      .get("/users")
      .set("Accept", "application/json")
      .expect("Content-Type", /json/)
      .expect(200)
      .then(response => {
        assert(response.body.email, "foo@bar.com");
      });
  });
});

describe("POST /user", function() {
  it('user.name should be an case-insensitive match for "john"', function(done) {
    request(app)
      .post("/user")
      .send("name=john") // x-www-form-urlencoded upload
      .set("Accept", "application/json")
      .expect(function(res) {
        res.body.id = "some fixed id";
        res.body.name = res.body.name.toLowerCase();
      })
      .expect(
        200,
        {
          id: "some fixed id",
          name: "john"
        },
        done
      );
  });
});
```
