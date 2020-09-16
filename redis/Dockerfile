FROM redis

COPY redis.conf /usr/local/etc/redis/redis.conf

RUN mkdir /var/lib/redis && chown redis:redis /var/lib/redis/

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]

# For latest redis.conf
# http://download.redis.io/redis-stable/redis.conf