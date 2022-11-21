create table employer
(
    id            serial8 not null
        primary key,
    alternate_url varchar(255),
    name          varchar(255),
    trusted       boolean,
    url           varchar(255),
    vacancies_url varchar(255)
);

alter table employer
    owner to postgres;

create table statistics
(
    id                     serial8 not null
        primary key,
    applicant_per_vacancy  numeric(19, 2),
    average_salary         integer,
    count_applicant        integer,
    date                   date,
    name                   varchar(255),
    salary_between100k150k integer,
    salary_between150k200k integer,
    salary_between200k250k integer,
    salary_between250k300k integer,
    salary_between300k350k integer,
    salary_between350k400k integer,
    salary_less100k        integer,
    salary_more400k        integer,
    total_vacancies        integer,
    without_salary         integer
);

alter table statistics
    owner to postgres;

create table vacancy
(
    id                       serial8 not null
        primary key,
    accept_temporary         boolean,
    alternate_url            varchar(255),
    archived                 boolean,
    area                     varchar(255),
    created_at               varchar(255),
    has_test                 boolean,
    name                     varchar(255),
    premium                  boolean,
    published_at             varchar(255),
    response_letter_required boolean,
    salary_avg               integer,
    salary_currency          varchar(255),
    salary_from              integer default 0,
    salary_gross             boolean,
    salary_to                integer default 0,
    schedule                 varchar(255),
    type                     varchar(255),
    url                      varchar(255),
    employer_id              bigint
);

alter table vacancy
    owner to postgres;


insert into statistics (applicant_per_vacancy, average_salary, count_applicant, date, name,
                               salary_between100k150k, salary_between150k200k, salary_between200k250k,
                               salary_between250k300k, salary_between300k350k, salary_between350k400k, salary_less100k,
                               salary_more400k, total_vacancies, without_salary)
values (4.30, 150049, 6189, '2022-07-15', '1C', 176, 258, 146, 31, 11, 3, 153, 0, 1439, 806),
       (null, 150220, null, '2022-07-16', '1C', 173, 259, 146, 31, 10, 3, 151, 0, 1429, 801),
       (null, 150294, null, '2022-07-17', '1C', 166, 249, 140, 30, 10, 3, 145, 0, 1388, 782),
       (null, 150311, null, '2022-07-18', '1C', 168, 249, 143, 31, 10, 3, 147, 0, 1392, 780),
       (null, 149785, null, '2022-07-19', '1C', 175, 255, 151, 33, 9, 3, 155, 0, 1461, 824),
       (5.65, 151266, 6189, '2022-07-23', '1C', 181, 263, 157, 32, 8, 2, 145, 0, 1435, 801),
       (null, 151521, null, '2022-07-25', '1C', 184, 262, 155, 33, 10, 2, 144, 0, 1434, 800),
       (null, 151103, null, '2022-07-27', '1C', 183, 269, 156, 32, 9, 2, 150, 0, 1456, 813),
       (null, 153355, null, '2022-07-29', '1C', 182, 259, 162, 36, 11, 2, 140, 0, 1436, 804),
       (null, 153344, null, '2022-08-01', '1C', 170, 255, 156, 32, 9, 2, 133, 0, 1390, 782),
       (null, 152015, null, '2022-08-03', '1C', 163, 258, 151, 31, 10, 3, 142, 0, 1398, 783),
       (4.22, 155949, 6068, '2022-08-12', '1C', 163, 274, 164, 40, 12, 3, 137, 0, 1436, 801),
       (null, 153698, null, '2022-08-25', '1C', 175, 264, 156, 42, 10, 3, 150, 0, 1433, 785),
       (null, 156064, 6314, '2022-09-01', '1C', 181, 254, 166, 47, 11, 3, 143, 0, 1448, 801),
       (null, 157127, null, '2022-09-05', '1C', 182, 248, 157, 51, 11, 3, 134, 0, 1426, 799),
       (4.36, 156459, 6348, '2022-09-10', '1C', 202, 259, 164, 50, 10, 3, 136, 0, 1453, 794),
       (null, 156911, null, '2022-09-13', '1C', 189, 257, 160, 49, 10, 3, 131, 0, 1436, 788),
       (null, 159479, null, '2022-09-16', '1C', 190, 266, 167, 51, 11, 4, 135, 1, 1466, 792),
       (null, 160059, null, '2022-09-19', '1C', 184, 260, 164, 50, 11, 4, 127, 1, 1424, 767),
       (null, 159756, null, '2022-09-23', '1C', 189, 264, 178, 48, 9, 2, 128, 1, 1462, 786),
       (null, 159610, null, '2022-09-25', '1C', 185, 257, 167, 46, 9, 2, 125, 1, 1438, 782),
       (null, 158212, null, '2022-10-06', '1C', 183, 246, 165, 54, 10, 1, 138, 1, 1407, 762),
       (null, 182899, null, '2022-09-16', 'FRONTEND', 91, 127, 104, 66, 33, 15, 68, 1, 1275, 830),
       (null, 180603, null, '2022-09-19', 'FRONTEND', 81, 120, 96, 59, 29, 13, 66, 1, 1220, 804),
       (null, 181007, null, '2022-09-23', 'FRONTEND', 81, 126, 99, 57, 30, 13, 65, 1, 1226, 803),
       (null, 181705, null, '2022-09-25', 'FRONTEND', 75, 115, 99, 54, 27, 13, 61, 1, 1180, 778),
       (null, 179712, null, '2022-10-06', 'FRONTEND', 76, 111, 86, 50, 26, 13, 60, 1, 1113, 736),
       (2.40, 245012, 876, '2022-07-15', 'GOLANG', 9, 17, 17, 15, 11, 9, 5, 3, 364, 287),
       (null, 248858, null, '2022-07-16', 'GOLANG', 8, 17, 17, 15, 11, 10, 4, 3, 359, 282),
       (null, 251044, null, '2022-07-17', 'GOLANG', 8, 15, 16, 15, 11, 10, 4, 3, 342, 268),
       (null, 253158, null, '2022-07-18', 'GOLANG', 7, 15, 16, 15, 11, 10, 4, 3, 339, 266),
       (null, 252806, null, '2022-07-19', 'GOLANG', 7, 16, 18, 18, 11, 10, 4, 3, 351, 274),
       (3.24, 243723, 1070, '2022-07-23', 'GOLANG', 7, 18, 17, 15, 11, 8, 7, 3, 330, 255),
       (null, 249505, null, '2022-07-25', 'GOLANG', 7, 16, 17, 17, 11, 9, 7, 3, 332, 257),
       (null, 254176, null, '2022-07-27', 'GOLANG', 7, 16, 17, 18, 12, 11, 7, 3, 339, 261),
       (null, 262225, null, '2022-07-29', 'GOLANG', 7, 14, 17, 18, 14, 10, 7, 4, 332, 255),
       (null, 263338, null, '2022-08-01', 'GOLANG', 8, 12, 17, 16, 14, 10, 7, 4, 342, 269),
       (null, 259253, null, '2022-08-03', 'GOLANG', 8, 11, 17, 15, 12, 10, 8, 4, 345, 273),
       (2.74, 256967, 961, '2022-08-12', 'GOLANG', 14, 13, 14, 15, 13, 9, 7, 5, 350, 276),
       (null, 238822, null, '2022-08-25', 'GOLANG', 20, 14, 12, 13, 13, 12, 10, 4, 375, 292),
       (null, 240048, 1024, '2022-09-01', 'GOLANG', 18, 11, 14, 12, 13, 13, 10, 4, 370, 285),
       (null, 240388, null, '2022-09-05', 'GOLANG', 18, 12, 14, 12, 13, 13, 8, 3, 358, 276),
       (2.87, 256530, 1051, '2022-09-10', 'GOLANG', 12, 12, 12, 11, 13, 12, 6, 3, 365, 291),
       (null, 249377, null, '2022-09-13', 'GOLANG', 12, 11, 11, 10, 12, 10, 6, 3, 355, 287),
       (null, 240375, null, '2022-09-16', 'GOLANG', 13, 12, 14, 11, 11, 9, 8, 3, 371, 297),
       (null, 241202, null, '2022-09-19', 'GOLANG', 13, 10, 13, 8, 10, 9, 8, 3, 368, 298),
       (null, 228507, null, '2022-09-23', 'GOLANG', 12, 12, 13, 14, 13, 10, 14, 2, 395, 312),
       (null, 230427, null, '2022-09-25', 'GOLANG', 11, 12, 13, 14, 13, 9, 12, 2, 375, 297),
       (null, 228900, null, '2022-10-06', 'GOLANG', 12, 12, 13, 14, 16, 6, 9, 2, 364, 287),
       (6.20, 236495, 6130, '2022-07-15', 'JAVA', 13, 39, 36, 32, 30, 18, 9, 1, 986, 815),
       (null, 234439, null, '2022-07-16', 'JAVA', 12, 39, 36, 29, 27, 17, 9, 1, 950, 786),
       (null, 233409, null, '2022-07-17', 'JAVA', 12, 40, 34, 29, 26, 17, 9, 1, 920, 760),
       (null, 234171, null, '2022-07-18', 'JAVA', 12, 39, 34, 27, 27, 17, 9, 1, 918, 759),
       (null, 237171, null, '2022-07-19', 'JAVA', 10, 42, 35, 30, 30, 16, 9, 2, 953, 787),
       (6.87, 236041, 6319, '2022-07-23', 'JAVA', 10, 38, 32, 28, 30, 18, 12, 0, 919, 756),
       (null, 236599, null, '2022-07-25', 'JAVA', 10, 38, 31, 26, 29, 18, 11, 0, 923, 764),
       (null, 239656, null, '2022-07-27', 'JAVA', 11, 39, 30, 26, 31, 19, 10, 1, 942, 779),
       (null, 243384, null, '2022-07-29', 'JAVA', 13, 38, 32, 26, 31, 17, 9, 4, 950, 790),
       (null, 240013, null, '2022-08-01', 'JAVA', 12, 38, 32, 25, 28, 17, 10, 4, 949, 789),
       (null, 236620, null, '2022-08-03', 'JAVA', 14, 42, 34, 27, 28, 16, 10, 4, 987, 815),
       (6.72, 230085, 6390, '2022-08-12', 'JAVA', 22, 37, 29, 22, 30, 16, 14, 4, 950, 785),
       (null, 220424, null, '2022-08-25', 'JAVA', 23, 29, 33, 22, 28, 11, 14, 3, 923, 764),
       (null, 217019, 6665, '2022-09-01', 'JAVA', 16, 30, 34, 22, 22, 9, 12, 1, 887, 742),
       (null, 220811, null, '2022-09-05', 'JAVA', 17, 29, 31, 25, 24, 8, 10, 1, 867, 727),
       (7.30, 222265, 6762, '2022-09-10', 'JAVA', 18, 33, 34, 29, 28, 9, 10, 0, 926, 772),
       (null, 219962, null, '2022-09-13', 'JAVA', 18, 35, 36, 26, 25, 7, 8, 0, 908, 756),
       (null, 220948, null, '2022-09-16', 'JAVA', 19, 32, 35, 27, 25, 6, 7, 0, 938, 789),
       (null, 222781, null, '2022-09-19', 'JAVA', 19, 32, 34, 27, 26, 8, 7, 0, 906, 758),
       (null, 221936, null, '2022-09-23', 'JAVA', 18, 33, 37, 27, 25, 7, 7, 0, 937, 786),
       (null, 224182, null, '2022-09-25', 'JAVA', 18, 31, 35, 28, 26, 8, 7, 0, 911, 765),
       (null, 235171, null, '2022-10-06', 'JAVA', 12, 27, 33, 29, 31, 8, 7, 1, 917, 780),
       (4.80, 171897, 2806, '2022-07-15', 'PHP', 77, 92, 69, 37, 18, 7, 47, 1, 589, 302),
       (null, 170960, null, '2022-07-16', 'PHP', 78, 91, 69, 36, 18, 7, 50, 1, 589, 299),
       (null, 171511, null, '2022-07-17', 'PHP', 77, 86, 65, 34, 17, 6, 45, 1, 561, 287),
       (null, 172185, null, '2022-07-18', 'PHP', 78, 85, 64, 34, 18, 6, 45, 1, 560, 286),
       (null, 172898, null, '2022-07-19', 'PHP', 84, 93, 69, 38, 21, 8, 49, 1, 597, 305),
       (4.97, 172302, 2890, '2022-07-23', 'PHP', 78, 87, 62, 34, 21, 8, 50, 1, 581, 299),
       (null, 172456, null, '2022-07-25', 'PHP', 76, 85, 64, 34, 19, 7, 46, 1, 580, 305),
       (null, 169945, null, '2022-07-27', 'PHP', 84, 90, 64, 32, 20, 6, 48, 1, 595, 311),
       (null, 172067, null, '2022-07-29', 'PHP', 81, 100, 68, 31, 19, 5, 48, 1, 595, 308),
       (null, 171372, null, '2022-08-01', 'PHP', 74, 97, 66, 28, 18, 5, 46, 1, 594, 316),
       (null, 170103, null, '2022-08-03', 'PHP', 79, 101, 63, 28, 19, 5, 48, 1, 602, 318),
       (4.87, 170040, 2916, '2022-08-12', 'PHP', 81, 101, 64, 29, 20, 5, 49, 0, 598, 307),
       (null, 171538, null, '2022-08-25', 'PHP', 80, 99, 73, 29, 19, 7, 58, 1, 602, 299),
       (null, 173071, 3012, '2022-09-01', 'PHP', 82, 102, 79, 31, 18, 7, 51, 0, 599, 294),
       (null, 174058, null, '2022-09-05', 'PHP', 77, 100, 77, 32, 17, 7, 45, 0, 581, 287),
       (4.94, 170631, 2968, '2022-09-10', 'PHP', 81, 100, 77, 32, 17, 7, 54, 0, 600, 298),
       (null, 172775, null, '2022-09-13', 'PHP', 82, 95, 78, 33, 19, 7, 49, 0, 593, 295),
       (null, 173361, null, '2022-09-16', 'PHP', 88, 102, 72, 39, 21, 5, 48, 0, 602, 292),
       (null, 172500, null, '2022-09-19', 'PHP', 88, 98, 66, 38, 18, 4, 45, 0, 566, 273),
       (null, 169908, null, '2022-09-23', 'PHP', 92, 105, 65, 41, 18, 3, 48, 0, 574, 266),
       (null, 169438, null, '2022-09-25', 'PHP', 97, 102, 65, 41, 17, 2, 46, 0, 570, 264),
       (null, 172738, null, '2022-10-06', 'PHP', 84, 89, 60, 41, 19, 3, 47, 1, 529, 242),
       (11.50, 192777, 6537, '2022-07-15', 'PYTHON', 37, 40, 42, 37, 18, 6, 25, 1, 568, 387),
       (null, 192991, null, '2022-07-16', 'PYTHON', 36, 37, 44, 37, 18, 6, 26, 1, 572, 392),
       (null, 193172, null, '2022-07-17', 'PYTHON', 34, 34, 40, 36, 18, 6, 26, 1, 548, 377),
       (null, 193300, null, '2022-07-18', 'PYTHON', 34, 36, 42, 39, 18, 6, 27, 1, 549, 375),
       (null, 196096, null, '2022-07-19', 'PYTHON', 36, 35, 40, 43, 20, 6, 26, 1, 567, 387),
       (12.05, 194450, 6764, '2022-07-23', 'PYTHON', 37, 33, 42, 39, 16, 5, 25, 1, 561, 381),
       (null, 197211, null, '2022-07-25', 'PYTHON', 35, 34, 42, 39, 18, 5, 24, 2, 554, 376),
       (null, 197920, null, '2022-07-27', 'PYTHON', 37, 34, 41, 39, 19, 6, 24, 2, 560, 381),
       (null, 197517, null, '2022-07-29', 'PYTHON', 37, 34, 39, 38, 18, 5, 25, 3, 556, 380),
       (null, 196831, null, '2022-08-01', 'PYTHON', 36, 37, 36, 37, 17, 5, 24, 3, 547, 376),
       (null, 195632, null, '2022-08-03', 'PYTHON', 40, 39, 36, 39, 18, 6, 24, 3, 562, 387),
       (12.15, 198062, 6817, '2022-08-12', 'PYTHON', 36, 40, 39, 37, 17, 5, 25, 4, 561, 385),
       (null, 192292, null, '2022-08-25', 'PYTHON', 37, 37, 46, 33, 16, 4, 26, 4, 542, 375),
       (null, 198029, 7072, '2022-09-01', 'PYTHON', 35, 29, 42, 31, 12, 4, 17, 3, 523, 377),
       (null, 197056, null, '2022-09-05', 'PYTHON', 35, 25, 39, 29, 10, 4, 17, 3, 500, 360),
       (13.21, 202036, 7137, '2022-09-10', 'PYTHON', 30, 30, 44, 33, 15, 7, 18, 2, 540, 390),
       (null, 199838, null, '2022-09-13', 'PYTHON', 33, 30, 42, 33, 17, 7, 20, 2, 538, 382),
       (null, 198434, null, '2022-09-16', 'PYTHON', 33, 28, 37, 30, 17, 7, 20, 2, 542, 391),
       (null, 200550, null, '2022-09-19', 'PYTHON', 33, 28, 39, 31, 17, 7, 18, 2, 527, 377),
       (null, 204954, null, '2022-09-23', 'PYTHON', 30, 32, 40, 36, 17, 7, 15, 2, 535, 380),
       (null, 203286, null, '2022-09-25', 'PYTHON', 29, 30, 37, 34, 15, 6, 15, 2, 517, 368),
       (null, 204658, null, '2022-10-06', 'PYTHON', 26, 40, 32, 30, 16, 7, 15, 2, 531, 386),
       (null, 161021, null, '2022-11-08', '1C', 178, 281, 175, 50, 11, 2, 128, 1, 1457, 796),
       (null, 185210, null, '2022-11-08', 'FRONTEND', 69, 86, 75, 54, 26, 18, 56, 1, 1093, 739),
       (null, 228565, null, '2022-11-08', 'GOLANG', 10, 13, 17, 20, 13, 8, 12, 3, 371, 291),
       (null, 248389, null, '2022-11-08', 'JAVA', 8, 31, 25, 35, 29, 10, 6, 2, 940, 803),
       (null, 178102, null, '2022-11-08', 'PHP', 54, 88, 59, 33, 14, 7, 39, 2, 507, 257),
       (null, 200301, null, '2022-11-08', 'PYTHON', 31, 32, 40, 28, 16, 8, 16, 2, 547, 405);