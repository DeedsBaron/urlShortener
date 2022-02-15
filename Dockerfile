FROM postgres

COPY ./urlShortener/init.sh ./

CMD bash ./init.sh