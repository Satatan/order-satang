
CREATE TABLE public.orders (
	id bigserial NOT NULL,
	customer_id varchar(100) NOT NULL,
	product varchar(200) NOT NULL,
	price int4 NOT NULL,
	quantity int4 NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);