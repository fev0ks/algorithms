select users.id, users.name
from
    users
        right join (select user_id from orders group by user_id) orders
                   on orders.user_id = users.id
order by lower(users.name), users.id
;