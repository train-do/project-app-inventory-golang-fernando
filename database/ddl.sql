create table if not exists "Goods" (
    "id" serial primary key,
    "name" varchar not null,
    "stock" int not null,
    "created_at" timestamp not null);
    
create table if not exists "Category" (
    "id" serial primary key,
    "name" varchar not null,
    "created_at" timestamp not null);
    
create table if not exists "Location" (
    "id" serial primary key,
    "warehouse" varchar not null,
    "rack" varchar not null,
    "created_at" timestamp not null);
    
create table if not exists "Conjuction" (
    "id" serial primary key,
    "good_id" int references "Goods"("id") not null,
    "category_id" int references "Category"("id") not null,
    "location_id" int references "Location"("id") not null);
    
create table if not exists "Log" (
    "id" serial primary key,
    "good_id" int  references "Goods"("id") not null,
    "information" varchar not null,
    "qty" int not null,
    "created_at" timestamp not null);

