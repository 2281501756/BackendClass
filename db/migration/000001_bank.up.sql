create table "user"
(
    id    integer                                   not null
        constraint user_pk
            primary key,
    name  varchar(10) default ''::character varying not null,
    email varchar(20) default ''::character varying not null,
    money integer     default 0                     not null
);

alter table "user"
    owner to root;

create table bill
(
    id      integer not null
        primary key,
    user_id integer
        constraint bill_user_id_fk
            references "user",
    money   integer not null
);

alter table bill
    owner to root;

