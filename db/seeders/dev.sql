INSERT INTO main.places(name, description, lat, lng, address, website)
VALUES
        ('Torro Grill', 'Какое-то описание ресторана Torro Grill', 55.778001, 37.586297, 'Россия, Москва, Лесная ул., 5, стр. Б', 'https://www.torrogrill.ru/'),
        ('Марукамэ', 'Азиатский ресторан Марукамэ', 55.778907, 37.587278, 'Россия, Москва, 4-й Лесной пер., 4', 'http://marukame.ru/'),
        ('Кофемания', 'Дорогая кафешка Кофемания', 55.777797, 37.586394, 'Россия, Москва, Лесная улица, 5', 'https://coffeemania.ru/'),
        ('Silver Panda', 'Недорогой азиатский ресторан, где можно дешёво пообедать.', 55.741206, 37.609505, 'Россия, Москва, Берсеневский переулок, 2с1', 'http://silverpanda.ru/'),
        ('Howard Loves Craft', 'Классный бар где можно попить кравтовое пиво с друзьями после работы.', 55.740347, 37.610544, 'Россия, Москва, Болотная набережная, 7с3', 'http://howardlovescraft.ru/');

INSERT INTO main.users(phone, name, email, password, photo_path)
VALUES
        ('+79151234567', 'Johny', 'johny@mail.ru', '123', null),
        ('+79151234568', 'Nick', 'nick@mail.ru', '1234', null),
        ('+79151234569', 'Alex', 'alex@mail.ru', '111', null);

INSERT INTO main.evaluation_criterions(name) VALUES ('Кухня'), ('Сервис'), ('Атмосфера');

INSERT INTO main.places_evaluations(place_id, user_id, comment)
VALUES
        (1, 3, 'Большие порции, высокое качество'),
        (4, 3, 'Есть недорогие зожные позиции'),
        (2, 1, 'Вкусная азиатская кафешка, невысокие цены и быстрое обслуживание'),
        (1, 1, 'Еда очень вкусная, но обслуживание могло бы быть получше =(');

INSERT INTO main.place_evaluation_marks(place_evaluation_id, evaluation_criterion_id, mark)
VALUES
        (1, 1, 'excellent'),
        (1, 2, 'excellent'),
        (1, 3, 'excellent'),
        (2, 1, 'good'),
        (2, 2, 'good'),
        (2, 3, 'good'),
        (3, 1, 'excellent'),
        (3, 2, 'good'),
        (3, 3, 'good'),
        (4, 1, 'excellent'),
        (4, 2, 'bad'),
        (4, 3, 'good');

INSERT INTO main.categories(name)
VALUES
        ('Кофейня'),
        ('Фастфуд'),
        ('Ресторан'),
        ('Бар'),
        ('На вынос');

INSERT INTO main.places_categories(place_id, category_id)
VALUES
        (1, 3),
        (2, 2),
        (2, 3),
        (2, 5),
        (3, 1),
        (3, 3),
        (4, 2),
        (4, 5),
        (5, 4);

INSERT INTO main.kitchen_categories(name)
VALUES
        ('Стейкхайс'),
        ('Пиццерия'),
        ('Азиатская кухня'),
        ('Бургерная'),
        ('Завтраки'),
        ('Кофе'),
        ('Чайная'),
        ('Итальянская кухня'),
        ('Французская кухня');

INSERT INTO main.places_kitchen_categories(place_id, kitchen_category_id)
VALUES
        (1, 1),
        (2, 3),
        (3, 5),
        (3, 6),
        (4, 3),
        (5, 4);

INSERT INTO main.restaurateur_roles(name) VALUES ('admin'), ('moderator');

INSERT INTO main.restaurateurs (email, password, email_verified, reg_datetime)
VALUES
       ('bob@torro.ru','$2a$10$R/jeXpRiTktW6IdEEgEmsuW3tSA41AJdL4i2agBDlOE.ul9q8Vgam',true,'2021-08-25 11:20:13.362000 +00:00'),
       ('john@mail.ru','$2a$10$R/jeXpRiTktW6IdEEgEmsuW3tSA41AJdL4i2agBDlOE.ul9q8Vgam',true,'2021-08-25 11:24:46.551000 +00:00');

INSERT INTO main.places_restaurateurs (restaurateur_id, place_id, restaurateur_role_id)
VALUES
       (1,1,1),
       (2,3,1),
       (2,4,1),
       (2,5,1);

INSERT INTO main.adverts(place_id, restaurateur_id, text, publish_datetime)
VALUES
        (1, 1, 'АКЦИЯ! Каждый вторник второй бургер бесплатно!', '2021-07-10 04:05:06 +3:00'),
        (5, 2, '24.07.2021 в 19:00 в нашем баре будет выступать группа "Бременские музыканты"', '2021-07-21 07:00:00 +3:00'),
        (1, 1, 'Мы открыли летнюю веранду в ресторане на Белорусской!', '2021-07-21 08:50:22 +3:00');

INSERT INTO main.users_subscriptions(follower_user_id, followed_user_id)
VALUES (1,3),
       (2,3),
       (1,2);
