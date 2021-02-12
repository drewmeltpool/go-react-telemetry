
drop table if exists telemetry;
drop table if exists tablet;



create table tablet(
    id int auto_increment,
    name varchar(60) not null,
    primary key(id)
);


create table telemetry(
    id int auto_increment,
    tabletId int not null,
    battery int not null,
    deviceTime varchar(60) not null,
    timeStamp varchar(60) not null,
    currentVideo varchar(60),
    primary key(id),
    foreign key (tabletId) references tablet(id)
);