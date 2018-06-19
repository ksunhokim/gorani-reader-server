FROM mysql:8

ADD start_db.sh init_db.sh init.sql schema.sql /tmp/

RUN /tmp/init_db.sh

CMD ["/tmp/start_db.sh"]
