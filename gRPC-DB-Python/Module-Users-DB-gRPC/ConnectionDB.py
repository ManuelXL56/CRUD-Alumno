from dotenv import load_dotenv
from pathlib import Path
import os
import pymysql

def getConnectionDB():
    dotenv_path = Path('./.env.local')
    load_dotenv(dotenv_path=dotenv_path)
    return pymysql.connect(
        host=os.getenv('MYSQL_HOST'),
        port=int(os.getenv('MYSQL_PORT')),
        database=os.getenv('MYSQL_DB'),
        user=os.getenv('MYSQL_USER'),
        passwd=os.getenv('MYSQL_PASSWORD'),
        cursorclass=pymysql.cursors.DictCursor
    )