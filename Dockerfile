FROM postgres

COPY ./urlShortener/urlShortener ./

CMD bash ./urlShortener