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

--Create one conservation
INSERT INTO public."conservation"(created, sent, unsent, keep_email_content) VALUES ('2019-01-25 13:44:45', '15', '10', true);
INSERT INTO public."conservation"(created, sent, unsent, keep_email_content) VALUES ('2018-01-25 08:44:45', '5', '20', false);
