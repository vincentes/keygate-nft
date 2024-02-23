-- +goose Up
-- +goose StatementBegin
-- UserKey
create table UserKey (
    ID varchar(255) primary key,
    UserID varchar(255) not null,
    KeyID varchar(255) not null,
    foreign key (UserID) references User(ID) on delete cascade,
    foreign key (KeyID) references `Key`(ID) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table UserKey;
-- +goose StatementEnd
