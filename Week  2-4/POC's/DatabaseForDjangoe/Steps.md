# first to install MYSQL and the lib run this command .
>sudo apt install mysql-server
>sudo apt install libmysqlclient-dev

# Second step inside the terminal open MySQL with username root by this command.
>sudo mysql -u root

# then create and give it all privileges by this commands .
>CREATE USER 'YOURUSERNAME'@'localhost' IDENTIFIED BY 'YOUR PASSWORD';

>GRANT ALL PRIVILEGES ON *.* TO 'YOURUSERNAME'@'localhost' WITH GRANT OPTION;

# Then create a database using this command .
>CREATE DATABASE YOURDATABASENAME;


# Inside Django go to your settings.py file and change the DATABASES to this .
>'default' : {
		'ENGINE' : 'django.db.backends.mysql',
		'NAME' : '', # The database you created above.
		'USER' : '', # The new user you created.
		'PASSWORD' : '', # The password you assigned to that user.
		'HOST' : 'localhost',
		'PORT' : '3306',
	}

# then Install MySQL Client Library usnig this command .
>pip install mysqlclient

# finally run this command to create tables in the MySQL database.
>python3 manage.py makemigrations

>python3 manage.py migrate


