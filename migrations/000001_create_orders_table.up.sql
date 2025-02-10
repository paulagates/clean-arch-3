create table orders (
   id          varchar(255) not null,
   price       float not null,
   tax         float not null,
   final_price float not null,
   primary key ( id )
);