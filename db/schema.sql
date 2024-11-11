create table if not exists accounts (
    id serial primary key,
    name text not null unique,
    notes text default '',
    align decimal(12, 2) default 0
);

create table if not exists expense_accounts (
    id serial primary key unique,
    name text not null,
    notes text default ''
);

create table if not exists income_accounts (
    id serial primary key,
    name text not null unique,
    notes text default ''
);

create table if not exists expense_categories (
    id serial primary key,
    name text not null unique,
    notes text default ''
);

create table if not exists income_categories (
    id serial primary key,
    name text not null unique,
    notes text default ''
);

create table if not exists transfer_categories (
    id serial primary key,
    name text not null unique,
    notes text default ''
);

create table if not exists expenses (
    id serial primary key,
    amount decimal(12, 2) not null,
    description text default '',
    date date not null,
    source serial,
    destination serial,
    category serial,
    constraint fk_source foreign key (source) references accounts(id),
    constraint fk_destination foreign key (destination) references expense_accounts(id),
    constraint fk_category foreign key (category) references expense_categories(id)
);

create table if not exists incomes (
    id serial primary key,
    amount decimal(12, 2) not null,
    description text default '',
    date date not null,
    source serial,
    destination serial,
    category serial,
    constraint fk_source foreign key (source) references income_accounts(id),
    constraint fk_destination foreign key (destination) references accounts(id),
    constraint fk_category foreign key (category) references income_categories(id)
);

create table if not exists transfers (
    id serial primary key,
    amount decimal(12, 2) not null,
    description text default '',
    date date not null,
    source serial,
    destination serial,
    category serial,
    constraint fk_source foreign key (source) references accounts(id),
    constraint fk_destination foreign key (destination) references accounts(id),
    constraint fk_category foreign key (category) references transfer_categories(id)
);

create table if not exists budgets (
    id serial primary key,
    amount decimal(12, 2) not null,
    category serial,
    constraint fk_category foreign key (category) references expense_categories(id)
);
