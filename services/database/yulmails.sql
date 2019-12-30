/*
__   ___   _ _     ____   ___  _     
\ \ / / | | | |   / ___| / _ \| |    
 \ V /| | | | |   \___ \| | | | |    
  | | | |_| | |___ ___) | |_| | |___ 
  |_|  \___/|_____|____/ \__\_\_____|
                                     
*/

CREATE TABLE public."conservation"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	sent int NOT NULL,
	unsent int NOT NULL,
	keep_email_content boolean,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."conservation"
	OWNER to postgres;

CREATE TABLE public."quota"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	duration_minutes int NOT NULL,
	quantity int NOT NULL,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."quota"
	OWNER to postgres;

CREATE TABLE public."abuse"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	name character varying(25) NOT NULL,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."abuse"
	OWNER to postgres;

CREATE TABLE public."entity"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	name character varying(50) NOT NULL,
	description character varying(250),
	conservation_id int NOT NULL REFERENCES "conservation" (id),
	abuse_id int NOT NULL REFERENCES "abuse" (id),
	PRIMARY KEY (id)
)
WITH (
	OIDS = FALSE
);

ALTER TABLE public."entity"
	OWNER to postgres;

CREATE TABLE public."entity_quotas"
(
	entity_id int NOT NULL REFERENCES "entity" (id),
	quota_id int NOT NULL REFERENCES "quota" (id)
)
WITH (
	OIDS=FALSE
);

CREATE TABLE public."environment"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	name character varying(50) NOT NULL,
	entity_id int NOT NULL REFERENCES "entity" (id),
	abuse_id int NOT NULL REFERENCES "abuse" (id),
	conservation_id int NOT NULL REFERENCES "conservation" (id),
	is_open boolean,
	description character varying(250),
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."environment"
	OWNER to postgres;

CREATE TABLE public."environment_quotas"
(
	environment_id int NOT NULL REFERENCES "environment" (id),
	quota_id int NOT NULL REFERENCES "quota" (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."environment_quotas"
	OWNER to postgres;

CREATE TABLE public."domain"
(
	id SERIAL NOT NULL,
	created timestamp without time zone NOT NULL,
	name character varying(25) NOT NULL,
	environment_id int NOT NULL REFERENCES "environment" (id),
	conservation_id int NOT NULL REFERENCES "conservation" (id),
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."domain"
	OWNER to postgres;

CREATE TABLE public."domain_quotas"
(
	domain_id int NOT NULL REFERENCES "domain" (id),
	quota_id int NOT NULL REFERENCES "quota" (id)
)
WITH (
	OIDS=FALSE
);

ALTER TABLE public."domain_quotas"
	OWNER to postgres;

/*
 _ __ ___   ___   ___| | __
| '_ ` _ \ / _ \ / __| |/ /
| | | | | | (_) | (__|   < 
|_| |_| |_|\___/ \___|_|\_\
*/

--Create two abuse address
INSERT INTO public."abuse"(created, name) VALUES ('2019-01-25 13:44:45', 'address@abuse.com');
INSERT INTO public."abuse"(created, name) VALUES ('2019-01-25 13:44:45', 'another_abuse@abuse.com');

--Create one conservation
INSERT INTO public."conservation"(created, sent, unsent, keep_email_content) VALUES ('2019-01-25 13:44:45', '15', '10', true);

--Create two entities
INSERT INTO public."entity"(created, name, description, conservation_id, abuse_id) VALUES ('2019-01-25 13:44:45', 'entity-1', 'this is the first entity', 1, 1);
INSERT INTO public."entity"(created, name, conservation_id, abuse_id) VALUES ('2019-01-25 13:44:45', 'entity-2', 1, 1);

--Create three environments
INSERT INTO public."environment"(created, name, description, conservation_id, entity_id, is_open, abuse_id) VALUES ('2019-01-25 13:44:45', 'environment-1', 'this is an environment', 1, 1, false, 1);
INSERT INTO public."environment"(created, name, conservation_id, entity_id, is_open, abuse_id) VALUES ('2019-01-25 13:44:45', 'environment-2', 1, 2, false, 1);
INSERT INTO public."environment"(created, name, conservation_id, entity_id, is_open, abuse_id) VALUES ('2019-01-25 13:44:45', 'environment-3', 1, 2, true, 2);

--Create five domains
INSERT INTO public."domain"(created, name, conservation_id, environment_id) VALUES ('2019-01-25 13:44:45', 'domain-1', 1, 3);
INSERT INTO public."domain"(created, name, conservation_id, environment_id) VALUES ('2019-01-25 13:44:45', 'domain-2', 1, 2);
INSERT INTO public."domain"(created, name, conservation_id, environment_id) VALUES ('2019-01-25 13:44:45', 'domain-3', 1, 1);
INSERT INTO public."domain"(created, name, conservation_id, environment_id) VALUES ('2019-01-25 13:44:45', 'domain-4', 1, 3);
INSERT INTO public."domain"(created, name, conservation_id, environment_id) VALUES ('2019-01-25 13:44:45', 'domain-5', 1, 2);


--Create default quota
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 10, 5);
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 60, 10);
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 1440, 10);
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 10080, 100);
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 44640, 1000);

--Create custom quota
INSERT INTO public."quota"(created, duration_minutes, quantity) VALUES ('2019-01-25 13:44:45', 30, 7);


--Add default quota for the two entities
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (1, 1);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (1, 2);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (1, 3);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (1, 4);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (1, 5);

INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (2, 1);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (2, 2);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (2, 3);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (2, 4);
INSERT INTO public."entity_quotas"(entity_id, quota_id) VALUES (2, 5);

--Add default quota for the three environment
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (1, 1);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (1, 2);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (1, 3);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (1, 4);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (1, 5);

INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (2, 1);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (2, 2);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (2, 3);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (2, 4);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (2, 5);

INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 1);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 2);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 3);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 4);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 5);
INSERT INTO public."environment_quotas"(environment_id, quota_id) VALUES (3, 6);

--Add default quotas for domain
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (1, 1);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (1, 2);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (1, 3);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (1, 4);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (1, 5);

INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (2, 1);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (2, 2);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (2, 3);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (2, 4);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (2, 5);

INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (3, 1);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (3, 2);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (3, 3);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (3, 4);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (3, 5);

INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (4, 1);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (4, 2);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (4, 3);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (4, 4);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (4, 5);

INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (5, 1);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (5, 2);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (5, 3);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (5, 4);
INSERT INTO public."domain_quotas"(domain_id, quota_id) VALUES (5, 5);
