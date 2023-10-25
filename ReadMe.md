## About


**We have 5 main layer:**

1) Deliver layer
2) Interactors (Use Cases)
3) Agify (Enrich)
4) Repositories
5) Domain

**1** -> **2** <-> **3** -> **4** -> **5**

## this is example params for fast check

Insert 

http://your_port/insert

{
"name":"Dmitry",
"surname":"Krapysh",
"patronymic":"Andreivich"
}

Update

http://your_port/update

{
"id" : 1,
"name":"Bob",
"surname":"Karapysh",
"patronymic":"Andreivich",
"age":26,
"sex":"male",
"nation":"rus"
}

Delete

http://your_port/delete

{
"id" : 1
}

Get

http://your_port/get

GetByID

http://your_port/getId?id=1

GetByName

http://your_port/getName?name=bob

GetByLimit

http://your_port/getLimit?limit=3

GetByOffset

http://your_port/getOffset?offset=2

GetByLimitAndOffset

http://your_port/getLimitOffset?limit=4&offset=2

