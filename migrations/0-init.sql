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
                            name varchar(255) not null,
                            phone varchar(15) not null,
                            email varchar(255),
                            password varchar(255),
                            email_verified bool not null default false,
                            reg_datetime timestamptz not null default now(),
                            photo_path varchar(255)
);

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
                              text text,
                              publish_datetime timestamptz not null default now()
);

create table main.review_medias (
                                    review_media_id serial primary key,
                                    review_id int references main.reviews(review_id),
                                    media_type main.media_type not null,
                                    media_path varchar not null,
                                    "order" int not null
);

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
                              place_id int references main.places(place_id),
                              text text,
                              publish_datetime timestamptz not null default now()
);

create table main.advert_medias (
                                    advert_media_id serial primary key,
                                    advert_id int references main.adverts(advert_id),
                                    media_type main.media_type not null,
                                    media_path varchar not null,
                                    "order" int not null
);

create table main.users_subscriptions (
                                          follower_user_id int references main.users(user_id),
                                          followed_user_id int references main.users(user_id),
                                          primary key (follower_user_id, followed_user_id)
);

-- +migrate Down