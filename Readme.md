# README 



* Database create tables: 
<pre>
create table Students(
student_id int,
first_name VARCHAR(50) NOT NULL,
last_name VARCHAR(50) NOT NULL,
login VARCHAR(50) not null,
password VARCHAR(50) not null,
birth_date Date not null,
registredOn timestamp(2),
lastlogin timestamp(2),
constraint student_pk primary key (student_id)
)

create table Courses(
course_id int,
classroom_id int,
courseCode Varchar(50) not null,
teacher_id int,
constraint course_pk primary key (course_id),
foreign key(classroom_id) references classrooms(classroom_id),
foreign key(teacher_id) references teachers(teacher_id)

)

create table Classrooms(
classroom_id int,
capasity int,
constraint classroom_pk primary key (classroom_id)
)

create table Teachers(
teacher_id int,
first_name VARCHAR(50) NOT NULL,
last_name VARCHAR(50) NOT NULL,
login VARCHAR(50) not null,
password VARCHAR(50) not null,
registredOn timestamp(2),
lastlogin timestamp(2),
constraint teacher_pk primary key (teacher_id)
)

create table Attendances(
course_id int,
student_id int,
attDate date,
stamp char(1),
foreign key(course_id) references courses(course_id),
foreign key(student_id) references students(student_id)
)
</pre>"# database_project" 
