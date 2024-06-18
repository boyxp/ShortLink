CREATE TABLE IF  NOT EXISTS url(
   id int(10) unsigned NOT NULL AUTO_INCREMENT,
   hash VARCHAR(16) default '' NOT NULL,
   url VARCHAR(300) default '' NOT NULL,
   password VARCHAR(40) default '' NOT NULL,
   expire_at timestamp NOT NULL DEFAULT '2029-01-01 00:00:00' COMMENT '失效时间',
   create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
   PRIMARY KEY ( id )
);
