-- +migrate Up

create schema main;

create table main.places (
                             place_id serial primary key,
                             name varchar(255),
                             description text,
                             lat double precision,
                             lng double precision,
                             address varchar(255),
                             website varchar(255),
                             rating numeric,
                             mark_amount int default 0 not null
);

create type main.media_type as enum ('photo', 'video');

create table main.place_media (
                                  place_media_id serial primary key,
                                  place_id int references main.places(place_id) not null,
                                  media_path varchar not null,
                                  media_type main.media_type not null,
                                  comment text,
                                  publish_datetime timestamptz default now() not null
);

create table main.users (
                            user_id serial primary key,
                            name varchar(255) not null unique,
                            phone varchar(15) not null unique,
                            email varchar(255) unique,
                            password varchar(255),
                            phone_verified bool not null default false,
                            email_verified bool not null default false,
                            reg_datetime timestamptz not null default now(),
                            photo_path varchar(255)
);

CREATE UNIQUE INDEX ON main.users (phone, phone_verified) WHERE (users.phone_verified is true);

create type main.mark as enum ('excellent', 'good', 'bad');

create table main.evaluation_criterions (
                                            evaluation_criterion_id serial primary key,
                                            name varchar(255) not null
);

create table main.places_evaluations (
                                         place_evaluation_id serial primary key,
                                         place_id int references main.places(place_id),
                                         user_id int references main.users(user_id),
                                         datetime timestamptz not null default now(),
                                         comment text,
                                         unique (place_id, user_id)
);

create table main.place_evaluation_marks (
                                             place_evaluation_id int references main.places_evaluations(place_evaluation_id),
                                             evaluation_criterion_id int references main.evaluation_criterions(evaluation_criterion_id) not null,
                                             mark main.mark not null,
                                             primary key (place_evaluation_id, evaluation_criterion_id)
);

create table main.places_schedules (
                                       place_schedule_id serial primary key,
                                       place_id int references main.places(place_id),
                                       day_of_week int,
                                       start_time time,
                                       end_time time,
                                       date_start date,
                                       date_stop date
);

create table main.categories (
                                 category_id serial primary key,
                                 name varchar(255) not null
);

create table main.places_categories (
                                        place_id int references main.places(place_id),
                                        category_id int references main.categories(category_id),
                                        primary key (place_id, category_id)
);

create table main.reviews (
                              review_id serial primary key,
                              user_id int references main.users(user_id),
                              place_id int references main.places(place_id) not null,
                              text text,
                              publish_datetime timestamptz not null default now()
);

CREATE INDEX idx_reviews_user_id_publish_datetime ON main.reviews(user_id, publish_datetime);

create table main.review_medias (
                                    review_media_id serial primary key,
                                    user_id int references main.users(user_id),
                                    media_type main.media_type not null,
                                    media_path varchar not null
);

create table main.review_review_medias (
                                    review_id int references main.reviews(review_id),
                                    review_media_id int references main.review_medias(review_media_id),
                                    "order" int not null,
                                    primary key (review_id, review_media_id)
);

CREATE UNIQUE INDEX ON main.review_review_medias (review_id, "order");

create table main.kitchen_categories (
                                         kitchen_category_id serial primary key,
                                         name varchar(255)
);

create table main.places_kitchen_categories (
                                                place_id int references main.places(place_id),
                                                kitchen_category_id int references main.kitchen_categories(kitchen_category_id),
                                                primary key (place_id, kitchen_category_id)
);

create table main.adverts (
                              advert_id serial primary key,
                              place_id int references main.places(place_id) not null,
                              restaurateur_id int references main.restaurateurs(restaurateur_id) not null,
                              text text,
                              publish_datetime timestamptz not null default now()
);

create table main.advert_medias (
                                    advert_media_id serial primary key,
                                    place_id integer references main.places(place_id) not null,
                                    restaurateur_id int references main.restaurateurs(restaurateur_id) not null,
                                    media_type main.media_type not null,
                                    media_path varchar not null
);

create table main.advert_advert_medias (
                                           advert_id int references main.adverts(advert_id),
                                           advert_media_id int references main.advert_medias(advert_media_id),
                                           "order" int not null,
                                           primary key (advert_id, advert_media_id)
);

create table main.users_subscriptions (
                                          follower_user_id int references main.users(user_id),
                                          followed_user_id int references main.users(user_id),
                                          primary key (follower_user_id, followed_user_id)
);

create table main.user_phone_codes (
                                       user_phone_code_id serial primary key,
                                       user_id serial references main.users(user_id),
                                       code varchar(10),
                                       create_datetime timestamptz default now() not null,
                                       actual bool default true,
                                       left_attempts integer not null default 0
);
CREATE UNIQUE INDEX ON main.user_phone_codes (user_id, actual) WHERE (actual is true);

create table main.restaurateurs (
                                    restaurateur_id serial primary key,
                                    email varchar(255) unique,
                                    password varchar(255),
                                    email_verified bool not null default false,
                                    reg_datetime timestamptz not null default now()
);

create table main.restaurateur_roles (
                                         restaurateur_role_id serial primary key,
                                         name varchar(255) unique
);

create table main.places_restaurateurs (
                                           restaurateur_id integer references main.restaurateurs(restaurateur_id),
                                           place_id integer references main.places(place_id),
                                           restaurateur_role_id integer references main.restaurateur_roles(restaurateur_role_id),
                                           primary key (restaurateur_id, place_id)
);

-- +migrate Down