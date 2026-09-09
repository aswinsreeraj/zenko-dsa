-- Write your PostgreSQL query statement below
select distinct p.email from Person p  where p.email in (
    select p1.email from Person p1 where p1.email = p.email and p.id <> p1.id
);