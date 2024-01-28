-- Write your PostgreSQL query statement below
select e.name
from employee e
inner join (
    select managerId, count(*)
    from employee
    group by managerId
    having count(*) >= 5
) temp
on e.id = temp.managerId
;
