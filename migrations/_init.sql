
-- Create the todos table
create table todos (
  id serial primary key not null,
  title text not null,
  completed boolean
);
