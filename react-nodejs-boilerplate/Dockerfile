FROM mhart/alpine-node

RUN mkdir -p /dummy_front
WORKDIR /dummy_front
ADD ./server /dummy_front
RUN yarn install
EXPOSE 8000

CMD [ "yarn", "start:docker" ]